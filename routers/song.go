package routers

import (
	"gozam/controllers"

	"github.com/gin-gonic/gin"
)

func SongRouter(router *gin.Engine) {
	router.GET("/song", controllers.SongHealthCheck)
}
