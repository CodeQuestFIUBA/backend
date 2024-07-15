package routes

import (
	"codequest/src/controllers"
	"github.com/gin-gonic/gin"
)

func LevelRoute(router *gin.Engine) {
	router.GET("/levels", controllers.GetAllLevels)
	router.GET("/levels/actual", controllers.GetMyActualLevel)
}