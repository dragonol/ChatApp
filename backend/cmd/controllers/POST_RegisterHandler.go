package controllers

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"time"

	database "../db"
	models "../models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

// Register :handler for /api/login
func Register(c *gin.Context) {
	// user information
	userEmail := c.PostForm("email")
	userPassword := c.PostForm("password")
	userDisplayName := c.PostForm("displayName")
	userDoB, err := time.Parse("20060102", c.PostForm("dob"))
	if err != nil {
		log.Fatal(err)
		return
	}
	userPhoneNumber := c.PostForm("phoneNumber")

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
		c.JSON(200, gin.H{
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
		Password:    passwordHash,
		PhoneNumber: userPhoneNumber,
		DisplayName: userDisplayName,
		DateOfBirth: userDoB,
	}

	_, err = userDetails.InsertOne(context.TODO(), registerDetailRecord)
	if err != nil {
		log.Fatal(err)
	}

	// response
	c.JSON(200, gin.H{
		"status": "success",
	})
}
