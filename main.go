package main

import (
	// // "context"
	// // // // "fmt"
	// "gozam/audiofingerprint"
	// // "gozam/db"

	// "gozam/wav"

	"log"

	// "strconv"

	"github.com/joho/godotenv"
)

// "gozam/downloader"
// "log"

func main() {

	//load env
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, continuing...")
	}

	// ctx := context.Background()

	// //new redis connection
	// redisClient, err := db.NewRedisClient()
	// if err != nil {
	// 	log.Fatal("failed to establish redis connection")
	// }
	// log.Print(redisClient)

	//<--------------1------------------------->
	//downloaded metadata and song
	//make custom function to name the song file based on metadata

	// url := "https://www.youtube.com/watch?v=dvgZkm1xWPE"
	// outputPath := "./downloads"

	// Get video metadata
	// data, err := downloader.GetVideoDetails(url)
	// if err != nil {
	// 	log.Fatalf("Error getting video details: %v", err)
	// }

	// songData := data.Items[0]
	// ytID := songData.ID
	// songTitle := songData.Snippet.Title
	// songArtist := songData.Snippet.ChannelTitle
	// songID := utils.GenerateUniqueID()
	// log.Print(songID, ytID, songTitle, songArtist)

	// // Download the audio
	// err := downloader.DownloadYTaudio(url, outputPath)
	// if err != nil {
	// 	log.Fatalf("Download failed: %v", err)
	// }

	//<--------------2------------------------->
	//converted the song to wav
	// song_path := "downloads/yellow coldplay.mp3"
	// channels := 1
	// wav.ConvertToWAV(song_path, channels)

	// // <--------------3------------------------->
	// // make wav into bytes
	// song_path := "downloads/yellow coldplay.wav"
	// waveInfo, err := wav.ReadWavInfo(song_path)
	// if err != nil {
	// 	log.Fatalf("error, %v", err)
	// }

	// // log.Print(waveInfo.SampleRate)

	// // <--------------4------------------------->
	// // making wavbytes from samples
	// samples, err := wav.WavBytesToSamples(waveInfo.Data)
	// if err != nil {
	// 	log.Fatalf("error converting wav bytes to float64: %v", err)
	// }

	// log.Print("erm what thw sigma")
	// // // log.Print(samples)

	// // <--------------5------------------------->
	// //creating spectogram
	// spectrogram, err := audiofingerprint.Spectrogram(samples, waveInfo.SampleRate)
	// if err != nil {
	// 	log.Fatalf("error creating spectrogram: %v", err)
	// }
	// log.Print(spectrogram)

	// // <--------------6------------------------->
	// //viusalize the spectrogram in freq vs time. intensity based on db
	// magSpec, err := audiofingerprint.MagnitudeSpectrogram(spectrogram)
	// if err != nil {
	// 	log.Fatalf("error getting magnitudes of the spectrogram: %v", err)
	// }

	// output_path := "./downloads/spectrograms/viva_la_vida_spectrogram.png"
	// audiofingerprint.SaveSpectrogramImage(magSpec, output_path)

	// // <--------------7------------------------->
	// // extract peaks ie most significant frequencies from each band
	// peaks := audiofingerprint.ExtractPeaks(spectrogram, waveInfo.Duration)
	// // log.Print(peaks[:10])

	// // <--------------8------------------------->
	// //create fingerprints
	// fingerprints := audiofingerprint.CreateFingerprint(peaks, songID)

	// log.Print(fingerprints)
	// count := 0
	// for k, v := range fingerprints {
	// 	log.Printf("%s: %d\n", k, v, "storing")
	// 	count++
	// 	if count >= 5 {
	// 		break
	// 	}
	// }

	//<---------------------8------------------------>
	//save fingerprints to redis

	// Store
	// err = db.StoreFingerprints(ctx, redisClient, fingerprints)
	// if err != nil {
	// 	log.Fatalf("Failed to store fingerprints: %v", err)
	// }

	// // Retrieve
	// retrieved, err := db.GetFingerprints(ctx, redisClient, fmt.Sprint(songID))
	// if err != nil {
	// 	log.Fatalf("Failed to retrieve fingerprints: %v", err)
	// }
	// // log.Print(retrieved)
	// count = 0
	// for k, v := range retrieved {
	// 	log.Printf("%s: %d\n", k, v, "retriving")
	// 	count++
	// 	if count >= 5 {
	// 		break
	// 	}
	// }

	//client side getting a clip of <10 sec and creating address:couple from them then from couples time and songid play the song

	//<---------------------9------------------------>
	// convert to wav
	// song_path := "./downloads/rickroll.mp3"
	// channels := 1
	// wav.ConvertToWAV(song_path, channels)

	// 	//<---------------------10------------------------>
	// make wav into bytes
	// song_path := "./downloads/viva la vida_cut.wav"
	// waveInfo, err := wav.ReadWavInfo(song_path)
	// if err != nil {
	// 	log.Fatalf("error, %v", err)
	// }

	// // <--------------11------------------------->
	// // making wavbytes from samples
	// samples, err := wav.WavBytesToSamples(waveInfo.Data)
	// if err != nil {
	// 	log.Fatalf("error converting wav bytes to float64: %v", err)
	// }

	// log.Print("erm what thw sigma: client side")
	// // log.Print(samples)

	// // <--------------11------------------------->
	// //creating spectogram
	// spectrogram, err := audiofingerprint.Spectrogram(samples, waveInfo.SampleRate)
	// if err != nil {
	// 	log.Fatalf("error creating spectrogram: %v", err)
	// }
	// // log.Print(spectrogram)

	// // <--------------12------------------------>
	// // extract peaks ie most significant frequencies from each band
	// peaks := audiofingerprint.ExtractPeaks(spectrogram, waveInfo.Duration)
	// // log.Print(peaks[:10])

	// // <--------------13------------------------->
	// //create fingerprints
	// TestsongID := utils.GenerateUniqueID()
	// log.Print(TestsongID)
	// samplefingerprint := audiofingerprint.CreateFingerprint(peaks, TestsongID)

	// sampleFingerprintMap := make(map[uint32]uint32)
	// for address, couple := range samplefingerprint {
	// 	sampleFingerprintMap[address] = couple.AnchorTimeMs
	// }
	// // log.Print(sampleFingerprintMap)

	// //match fingerprints
	// matches, timeDur, err := audiofingerprint.FindMatchesFGP(sampleFingerprintMap)
	// if err != nil {
	// 	log.Fatal("matching algo failed")
	// }
	// log.Print(matches, timeDur)

	// for _, match := range matches {
	// 	log.Print(match.Timestamp)
	// }

}
