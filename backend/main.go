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

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func main() {
	r := gin.Default()

	// Configurar CORS
	//r.Use(cors.New(cors.Config{
	//	AllowOrigins:     []string{"http://localhost:3006", "http://localhost:5173"},
	//	AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	//	AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
	//	AllowCredentials: true,
	//}))

	//r.Use(func(c *gin.Context) {
	//      c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	//      c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	//      c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
	//      c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")
	//
	//      if c.Request.Method == "OPTIONS" {
	//          c.AbortWithStatus(204)
	//          return
	//      }
	//
	//      c.Next()
	//  })

	r.Use(CORSMiddleware())

	_ = database.CreateMongoDBInstance()
	routes.UserRoute(r)
	routes.AdminRoute(r)

	r.Use(middleware.Authentication())

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"code": "PAGE_NOT_FOUND", "message": "404 page not found"})
	})

	routes.JsExecutorRoute(r)
	routes.IntroProgrammingLevelsRoutes(r)
	routes.FunctionsLevelsRoutes(r)
	routes.VectorLevelsRoutes(r)
	routes.MatrixLevelsRoutes(r)
	routes.ScoreRoute(r)
	routes.LevelRoute(r)
	routes.ClassRoomRoute(r)

	_ = r.Run(":8080")
}
