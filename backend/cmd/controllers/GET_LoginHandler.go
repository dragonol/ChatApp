package controllers

import (
	"context"
	"fmt"
	"log"

	database "../db"
	"../models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

// Login :handler for /api/login
func Login(c *gin.Context) {
	// email and password
	userEmail := c.DefaultQuery("email", "")
	userPassword := c.DefaultQuery("password", "")
	// get database
	db, err := database.GetDatabase()
	if err != nil {
		log.Fatal(err)
		return
	}
	// check if email exist in database
	collection := db.Collection("user_authenticate")
	var result models.UserAuthenticate

	err = collection.FindOne(context.TODO(), bson.M{
		"email": userEmail,
	}).Decode(&result)

	fmt.Println(result)

	// compare password hash
	compareErr := bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(userPassword))

	// response
	if compareErr != nil {
		c.JSON(200, gin.H{
			"status": "fail",
		})
	} else {
		c.JSON(200, gin.H{
			"status": "success",
			"userId": result.Id,
		})
	}
}
