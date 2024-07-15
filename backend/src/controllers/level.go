package controllers

import (
	"codequest/src/database"
	helper "codequest/src/helpers"
	"codequest/src/models"
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetAllLevels(c *gin.Context) {
	responseLevels := helper.ConvertLevelsToResponseFormat()

	c.JSON(http.StatusOK, models.StandardResponse{
		Code:    http.StatusOK,
		Message: "Scores fetched successfully",
		Data:    responseLevels,
	})
}

func GetMyActualLevel(c *gin.Context) {

	var scoreCollection *mongo.Collection = database.OpenCollection(database.MongoClient, ScoreCollectionName)

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

	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	if msg != "" {
		c.JSON(http.StatusBadRequest,
			models.StandardResponse{
				Code:    http.StatusBadRequest,
				Message: msg,
				Data:    nil})
		c.Abort()
		return
	}
	var scores []models.Score
	var userId = claims.Uid

	cursor, err := scoreCollection.Find(ctx, bson.M{"user_id": userId}, options.Find())
	if err != nil {
		log.Panic(err)
		c.JSON(http.StatusInternalServerError, models.StandardResponse{
			Code:    http.StatusInternalServerError,
			Message: "error occurred while fetching the scores",
			Data:    nil})
		return
	}
	defer cursor.Close(ctx)
	if err = cursor.All(ctx, &scores); err != nil {
		log.Panic(err)
		c.JSON(http.StatusInternalServerError, models.StandardResponse{
			Code:    http.StatusInternalServerError,
			Message: "error occurred while decoding the scores",
			Data:    nil})
		return
	}

	var levelKeys []models.LevelKeys
	for _, score := range scores {
		levelKeys = append(levelKeys, models.LevelKeys{
			Level:     score.Level,
			SubLevel:  score.SubLevel,
			Completed: score.Points > 0,
		})
	}

	c.JSON(http.StatusOK, models.StandardResponse{
		Code:    http.StatusOK,
		Message: "Scores fetched successfully",
		Data:    gin.H{"levels": levelKeys},
	})
}
