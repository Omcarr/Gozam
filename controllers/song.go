package controllers

import "github.com/gin-gonic/gin"

func SongHealthCheck(cxt *gin.Context) {
	cxt.JSON(200, gin.H{
		"message": "apdi pode",
	})
}
