package routes

import (
	"codequest/controllers"

	"github.com/gin-gonic/gin"
)

func JsExecutorRoute(router *gin.Engine) {
	router.POST("/execute", controllers.JsExecute())
}
