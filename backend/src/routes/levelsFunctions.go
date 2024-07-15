package routes

import (
	"codequest/src/controllers"

	"github.com/gin-gonic/gin"
)

func FunctionsLevelsRoutes(router *gin.Engine) {
	router.POST("/functions/intro", controllers.FunctionsIntro())
	router.POST("/functions/ifelse", controllers.FunctionsIfElse())
	router.POST("/functions/object", controllers.FunctionsObject())
}
