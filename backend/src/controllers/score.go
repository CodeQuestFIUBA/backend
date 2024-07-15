package controllers

import (
	"context"
	"fmt"
	"strings"

	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"codequest/src/database"
	helper "codequest/src/helpers"

	"codequest/src/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const ScoreCollectionName = "scores"

func GetMyScore() gin.HandlerFunc {
	return func(c *gin.Context) {
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
		var classRoom = claims.ClassRoom
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

		var scoreRespoonses []models.ScoreResponse
		for _, score := range scores {
			var scoreResponse models.ScoreResponse
			scoreResponse.Level = score.SubLevelTitle
			scoreResponse.Complete = score.Points > 0
			scoreResponse.Score = score.Points
			if score.Points == 0 {
				scoreResponse.Qualification = 0
			} else {
				if score.Attempts == 0 {
					scoreResponse.Qualification = 1
				} else if score.Attempts == 1 {
					scoreResponse.Qualification = 2
				} else {
					scoreResponse.Qualification = 3
				}
			}
			scoreRespoonses = append(scoreRespoonses, scoreResponse)
		}

		pipeline := mongo.Pipeline{
			{{Key: "$match", Value: bson.D{{Key: "classRoom", Value: classRoom}}}},
			{{Key: "$lookup", Value: bson.D{
				{Key: "from", Value: "users"},
				{Key: "localField", Value: "user_id"},
				{Key: "foreignField", Value: "userid"},
				{Key: "as", Value: "user_details"},
			}}},
			{{Key: "$unwind", Value: bson.D{
				{Key: "path", Value: "$user_details"},
				{Key: "preserveNullAndEmptyArrays", Value: true},
			}}},
			{{Key: "$group", Value: bson.D{
				{Key: "_id", Value: "$user_details.userid"},
				{Key: "totalPoints", Value: bson.D{{Key: "$sum", Value: "$points"}}},
				{Key: "firstName", Value: bson.D{{Key: "$first", Value: "$user_details.firstname"}}},
				{Key: "lastName", Value: bson.D{{Key: "$first", Value: "$user_details.lastname"}}},
			}}},
			{{Key: "$project", Value: bson.D{
				{Key: "user_id", Value: "$_id"},
				{Key: "score", Value: "$totalPoints"},
				{Key: "user", Value: bson.D{
					{Key: "$concat", Value: bson.A{"$firstName", " ", "$lastName"}},
				}},
			}}},
			{{Key: "$sort", Value: bson.D{
				{Key: "score", Value: -1},
			}}},
		}

		var scoresByClassRoom []models.ScoreByClassRoom
		cursor2, err2 := scoreCollection.Aggregate(ctx, pipeline)
		if err2 != nil {
			log.Fatal(err)
		}

		var results []bson.M
		if err2 = cursor2.All(ctx, &results); err2 != nil {
			log.Fatal(err2)
		}

		for _, result := range results {
			var scoreByClassRoom models.ScoreByClassRoom
			scoreByClassRoom.User = result["user"].(string)
			scoreByClassRoom.Score = int(result["score"].(int32))
			userID := result["_id"].(string)
			scoreByClassRoom.MyUser = userID == userId

			scoresByClassRoom = append(scoresByClassRoom, scoreByClassRoom)
		}

		c.JSON(http.StatusOK, models.StandardResponse{
			Code:    http.StatusOK,
			Message: "Score updated successfully",
			Data:    gin.H{"scores": scoreRespoonses, "scoresByClassRoom": scoresByClassRoom},
		})

	}
}

func GetScoreByUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		var scoreCollection *mongo.Collection = database.OpenCollection(database.MongoClient, ScoreCollectionName)

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		var scores []models.Score
		var userID = c.Param("user_id")

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

		var totalScore int = 0
		for _, score := range scores {
			totalScore += score.Points
		}

		c.JSON(http.StatusOK, models.StandardResponse{
			Code:    http.StatusOK,
			Message: "Scores fetched successfully",
			Data:    gin.H{"scores": scores, "totalScore": totalScore},
		})

	}

}

func UpdateScore() gin.HandlerFunc {
	return func(c *gin.Context) {
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

		claims, msg := helper.ValidateAdminToken(strings.ReplaceAll(clientToken, "Bearer ", ""))

		if msg != "" {
			c.JSON(http.StatusBadRequest,
				models.StandardResponse{
					Code:    http.StatusBadRequest,
					Message: msg,
					Data:    nil})
			c.Abort()
			return
		}

		var userId = claims.Uid

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		var score models.Score

		var level = c.Param("level")
		var sublevel = c.Param("sublevel")

		errScore := scoreCollection.FindOne(ctx, bson.M{"level": level, "sub_level": sublevel, "user_id": userId}).Decode(&score)
		if errScore != nil {
			log.Panic(errScore)
			c.JSON(http.StatusInternalServerError, models.StandardResponse{
				Code:    http.StatusInternalServerError,
				Message: "error occurred while fetching the score",
				Data:    nil})
			return
		}

		var points = 0
		if score.Attempts == 0 {
			points = 100
		} else if score.Attempts == 1 {
			points = 50
		} else {
			points = 25
		}

		if score.Points > 0 {
			c.JSON(http.StatusOK, models.StandardResponse{
				Code:    http.StatusOK,
				Message: "Score already updated",
				Data:    gin.H{"updated": "no"},
			})
			return
		}

		var filter = bson.M{"level": level, "sub_level": sublevel, "user_id": userId}
		var update = bson.M{"$inc": bson.M{"points": points}}

		var _, err = scoreCollection.UpdateOne(ctx, filter, update)
		if err != nil {
			log.Panic(err)
			c.JSON(http.StatusInternalServerError, models.StandardResponse{
				Code:    http.StatusInternalServerError,
				Message: "error occurred while updating the score",
				Data:    nil})
			return
		}

		c.JSON(http.StatusOK, models.StandardResponse{
			Code:    http.StatusOK,
			Message: "Score updated successfully",
			Data:    gin.H{"updated": "ok"},
		})
	}
}

func UpdateScoreAttempts() gin.HandlerFunc {
	return func(c *gin.Context) {
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

		claims, msg := helper.ValidateAdminToken(strings.ReplaceAll(clientToken, "Bearer ", ""))

		if msg != "" {
			c.JSON(http.StatusBadRequest,
				models.StandardResponse{
					Code:    http.StatusBadRequest,
					Message: msg,
					Data:    nil})
			c.Abort()
			return
		}

		var userId = claims.Uid

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		var score models.Score

		var level = c.Param("level")
		var sublevel = c.Param("sublevel")

		errScore := scoreCollection.FindOne(ctx, bson.M{"level": level, "sub_level": sublevel, "user_id": userId}).Decode(&score)
		if errScore != nil {
			log.Panic(errScore)
			c.JSON(http.StatusInternalServerError, models.StandardResponse{
				Code:    http.StatusInternalServerError,
				Message: "error occurred while fetching the score",
				Data:    nil})
			return
		}

		log.Println("UserID: ", userId)
		log.Println("level: ", level)
		log.Println("sublevel: ", sublevel)
		log.Println("score: ", score)

		if score.Points > 0 {
			c.JSON(http.StatusOK, models.StandardResponse{
				Code:    http.StatusOK,
				Message: "Score already updated",
				Data:    gin.H{"updated": "no"},
			})
			return
		}

		var filter = bson.M{"level": level, "sub_level": sublevel, "user_id": userId}
		var update = bson.M{"$inc": bson.M{"attempts": 1}}

		var _, err = scoreCollection.UpdateOne(ctx, filter, update)
		if err != nil {
			log.Panic(err)
			c.JSON(http.StatusInternalServerError, models.StandardResponse{
				Code:    http.StatusInternalServerError,
				Message: "error occurred while updating the score",
				Data:    nil})
			return
		}

		c.JSON(http.StatusOK, models.StandardResponse{
			Code:    http.StatusOK,
			Message: "Score updated successfully",
			Data:    gin.H{"updated": "ok"},
		})
	}
}
