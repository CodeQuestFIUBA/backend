package routes

import (
	"codequest/src/controllers"

	"github.com/gin-gonic/gin"
)

func MatrixLevelsRoutes(router *gin.Engine) {
	router.POST("/search-matrix", controllers.ExecuteSearchMatrixLevel())
}
