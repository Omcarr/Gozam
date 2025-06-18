package controllers

import (
	"gozam/downloader"
	"net/http"

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

	err := downloader.RegisterSong(c.Request.Context(), redisClient, postgresClient, req.URL)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Song successfully registered and downloaded"})
}
