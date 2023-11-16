package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"rogchap.com/v8go"
)

func JsExecute() gin.HandlerFunc {
	return func(c *gin.Context) {

		body, err := c.GetRawData()

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read request body"})
			return
		}

		fmt.Printf("Request Body:\n\n%s\n\n", body)

		ctx := v8go.NewContext()
		_, err = ctx.RunScript(string(body), "main.js")

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed executing de JavaScript code -> " + err.Error()})
			return
		}

		val, _ := ctx.RunScript("result", "value.js")

		fmt.Printf("Value: %s\n\n", val)

		c.JSON(http.StatusOK, gin.H{"message": "Request body received", "value": val})
	}
}
