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

// Register :handler for /api/login
func Register(c *gin.Context) {
	// get database
	db, err := mongodb.GetDatabase()
	if err != nil {
		log.Fatal(err)
		return
	}
	collection:= db.Collection("users")

	// check if email exist in database
	var result models.UserLogin
	collection.FindOne(context.TODO(), bson.D{
		{"email", c.PostForm("email")},
	}).Decode(&result)

	if !result.IsEmpty() {
		c.JSON(200, gin.H{
			"status": "fail",
		})
		return
	}

	// hash password
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(c.PostForm("password")), 14)
	if err != nil {
		log.Fatal(err)
	}
	record := models.UserLogin{c.PostForm("email"), passwordHash}

	// insert to database
	_, err = collection.InsertOne(context.TODO(), record)
	if err != nil {
		log.Fatal(err)
	}

	// response
	c.JSON(200, gin.H{
		"status": "success",
	})
}