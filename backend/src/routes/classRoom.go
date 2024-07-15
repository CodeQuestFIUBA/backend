package routes

import (
	"codequest/src/controllers"

	"github.com/gin-gonic/gin"
)

func ClassRoomRoute(router *gin.Engine) {
	router.GET("/classroom/all", controllers.GetAllClassRoom())
	router.POST("/classroom", controllers.CreateClassRoom())
	router.GET("/classroom/:id/users", controllers.GetUsersByClassRoom())
}
