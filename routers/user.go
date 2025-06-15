package routers

import (
	"gozam/controllers"

	"github.com/gin-gonic/gin"
)

func UserRouter(router *gin.Engine) {
	router.GET("/", controllers.HomePage)
}
