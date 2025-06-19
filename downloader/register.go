package downloader

import (
	"context"
	"errors"
	"gozam/audiofingerprint"
	"gozam/db"
	"gozam/models"
	"gozam/utils"
	"gozam/wav"
	"log"
	"path/filepath"
	"strings"

	"github.com/redis/go-redis/v9"

	"gorm.io/gorm"
)

var DOWNLOAD_PATH = "./downloads"

func ProcessSong(songTitle string, songID uint32) (map[uint32]models.Couple, error) {
	// converted the song to wav
	songPath, err := utils.FindDownloadedFile(songTitle, DOWNLOAD_PATH)
	if err != nil {
		return nil, errors.New("song file can not be found")
	}
	wav.ConvertToWAV(songPath, 1) //stereo to mono audio and .wav format

	log.Print("converted the video to wav")

	// make wav into bytes
	wavPath := strings.TrimSuffix(songPath, filepath.Ext(songPath)) + ".wav"
	wavInfo, err := wav.ReadWavInfo(wavPath)
	if err != nil {
		log.Fatalf("error, %v", err)
		return nil, err
	}

	// making wavbytes from samples
	samples, err := wav.WavBytesToSamples(wavInfo.Data)
	if err != nil {
		log.Fatalf("error converting wav bytes to float64: %v", err)
		return nil, err

	}

	log.Print("converted to samples")
	// log.Print("erm what thw sigma")
	// log.Print(samples)

	//creating spectogram
	spectrogram, err := audiofingerprint.Spectrogram(samples, wavInfo.SampleRate)
	if err != nil {
		log.Fatalf("error creating spectrogram: %v", err)
		return nil, err
	}
	log.Print("created the spectogram")

	//viusalize the spectrogram in freq vs time. intensity based on db
	// magSpec, err := audiofingerprint.MagnitudeSpectrogram(spectrogram)
	// if err != nil {
	// 	log.Fatalf("error getting magnitudes of the spectrogram: %v", err)
	// }

	// output_path := "./downloads/spectrograms/viva_la_vida_spectrogram.png"
	// audiofingerprint.SaveSpectrogramImage(magSpec, output_path)

	// extract peaks ie most significant frequencies from each band
	peaks := audiofingerprint.ExtractPeaks(spectrogram, wavInfo.Duration)
	log.Print("extracted the peaks")

	//create fingerprints
	fingerprints := audiofingerprint.CreateFingerprint(peaks, songID)
	log.Print("created the fingerprints. count: ", len(fingerprints))

	return fingerprints, nil
}

// single song download-->process-->fingerprints in redis and metaData in postgres
func RegisterSong(ctx context.Context, redisClient *redis.Client, postgresClient *gorm.DB, url string) error {

	//download the metadata of song
	data, err := GetVideoDetails(url)
	if err != nil {
		return err
	}
	if len(data.Items) == 0 {
		return errors.New("no video metadata found for the provided URL")
	}

	songData := data.Items[0]
	ytID := songData.ID
	songTitle := songData.Snippet.Title
	songArtist := songData.Snippet.ChannelTitle
	songID := utils.GenerateUniqueID()
	log.Print(songID, ytID, songTitle, songArtist)

	//store the song in postgres
	NewSong := models.Song{
		ID:     songID,
		Title:  songTitle,
		Artist: songArtist,
		YtID:   ytID,
	}

	//song already in database
	exists, err := db.IsSonginDB(ytID, postgresClient)
	if err != nil {
		return err
	}
	if exists {
		log.Println("Song already exists in the database. Skipping download and processing.")
		return nil
	}

	//downlaod the song
	err = DownloadYTaudio(url, DOWNLOAD_PATH)
	if err != nil {
		return err
	}

	log.Print("downloaded the video")

	fingerprints, err := ProcessSong(songTitle, songID)
	if err != nil {
		return err

	}
	// save fingerprints to redis
	err = db.StoreFingerprints(ctx, redisClient, fingerprints)
	if err != nil {
		log.Fatalf("Failed to store fingerprints: %v", err)
	}
	log.Print("succesfully saved the fingerprints in redis")

	err = db.InsertSong(&NewSong, postgresClient)
	if err != nil {
		log.Fatalf("Failed to store song: %v", err)
	}
	log.Print("succesfully saved the song in postgres")

	return nil
}

// handle yt playlist download
func RegisterPlaylist(ctx context.Context, redisClient *redis.Client, postgresClient *gorm.DB, playlistURL string) error {
	//download the metadata of song
	data, err := GetPlaylistDetails(playlistURL)
	if err != nil {
		return err
	}
	if len(data.Items) == 0 {
		return errors.New("no video metadata found for the provided URL")
	}

	for x := range len(data.Items) {
		songData := data.Items[x]
		ytID := songData.ContentDetails.ID
		songTitle := songData.Snippet.Title
		songArtist := songData.Snippet.ChannelTitle
		songID := utils.GenerateUniqueID()
		log.Print(songID, ytID, songTitle, songArtist)

		//song already in database
		exists, err := db.IsSonginDB(ytID, postgresClient)
		if err != nil {
			return err
		}
		if exists {
			log.Println("Song already exists in the database. Skipping download and processing.")

		} else {
			//downlaod the song
			url := "https://www.youtube.com/watch?v=" + ytID
			err = DownloadYTaudio(url, DOWNLOAD_PATH)
			if err != nil {
				return err
			}

			log.Print("downloaded the video")

			fingerprints, err := ProcessSong(songTitle, songID)
			if err != nil {
				return err

			}

			// save fingerprints to redis
			err = db.StoreFingerprints(ctx, redisClient, fingerprints)
			if err != nil {
				log.Fatalf("Failed to store fingerprints: %v", err)
			}
			log.Print("succesfully saved the fingerprints in redis")

			//store the song in postgres
			NewSong := models.Song{
				ID:     songID,
				Title:  songTitle,
				Artist: songArtist,
				YtID:   ytID,
			}

			err = db.InsertSong(&NewSong, postgresClient)
			if err != nil {
				log.Fatalf("Failed to store song: %v", err)
			}
			log.Print("succesfully saved the song in postgres")
		}

	}
	return nil

}

// finds top 20 macthes for requested file
func FindClientMatches(songPath string, postgresClient *gorm.DB) ([]audiofingerprint.Match, error) {
	//convert to wav
	wav.ConvertToWAV(songPath, 1) //stereo to mono audio and .wav format
	log.Print("converted the video to wav")
	// make wav into bytes

	wavPath := strings.TrimSuffix(songPath, filepath.Ext(songPath)) + ".wav"
	wavInfo, err := wav.ReadWavInfo(wavPath)
	if err != nil {
		log.Fatalf("error, %v", err)
		return nil, err
	}

	// making wavbytes from samples
	samples, err := wav.WavBytesToSamples(wavInfo.Data)
	if err != nil {
		log.Fatalf("error converting wav bytes to float64: %v", err)
		return nil, err

	}

	log.Print("converted to samples")
	// log.Print("erm what thw sigma")
	// log.Print(samples)

	matches, searchDuration, err := audiofingerprint.FindMatches(samples, wavInfo.Duration, wavInfo.SampleRate, postgresClient)
	if err != nil {
		log.Fatalf("failed to find matches: %v", err)
		return nil, err
	}

	if len(matches) == 0 {
		log.Println("\nNo match found.")
		log.Printf("\nSearch took: %s\n", searchDuration)
		return nil, nil
	}

	topMatches := matches
	if len(matches) >= 20 {
		topMatches = matches[:20]
	}
	return topMatches, nil
}
