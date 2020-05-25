package controllers

import (
	"context"
	"log"

	mongodb "../db"
	models "../models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

// Login :handler for /api/login
func Login(c *gin.Context) {
	// get database
	db, err := mongodb.GetDatabase()
	if err != nil {
		log.Fatal(err)
		return
	}
	collection := db.Collection("users")

	// check if email exist in database
	var result models.UserLogin
	collection.FindOne(context.TODO(), bson.D{
		{"email", c.PostForm("email")},
	}).Decode(&result)

	// compare password hash
	compareErr:= bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(c.PostForm("password")))

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


