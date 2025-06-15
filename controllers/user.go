package controllers

import "github.com/gin-gonic/gin"

func HomePage(cxt *gin.Context) {
	cxt.JSON(200, gin.H{
		"message": "welcome to gozam",
	})
}
