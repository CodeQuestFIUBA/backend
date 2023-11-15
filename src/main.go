package main

import (
	"codequest/configs"
	"codequest/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	configs.ConnectToMongoDB()

	routes.UserRoute(r)

	r.Run(":8080")
}
