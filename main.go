package main

import (
	"codequest/src/configs"
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

	_ = configs.ConnectToMongoDB()

	routes.UserRoute(r)
	routes.JsExecutorRoute(r)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	_ = r.Run(":8080")
}
