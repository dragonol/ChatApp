package controllers

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"

	database "../db"
	models "../models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

// Register :handler for /api/login
func Register(c *gin.Context) {
	// user information
	userEmail := c.PostForm("email")
	userPassword := c.PostForm("password")
	userDisplayName := c.PostForm("displayName")

	// get database
	db, err := database.GetDatabase()
	if err != nil {
		log.Fatal(err)
		return
	}

	// get collections
	userAthenticate := db.Collection("user_authenticate")
	userDetails := db.Collection("user_details")

	// check if email exist in database
	var result models.UserAuthenticate
	userAthenticate.FindOne(context.TODO(), bson.D{
		{Key: "email", Value: c.PostForm("email")},
	}).Decode(&result)

	// check if no email found
	if !result.IsEmpty() {
		c.JSON(http.StatusOK, gin.H{
			"status": "fail",
		})
		return
	}

	// hash password
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(userPassword), 14)
	if err != nil {
		log.Fatal(err)
	}

	// create userAuthenticate record
	registerRecord := models.UserAuthenticate{
		Id:       primitive.NewObjectID(),
		Email:    userEmail,
		Password: passwordHash,
	}

	// insert to database
	userAuthenticateInsertResult, err := userAthenticate.InsertOne(context.TODO(), registerRecord)
	if err != nil {
		log.Fatal(err)
	}

	// create userDetails record
	registerDetailRecord := models.UserDetails{
		Id:          userAuthenticateInsertResult.InsertedID.(primitive.ObjectID),
		Email:       userEmail,
		DisplayName: userDisplayName,
	}

	_, err = userDetails.InsertOne(context.TODO(), registerDetailRecord)
	if err != nil {
		log.Fatal(err)
	}

	// response
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
	})
}
