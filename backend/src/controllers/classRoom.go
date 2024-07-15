package controllers

import (
	"context"
	"strconv"

	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"

	"codequest/src/database"
	helper "codequest/src/helpers"

	"codequest/src/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const ClassRoomCollectionName = "classRoom"

func CreateClassRoom() gin.HandlerFunc {
	return func(c *gin.Context) {
		var classRoomCollection *mongo.Collection = database.OpenCollection(database.MongoClient, ClassRoomCollectionName)
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		clientToken := c.Request.Header.Get("Authorization")
		if clientToken == "" {
			c.JSON(http.StatusInternalServerError, models.StandardResponse{
				Code:    http.StatusBadRequest,
				Message: fmt.Sprintf("no authorization header provided"),
				Data:    nil})
			c.Abort()
			return
		}

		claims, _ := helper.ValidateAdminToken(strings.ReplaceAll(clientToken, "Bearer ", ""))

		var classRoom models.ClassRoom
		if err := c.BindJSON(&classRoom); err != nil {
			c.JSON(http.StatusBadRequest, models.StandardResponse{
				Code:    http.StatusBadRequest,
				Message: err.Error(),
				Data:    nil,
			})
			return
		}

		validationErr := validate.Struct(classRoom)
		if validationErr != nil {
			c.JSON(http.StatusBadRequest, models.StandardResponse{
				Code:    http.StatusBadRequest,
				Message: validationErr.Error(),
				Data:    nil,
			})
			return
		}

		newCode, err := helper.GenerateNewCode(ctx, classRoomCollection, classRoom.Code)

		if err != nil {
			log.Panic(err)
			c.JSON(http.StatusInternalServerError, models.StandardResponse{
				Code:    http.StatusInternalServerError,
				Message: "error occurred while generating new code",
				Data:    nil,
			})
			return
		}

		classRoom.Code = newCode
		classRoom.ID = primitive.NewObjectID()
		classRoom.Owner = claims.Uid
		classRoom.Users = []string{}
		classRoom.CreatedAt = time.Now()

		_, insertErr := classRoomCollection.InsertOne(ctx, classRoom)
		if insertErr != nil {
			msg := fmt.Sprintf("ClassRoom item was not created")
			c.JSON(http.StatusInternalServerError, models.StandardResponse{
				Code:    http.StatusInternalServerError,
				Message: msg,
				Data:    nil,
			})
			return
		}

		c.JSON(http.StatusOK, models.StandardResponse{
			Code:    http.StatusOK,
			Message: "ClassRoom created successfully",
			Data:    gin.H{"classRoom": classRoom},
		})

	}
}

func GetAllClassRoom() gin.HandlerFunc {
	return func(c *gin.Context) {
		var classRoomCollection *mongo.Collection = database.OpenCollection(database.MongoClient, ClassRoomCollectionName)
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

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

		page, err := strconv.Atoi(c.DefaultQuery("page", "0"))
		if err != nil || page < 1 {
			page = 0
		}
		limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))
		if err != nil || limit < 1 {
			limit = 10
		}

		skip := page * limit

		var classRooms []models.ClassRoomScores
		findOptions := options.Find()
		findOptions.SetSort(bson.D{{"created_at", -1}})
		findOptions.SetSkip(int64(skip))
		findOptions.SetLimit(int64(limit))

		count, err := classRoomCollection.CountDocuments(ctx, bson.M{"owner": userId})

		cursor, err := classRoomCollection.Find(ctx, bson.M{"owner": userId}, findOptions)
		if err != nil {
			log.Panic(err)
			c.JSON(http.StatusInternalServerError, models.StandardResponse{
				Code:    http.StatusInternalServerError,
				Message: "error occurred while fetching the classRoom",
				Data:    nil})
			return
		}
		defer cursor.Close(ctx)
		if err = cursor.All(ctx, &classRooms); err != nil {
			log.Panic(err)
			c.JSON(http.StatusInternalServerError, models.StandardResponse{
				Code:    http.StatusInternalServerError,
				Message: "error occurred while decoding the classRoom",
				Data:    nil})
			return
		}

		var total = 0
		for _, level := range models.Levels {
			total += len(level.SubLevels)
		}

		for i, classRoom := range classRooms {

			var totalScores int
			for _, user := range classRoom.Users {
				var scoreCollection *mongo.Collection = database.OpenCollection(database.MongoClient, "scores")
				var scores []models.Score
				cursor, err := scoreCollection.Find(ctx, bson.M{"user_id": user, "points": bson.M{"$gt": 0}}, options.Find())
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
				totalScores += len(scores)
			}
			classRooms[i].TotalScores = totalScores
			classRooms[i].TotalQuestions = total * len(classRoom.Users)
		}

		c.JSON(http.StatusOK, models.StandardResponse{
			Code:    http.StatusOK,
			Message: "ClassRooms fetched successfully",
			Data:    gin.H{"classRoom": classRooms, "total": count, "page": page, "limit": limit},
		})
	}
}

func GetUsersByClassRoom() gin.HandlerFunc {
	return func(c *gin.Context) {
		var classRoomCollection *mongo.Collection = database.OpenCollection(database.MongoClient, ClassRoomCollectionName)
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

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

		//Quiero transformar el id que recibo por params a un ObjectID
		classRoomId := c.Param("id")

		if classRoomId == "" {
			c.JSON(http.StatusBadRequest, models.StandardResponse{
				Code:    http.StatusBadRequest,
				Message: "classRoom id is required",
				Data:    nil,
			})
			return
		}

		var classRoom models.ClassRoom

		var objID, _ = primitive.ObjectIDFromHex(classRoomId)

		err := classRoomCollection.FindOne(ctx, bson.M{"_id": objID, "owner": userId}).Decode(&classRoom)
		if err != nil {
			fmt.Println("err", err)
			c.JSON(http.StatusInternalServerError, models.StandardResponse{
				Code:    http.StatusInternalServerError,
				Message: "error occurred while fetching the classRoom",
				Data:    nil,
			})
			return
		}

		var usersCollection *mongo.Collection = database.OpenCollection(database.MongoClient, "users")
		var users []models.UserScore

		var ids = []primitive.ObjectID{}
		for _, uId := range classRoom.Users {
			var uIdObjID, _ = primitive.ObjectIDFromHex(uId)
			ids = append(ids, uIdObjID)
		}

		cursor, err := usersCollection.Find(ctx, bson.M{"_id": bson.M{"$in": ids}}, options.Find())
		if err != nil {
			log.Panic(err)
			c.JSON(http.StatusInternalServerError, models.StandardResponse{
				Code:    http.StatusInternalServerError,
				Message: "error occurred while fetching the users",
				Data:    nil})
			return
		}
		defer cursor.Close(ctx)
		if err = cursor.All(ctx, &users); err != nil {
			log.Panic(err)
			c.JSON(http.StatusInternalServerError, models.StandardResponse{
				Code:    http.StatusInternalServerError,
				Message: "error occurred while decoding the users",
				Data:    nil})
			return
		}

		var total = 0
		for _, level := range models.Levels {
			total += len(level.SubLevels)
		}

		for i, user := range users {
			var scoreCollection *mongo.Collection = database.OpenCollection(database.MongoClient, "scores")
			var scores []models.Score
			cursor, err := scoreCollection.Find(ctx, bson.M{"user_id": user.UserId, "points": bson.M{"$gt": 0}}, options.Find())
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

			var total_scores = 0
			for _, sc := range scores {
				total_scores += sc.Points
			}

			users[i].CompletedLevels = len(scores)
			users[i].TotalLevels = total
			users[i].Score = total_scores
		}

		c.JSON(http.StatusOK, models.StandardResponse{
			Code:    http.StatusOK,
			Message: "ClassRoom fetched successfully",
			Data:    gin.H{"users": users},
		})
	}
}
