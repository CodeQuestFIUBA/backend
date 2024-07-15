package controllers

import (
	"context"

	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	"codequest/src/database"

	helper "codequest/src/helpers"
	"codequest/src/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

const UsersCollectionName = "users"

var validate = validator.New()

func SignUp() gin.HandlerFunc {
	return func(c *gin.Context) {
		var userCollection *mongo.Collection = database.OpenCollection(database.MongoClient, "users")
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		var user models.User

		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, models.StandardResponse{
				Code:    http.StatusBadRequest,
				Message: err.Error(),
				Data:    nil,
			})
			return
		}

		validationErr := validate.Struct(user)
		if validationErr != nil {
			c.JSON(http.StatusBadRequest, models.StandardResponse{
				Code:    http.StatusBadRequest,
				Message: validationErr.Error(),
				Data:    nil,
			})
			return
		}

		count, err := userCollection.CountDocuments(ctx, bson.M{"email": user.Email})
		if err != nil {
			log.Panic(err)
			c.JSON(http.StatusInternalServerError, models.StandardResponse{
				Code:    http.StatusInternalServerError,
				Message: "error occurred while checking for the email",
				Data:    nil,
			})
			return
		}

		if count > 0 {
			c.JSON(http.StatusBadRequest, models.StandardResponse{
				Code:    http.StatusBadRequest,
				Message: "the email is already registered",
				Data:    nil,
			})
			return
		}

		hashedPassword, err := helper.HashPassword(*user.Password, bcrypt.DefaultCost)
		if err != nil {
			c.JSON(http.StatusInternalServerError, models.StandardResponse{
				Code:    http.StatusInternalServerError,
				Message: "failed to hash password",
				Data:    nil,
			})
			return
		}
		hashedPasswordString := string(hashedPassword)
		user.Password = &hashedPasswordString

		count, err = userCollection.CountDocuments(ctx, bson.M{"username": user.Username})
		if err != nil {
			log.Panic(err)
			c.JSON(http.StatusInternalServerError, models.StandardResponse{
				Code:    http.StatusInternalServerError,
				Message: "error occurred while checking for the username",
				Data:    nil,
			})
			return
		}

		if count > 0 {
			c.JSON(http.StatusBadRequest, models.StandardResponse{
				Code:    http.StatusBadRequest,
				Message: "the username is already registered",
				Data:    nil,
			})
			return
		}

		user.CreatedAt = time.Now()
		user.UpdatedAt = time.Now()
		user.ID = primitive.NewObjectID()
		user.UserId = user.ID.Hex()
		token, refreshToken, _ := helper.GenerateAllTokens(*user.Email, *user.Username, *user.FirstName, *user.LastName, user.UserId, user.ClassRoom)

		_, insertErr := userCollection.InsertOne(ctx, user)
		if insertErr != nil {
			msg := fmt.Sprintf("User item was not created")
			c.JSON(http.StatusInternalServerError, models.StandardResponse{
				Code:    http.StatusInternalServerError,
				Message: msg,
				Data:    nil,
			})
			return
		}

		if user.ClassRoom != "" {
			var classRoomCollection *mongo.Collection = database.OpenCollection(database.MongoClient, "classRoom")
			var classRoom models.ClassRoom
			err := classRoomCollection.FindOne(ctx, bson.M{"code": user.ClassRoom}).Decode(&classRoom)
			if err != nil {
				c.JSON(http.StatusBadRequest, models.StandardResponse{
					Code:    http.StatusBadRequest,
					Message: "classRoom does not exist",
					Data:    nil,
				})
				return
			}
			_, err = classRoomCollection.UpdateOne(ctx, bson.M{"code": user.ClassRoom}, bson.M{"$push": bson.M{"users": user.UserId}})
			if err != nil {
				c.JSON(http.StatusInternalServerError, models.StandardResponse{
					Code:    http.StatusInternalServerError,
					Message: "failed to add user to classRoom",
					Data:    nil,
				})
				return
			}
		}

		room := user.ClassRoom
		_, err = helper.GenerateScoresForUserAndRoom(user.UserId, room)
		if err != nil {
			c.JSON(http.StatusInternalServerError, models.StandardResponse{
				Code:    http.StatusInternalServerError,
				Message: "failed to generate scores",
				Data:    nil,
			})
			return
		}

		c.JSON(http.StatusOK, models.StandardResponse{
			Code:    http.StatusOK,
			Message: "user created successfully",
			Data:    models.TokenResponse{User: user, Token: &token, RefreshToken: &refreshToken},
		})
	}
}

// Login is the api used to tget a single user
func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		var userCollection *mongo.Collection = database.OpenCollection(database.MongoClient, UsersCollectionName)

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var user models.User
		var foundUser models.User

		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, models.StandardResponse{
				Code:    http.StatusBadRequest,
				Message: err.Error(),
				Data:    nil})
			return
		}

		err := userCollection.FindOne(ctx, bson.M{"email": user.Email}).Decode(&foundUser)
		defer cancel()
		if err != nil {
			c.JSON(http.StatusBadRequest,
				models.StandardResponse{
					Code:    http.StatusBadRequest,
					Message: "email does not exist",
					Data:    nil})
			return
		}

		err = helper.VerifyPassword(*user.Password, *foundUser.Password)
		defer cancel()
		if err != nil {
			c.JSON(http.StatusBadRequest,
				models.StandardResponse{
					Code:    http.StatusBadRequest,
					Message: err.Error(),
					Data:    nil})
			return
		}

		token, refreshToken, _ := helper.GenerateAllTokens(*foundUser.Email, *foundUser.Username, *foundUser.FirstName, *foundUser.LastName, foundUser.UserId, foundUser.ClassRoom)

		c.JSON(http.StatusOK,
			models.StandardResponse{
				Code:    http.StatusOK,
				Message: "",
				Data:    models.TokenResponse{User: foundUser, Token: &token, RefreshToken: &refreshToken},
			})
	}
}
