package routes

import (
	"codequest/src/controllers"

	"github.com/gin-gonic/gin"
)

func VectorLevelsRoutes(router *gin.Engine) {
	router.POST("/sort-vector", controllers.ExecuteVectorLevel())
	router.POST("/binary-search", controllers.ExecuteBinarySearchLevel())
	router.POST("/vectors/max", controllers.VectorMaxLevel())
	router.POST("/vectors/lineal", controllers.VectoLinealSearchLevel())
}
