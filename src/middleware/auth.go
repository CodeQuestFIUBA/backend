package middleware

import (
	"fmt"
	"net/http"
	"strings"

	helper "codequest/src/helpers"
	"codequest/src/models"

	"github.com/gin-gonic/gin"
)

// Validates token and authorizes users
func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		clientToken := c.Request.Header.Get("Authorization")
		if clientToken == "" {
			c.JSON(http.StatusInternalServerError, models.StandardResponse{
				Code:    http.StatusBadRequest,
				Message: fmt.Sprintf("no authorization header provided"),
				Data:    nil})
			c.Abort()
			return
		}

		claims, msg := helper.ValidateToken(strings.ReplaceAll(clientToken, "Bearer ", ""))

		if msg != "" {
			c.JSON(http.StatusBadRequest,
				models.StandardResponse{
					Code:    http.StatusBadRequest,
					Message: msg,
					Data:    nil})
			c.Abort()
			return
		}

		c.Set("email", claims.Email)
		c.Set("first_name", claims.FirstName)
		c.Set("last_name", claims.LastName)
		c.Set("uid", claims.Uid)

		c.Next()
	}
}
