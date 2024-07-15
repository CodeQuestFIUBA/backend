package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"codequest/src/models"
)

func Ping() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.JSON(http.StatusOK, models.StandardResponse{Code: http.StatusOK, Message: "Request body executed",
			Data: "Pong!"})
	}
}
