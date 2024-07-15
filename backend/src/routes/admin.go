package routes

import (
	"codequest/src/controllers"
	"github.com/gin-gonic/gin"
)

func AdminRoute(router *gin.Engine) {
	router.POST("/admin/signup", controllers.SignUpAdmin())
	router.POST("/admin/login", controllers.LoginAdmin())
	router.GET("/admin", controllers.GetAdmin())
}
