package routes

import (
	"codequest/src/controllers"

	"github.com/gin-gonic/gin"
)

func IntroProgrammingLevelsRoutes(router *gin.Engine) {
	router.POST("/intro/functions", controllers.IntroFunctions())
	router.POST("/intro/ifElse", controllers.IntroIfElse())
	router.POST("/intro/for", controllers.IntroFor())
	router.POST("/intro/while", controllers.IntroWhile())
}
