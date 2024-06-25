package main

import (
	"codequest/src/database"
	middleware "codequest/src/middlewares"
	"codequest/src/routes"

	"net/http"

	"github.com/gin-gonic/gin"
)

//	@title			CodeQuest API
//	@version		1.0
//	@description	Documentantion for CodeQuest API endpoints.

func main() {
	r := gin.Default()

	_ = database.CreateMongoDBInstance()
	routes.UserRoute(r)

	r.Use(middleware.Authentication())

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"code": "PAGE_NOT_FOUND", "message": "404 page not found"})
	})

	routes.JsExecutorRoute(r)
	routes.VectorLevelsRoutes(r)
	routes.MatrixLevelsRoutes(r)
	routes.ScoreRoute(r)
	routes.LevelRoute(r)

	_ = r.Run(":8080")
}
