package controllers

import (
	"context"
	"log"

	database "../db"
	"../models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

// Login :handler for /api/login
func Login(c *gin.Context) {
	// get database
	db, err := database.GetDatabase()
	if err != nil {
		log.Fatal(err)
		return
	}
	// check if email exist in database
	collection := db.Collection("users")
	var result models.UserLogin
	err = collection.FindOne(context.TODO(), bson.D{
		{Key: "email", Value: c.PostForm("email")},
	}).Decode(&result)

	// compare password hash
	compareErr := bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(c.PostForm("password")))

	// response
	if compareErr != nil {
		c.JSON(200, gin.H{
			"status": "fail",
		})
	} else {
		c.JSON(200, gin.H{
			"status": "success",
		})
	}
}
