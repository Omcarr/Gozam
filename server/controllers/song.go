package controllers

import (
	"fmt"
	"gozam/downloader"
	"gozam/utils"
	// "log"
	"net/http"

	"os"
	"path/filepath"
	// "strings"
	"time"

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

	urlType := utils.GetYouTubeURLType(req.URL)

	if urlType == "video" {
		err := downloader.RegisterSong(c.Request.Context(), redisClient, postgresClient, req.URL)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	} else if urlType == "playlist" {
		err := downloader.RegisterPlaylist(c.Request.Context(), redisClient, postgresClient, req.URL)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing or invalid URL in request body"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Song successfully registered and downloaded"})
}

func FindMatchesHandler(c *gin.Context, postgresClient *gorm.DB) {
	// 1. Parse uploaded file
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "file is required"})
		return
	}

	// Ensure the tmp directory exists
	if err := os.MkdirAll("tmp", os.ModePerm); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create tmp directory"})
		return
	}

	// Save file in local tmp directory with unique name
	tmpPath := filepath.Join("tmp", fmt.Sprintf("%d_%s", time.Now().UnixNano(), file.Filename))
	if err := c.SaveUploadedFile(file, tmpPath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to save uploaded file"})
		return
	}

	//find matches
	matches, err := downloader.FindClientMatches(tmpPath, postgresClient)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// wavPath := strings.TrimSuffix(tmpPath, filepath.Ext(tmpPath)) + ".wav"
	// // Now cleanup:
	// if err := os.Remove(tmpPath); err != nil {
	// 	log.Printf("Failed to delete uploaded file: %s", err)
	// }

	// if err := os.Remove(wavPath); err != nil {
	// 	log.Printf("Failed to delete wav file: %s", err)
	// }
	c.JSON(http.StatusOK, gin.H{"matches": matches})
}
