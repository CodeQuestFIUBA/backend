package controllers

import (
	"context"

	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"codequest/src/database"

	"codequest/src/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const ScoreCollectionName = "scores"

// quiero obtener todas las colecciones del usuario tal que tenga user_id pasado en params
func GetScoreByUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		var scoreCollection *mongo.Collection = database.OpenCollection(database.MongoClient, ScoreCollectionName)

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		var scores []models.Score
		var userID = c.Param("user_id")

		log.Println("Error inserting scores:", userID)

		log.Default().Println(userID)
		cursor, err := scoreCollection.Find(ctx, bson.M{"user_id": userID}, options.Find())
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
		log.Println("Error inserting scores:", scores)
		// quiero obtener la suma de todos los scores.points en GO
		var totalScore int = 0
		for _, score := range scores {
			totalScore += score.Points
		}

		//quiero retornar en Data los scores y el totalScore
		c.JSON(http.StatusOK, models.StandardResponse{
			Code:    http.StatusOK,
			Message: "Scores fetched successfully",
			Data:    gin.H{"scores": scores, "totalScore": totalScore},
		})
	}
}

func InsertScore() gin.HandlerFunc {
	return func(c *gin.Context) {
		var scoreCollection *mongo.Collection = database.OpenCollection(database.MongoClient, ScoreCollectionName)

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		var score models.Score

		if err := c.BindJSON(&score); err != nil {
			c.JSON(http.StatusBadRequest, models.StandardResponse{
				Code:    http.StatusBadRequest,
				Message: err.Error(),
				Data:    nil})
			return
		}

		score.ID = primitive.NewObjectID()

		validationErr := validate.Struct(score)
		if validationErr != nil {
			c.JSON(http.StatusBadRequest, models.StandardResponse{
				Code:    http.StatusBadRequest,
				Message: validationErr.Error(),
				Data:    nil})
			return
		}

		_, insertErr := scoreCollection.InsertOne(ctx, score)
		if insertErr != nil {
			log.Panic(insertErr)
			c.JSON(http.StatusInternalServerError, models.StandardResponse{
				Code:    http.StatusInternalServerError,
				Message: "error occurred while inserting the score",
				Data:    nil})
			return
		}

		c.JSON(http.StatusOK, models.StandardResponse{
			Code:    http.StatusOK,
			Message: "Score inserted successfully",
			Data:    score})
	}
}
