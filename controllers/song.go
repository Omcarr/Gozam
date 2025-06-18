package controllers

import (
	"context"
	"errors"
	"gozam/audiofingerprint"
	"gozam/db"
	"gozam/downloader"
	"gozam/utils"
	"gozam/wav"
	"log"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"

	"gorm.io/gorm"
)

var DOWNLOAD_PATH = "./downloads"

func SongHealthCheck(cxt *gin.Context) {
	cxt.JSON(200, gin.H{
		"message": "apdi pode",
	})
}

func RegisterSongHandler(c *gin.Context, redisClient *redis.Client, postgresClient *gorm.DB) {
	type RequestBody struct {
		URL string `json:"url" binding:"required"`
	}

	var req RequestBody
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing or invalid URL in request body"})
		return
	}

	err := RegisterSong(c.Request.Context(), redisClient, postgresClient, req.URL)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Song successfully registered and downloaded"})
}

// single song download-->process-->fingerprints in redis and metaData in postgres
func RegisterSong(ctx context.Context, redisClient *redis.Client, postgresClient *gorm.DB, url string) error {
	//validate the url
	if !utils.IsValidYouTubeURL(url) {
		return errors.New("invalid url. please provide a valid youtube url")
	}

	// //downlaod the song
	// err := downloader.DownloadYTaudio(url, DOWNLOAD_PATH)
	// if err != nil {
	// 	return err
	// }

	log.Print("downloaded the video")

	//download the metadata of song
	data, err := downloader.GetVideoDetails(url)
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

	// converted the song to wav
	songPath, err := utils.FindDownloadedFile(songTitle, DOWNLOAD_PATH)
	if err != nil {
		return errors.New("song file can not be found")
	}
	wav.ConvertToWAV(songPath, 1) //stereo to mono audio and .wav format

	log.Print("converted the video to wav")

	// make wav into bytes
	wavPath := strings.TrimSuffix(songPath, filepath.Ext(songPath)) + ".wav"
	waveInfo, err := wav.ReadWavInfo(wavPath)
	if err != nil {
		log.Fatalf("error, %v", err)
	}

	// making wavbytes from samples
	samples, err := wav.WavBytesToSamples(waveInfo.Data)
	if err != nil {
		log.Fatalf("error converting wav bytes to float64: %v", err)
	}

	log.Print("converted to samples")
	// log.Print("erm what thw sigma")
	// log.Print(samples)

	//creating spectogram
	spectrogram, err := audiofingerprint.Spectrogram(samples, waveInfo.SampleRate)
	if err != nil {
		log.Fatalf("error creating spectrogram: %v", err)
	}
	log.Print("created the spectogram")

	// //viusalize the spectrogram in freq vs time. intensity based on db
	// magSpec, err := audiofingerprint.MagnitudeSpectrogram(spectrogram)
	// if err != nil {
	// 	log.Fatalf("error getting magnitudes of the spectrogram: %v", err)
	// }

	// output_path := "./downloads/spectrograms/viva_la_vida_spectrogram.png"
	// audiofingerprint.SaveSpectrogramImage(magSpec, output_path)

	// extract peaks ie most significant frequencies from each band
	peaks := audiofingerprint.ExtractPeaks(spectrogram, waveInfo.Duration)
	log.Print("extracted the peaks")

	// log.Print(peaks[:10])

	//create fingerprints
	fingerprints := audiofingerprint.CreateFingerprint(peaks, songID)
	log.Print("created the fingerprints")

	// save fingerprints to redis
	err = db.StoreFingerprints(ctx, redisClient, fingerprints)
	if err != nil {
		log.Fatalf("Failed to store fingerprints: %v", err)
	}
	log.Print("succesfully saved the fingerprints in redis")

	return nil
}
