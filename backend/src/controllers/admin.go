package controllers

import (
	"context"
	"strings"

	"fmt"
	"log"
	"net/http"
	"time"

	"codequest/src/database"

	"github.com/gin-gonic/gin"

	helper "codequest/src/helpers"
	"codequest/src/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

const AdminsCollectionName = "admin"

func SignUpAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		var adminCollection *mongo.Collection = database.OpenCollection(database.MongoClient, AdminsCollectionName)
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		var admin models.Admin

		if err := c.BindJSON(&admin); err != nil {
			c.JSON(http.StatusBadRequest, models.StandardResponse{
				Code:    http.StatusBadRequest,
				Message: err.Error(),
				Data:    nil,
			})
			return
		}

		validationErr := validate.Struct(admin)
		if validationErr != nil {
			c.JSON(http.StatusBadRequest, models.StandardResponse{
				Code:    http.StatusBadRequest,
				Message: validationErr.Error(),
				Data:    nil,
			})
			return
		}

		count, err := adminCollection.CountDocuments(ctx, bson.M{"email": admin.Email})
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

		hashedPassword, err := helper.HashAdminPassword(*admin.Password, bcrypt.DefaultCost)
		if err != nil {
			c.JSON(http.StatusInternalServerError, models.StandardResponse{
				Code:    http.StatusInternalServerError,
				Message: "failed to hash password",
				Data:    nil,
			})
			return
		}
		hashedPasswordString := string(hashedPassword)
		admin.Password = &hashedPasswordString

		admin.ID = primitive.NewObjectID()
		admin.AdminId = admin.ID.Hex()
		token, refreshToken, _ := helper.GenerateAllAdminTokens(*admin.Email, *admin.Name, admin.AdminId)

		_, insertErr := adminCollection.InsertOne(ctx, admin)
		if insertErr != nil {
			msg := fmt.Sprintf("admin item was not created")
			c.JSON(http.StatusInternalServerError, models.StandardResponse{
				Code:    http.StatusInternalServerError,
				Message: msg,
				Data:    nil,
			})
			return
		}

		c.JSON(http.StatusOK, models.StandardResponse{
			Code:    http.StatusOK,
			Message: "Admin created successfully",
			Data:    models.TokenAdminResponse{Admin: admin, Token: &token, RefreshToken: &refreshToken},
		})
	}
}

func LoginAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		var adminCollection *mongo.Collection = database.OpenCollection(database.MongoClient, AdminsCollectionName)

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var admin models.Admin
		var foundAdmin models.Admin

		if err := c.BindJSON(&admin); err != nil {
			c.JSON(http.StatusBadRequest, models.StandardResponse{
				Code:    http.StatusBadRequest,
				Message: err.Error(),
				Data:    nil})
			return
		}

		err := adminCollection.FindOne(ctx, bson.M{"email": admin.Email}).Decode(&foundAdmin)
		defer cancel()
		if err != nil {
			c.JSON(http.StatusBadRequest,
				models.StandardResponse{
					Code:    http.StatusBadRequest,
					Message: "email does not exist",
					Data:    nil})
			return
		}

		err = helper.VerifyAdminPassword(*admin.Password, *foundAdmin.Password)
		defer cancel()
		if err != nil {
			c.JSON(http.StatusBadRequest,
				models.StandardResponse{
					Code:    http.StatusBadRequest,
					Message: err.Error(),
					Data:    nil})
			return
		}

		token, refreshToken, _ := helper.GenerateAllAdminTokens(*foundAdmin.Email, *foundAdmin.Name, foundAdmin.AdminId)

		c.JSON(http.StatusOK,
			models.StandardResponse{
				Code:    http.StatusOK,
				Message: "",
				Data:    models.TokenAdminResponse{Admin: foundAdmin, Token: &token, RefreshToken: &refreshToken},
			})
	}
}

func GetAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		var adminCollection *mongo.Collection = database.OpenCollection(database.MongoClient, AdminsCollectionName)
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

		var admin models.Admin
		err := adminCollection.FindOne(ctx, bson.M{"adminid": claims.Uid}).Decode(&admin)
		if err != nil {
			c.JSON(http.StatusBadRequest, models.StandardResponse{
				Code:    http.StatusBadRequest,
				Message: "admin does not exist",
				Data:    nil})
			return
		}

		c.JSON(http.StatusOK, models.StandardResponse{
			Code:    http.StatusOK,
			Message: "",
			Data:    admin,
		})
	}
}
