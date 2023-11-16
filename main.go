package main

import (
	"codequest/src/configs"
	"codequest/src/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	_ = configs.ConnectToMongoDB()

	routes.UserRoute(r)
	routes.JsExecutorRoute(r)

	_ = r.Run(":8080")
}
