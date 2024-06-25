package routes

import (
	"codequest/src/controllers"

	"github.com/gin-gonic/gin"
)

func ScoreRoute(router *gin.Engine) {
    router.GET("/score/:user_id", controllers.GetScoreByUser())
   // router.POST("/score", controllers.InsertScore())
}