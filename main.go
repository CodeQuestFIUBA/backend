package main

import (
	"codequest/src/database"
	"codequest/src/middleware"
	"codequest/src/routes"

	_ "codequest/docs"

	"github.com/gin-gonic/gin"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

//	@title			CodeQuest API
//	@version		1.0
//	@description	Documentantion for CodeQuest API endpoints.

func main() {
	r := gin.Default()

	_ = database.CreateMongoDBInstance()
	routes.UserRoute(r)

	r.Use(middleware.Authentication())

	routes.JsExecutorRoute(r)
	routes.VectorLevelsRoutes(r)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	_ = r.Run(":8080")
}
