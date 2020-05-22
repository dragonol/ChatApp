package controllers

import (
	"context"
	"fmt"
	"log"

	mongodb "../db"
	models "../models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

//Login handler for /api/login
func Login(c *gin.Context) {
	db, err := mongodb.GetDatabase()

	if err != nil {
		log.Fatal(err)
		return
	}

	collection := db.Collection("users")
	var result models.UserLogin
	collection.FindOne(context.TODO(), bson.D{
		{"email", c.PostForm("email")},
		{"password", c.PostForm("password")},
	}).Decode(&result)

	fmt.Println(result)
	// if result != nil {
	// 	c.JSON(200, gin.H{
	// 		"status": "success",
	// 	})
	// } else {
	// 	c.JSON(200, gin.H{
	// 		"status": "fail",
	// 	})
	// }
}

func Register(c *gin.Context) {
	db, err := mongodb.GetDatabase()

	if err != nil {
		log.Fatal(err)
		return
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(c.PostForm("password")), 14)
	record := models.UserLogin{c.PostForm("email"), passwordHash}
	_, err = db.Collection("users").InsertOne(context.TODO(), record)

	if err != nil {
		log.Fatal(err)
	}
}
