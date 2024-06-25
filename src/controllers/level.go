package controllers

import (
	helper "codequest/src/helpers"
	"codequest/src/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllLevels(c *gin.Context) {
	responseLevels := helper.ConvertLevelsToResponseFormat()

	c.JSON(http.StatusOK, models.StandardResponse{
		Code:    http.StatusOK,
		Message: "Scores fetched successfully",
		Data:    responseLevels,
	})
}
