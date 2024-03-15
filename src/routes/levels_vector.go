package routes

import (
	"codequest/src/controllers"

	"github.com/gin-gonic/gin"
)

func VectorLevelsRoutes(router *gin.Engine) {
	router.POST("/sort-vector", controllers.ExecuteVectorLevel())
}
