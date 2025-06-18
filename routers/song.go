package routers

import (
	"gozam/controllers"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

func SongRouter(router *gin.Engine, redisClient *redis.Client, db *gorm.DB) {
	router.GET("/song", controllers.SongHealthCheck)
	router.POST("/register_song", func(c *gin.Context) {
		controllers.RegisterSongHandler(c, redisClient, db)
	})
}
