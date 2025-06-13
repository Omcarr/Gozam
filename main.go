package main

import (
	// "fmt"

	"gozam/audiofingerprint"
	"gozam/utils"
	"gozam/wav"
	"log"
)

// "gozam/downloader"
// "log"

func main() {
	//<--------------1------------------------->
	//downloaded metadata and song
	//make custom function to name the song file based on metadata

	// url := "https://www.youtube.com/watch?v=dvgZkm1xWPE"
	// outputPath := "./downloads"

	// // Get video metadata
	// // err := downloader.GetVideoDetails(url)
	// // if err != nil {
	// // 	log.Fatalf("Error getting video details: %v", err)
	// // }

	// // Download the audio
	// err := downloader.DownloadYTaudio(url, outputPath)
	// if err != nil {
	// 	log.Fatalf("Download failed: %v", err)
	// }

	//<--------------2------------------------->
	//converted the song to wav
	// song_path := "downloads/Coldplay - Viva La Vida (Official Video).m4a"
	// channels := 1
	// wav.ConvertToWAV(song_path, channels)

	//<--------------3------------------------->
	//make wav into bytes
	song_path := "downloads/Coldplay - Viva La Vida (Official Video).wav"
	waveInfo, err := wav.ReadWavInfo(song_path)
	if err != nil {
		log.Fatalf("error, %v", err)
	}

	// log.Print(waveInfo.SampleRate)

	// <--------------4------------------------->
	// making wavbytes from samples
	samples, err := wav.WavBytesToSamples(waveInfo.Data)
	if err != nil {
		log.Fatalf("error converting wav bytes to float64: %v", err)
	}

	log.Print("erm what thw sigma")
	// log.Print(samples)

	// <--------------5------------------------->
	//creating spectogram
	spectrogram, err := audiofingerprint.Spectrogram(samples, waveInfo.SampleRate)
	if err != nil {
		log.Fatalf("error creating spectrogram: %v", err)
	}
	// log.Print(spectrogram)

	// <--------------6------------------------->
	//viusalize the spectrogram in freq vs time. intensity based on db
	magSpec, err := audiofingerprint.MagnitudeSpectrogram(spectrogram)
	if err != nil {
		log.Fatalf("error getting magnitudes of the spectrogram: %v", err)
	}

	output_path := "./downloads/spectrograms/viva_la_vida_spectrogram.png"
	audiofingerprint.SaveSpectrogramImage(magSpec, output_path)

	// <--------------7------------------------->
	// extract peaks ie most significant frequencies from each band
	peaks := audiofingerprint.ExtractPeaks(spectrogram, waveInfo.Duration)
	log.Print(peaks[:10])

	// <--------------7------------------------->
	//create fingerprints
	songID := utils.GenerateUniqueID()
	fingerprints := audiofingerprint.CreateFingerprint(peaks, songID)

	// log.Print(fingerprints)
	// count := 0
	// for k, v := range fingerprints {
	// 	log.Printf("%s: %d\n", k, v)
	// 	count++
	// 	if count >= 10 {
	// 		break
	// 	}
	// }



	//<---------------------8------------------------>
	//save fingerprints to redis



}
