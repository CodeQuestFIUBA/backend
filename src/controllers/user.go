package controllers

import (
	"codequest/configs"
	"codequest/models"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const UsersCollectionName = "users"

func GetAllUsers() gin.HandlerFunc {
	return func(c *gin.Context) {
		cursor, err := configs.MongoClient.Database(configs.EnvDBName()).
			Collection(UsersCollectionName).
			Find(context.TODO(), bson.D{{}})

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		var users []bson.M
		if err = cursor.All(context.TODO(), &users); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, users)
	}
}

func GetUserByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := primitive.ObjectIDFromHex(idStr)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var user bson.M
		err = configs.MongoClient.Database(configs.EnvDBName()).
			Collection(UsersCollectionName).
			FindOne(context.TODO(), bson.D{{Key: "_id", Value: id}}).
			Decode(&user)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, user)
	}
}

func PostUser() gin.HandlerFunc {
	return func(c *gin.Context) {

		var pipeline models.User
		if err := c.BindJSON(&pipeline); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		result, err := configs.MongoClient.Database(configs.EnvDBName()).
			Collection(UsersCollectionName).
			InsertOne(context.TODO(), pipeline)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, result)
	}
}
