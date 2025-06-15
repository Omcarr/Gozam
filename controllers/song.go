package controllers

import (
	"errors"
	"gozam/downloader"
	"gozam/utils"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

var DOWNLOAD_PATH = "./downloads"

func SongHealthCheck(cxt *gin.Context) {
	cxt.JSON(200, gin.H{
		"message": "apdi pode",
	})
}

func RegisterSongHandler(c *gin.Context) {
	type RequestBody struct {
		URL string `json:"url" binding:"required"`
	}

	var req RequestBody
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing or invalid URL in request body"})
		return
	}

	err := RegisterSong(req.URL)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Song successfully registered and downloaded"})
}

// single song download-->process-->fingerprints in redis and metaData in postgres
func RegisterSong(url string) error {

	//validate the url
	if !utils.IsValidYouTubeURL(url) {
		return errors.New("invalid url. please provide a valid youtube url")
	}

	//downlaod the song
	err := downloader.DownloadYTaudio(url, DOWNLOAD_PATH)
	if err != nil {
		return err
	}

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

	return nil
}
