package routes

import (
	"codequest/src/controllers"

	"github.com/gin-gonic/gin"
)

func ScoreRoute(router *gin.Engine) {
	router.GET("/score", controllers.GetMyScore())
	router.GET("/score/:user_id", controllers.GetScoreByUser())
	router.PUT("/score/complete/:level/:sublevel", controllers.UpdateScore())
	router.PUT("/score/attempts/:level/:sublevel", controllers.UpdateScoreAttempts())
}
