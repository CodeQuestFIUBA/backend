package routes

import (
	"codequest/src/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine) {
	router.GET("/users", controllers.GetAllUsers())
	router.GET("/users/:id", controllers.GetUserByID())
	router.POST("/users", controllers.PostUser())
}
