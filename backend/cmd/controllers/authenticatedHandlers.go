package controllers

import (
	mongodb "../db"
	"context"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"log"
)

//Login handler for /api/login
func Login(c *gin.Context) {
	db, err := mongodb.GetDatabase()

	if err != nil {
		log.Fatal(err)
		return
	}

	collection := db.Collection("users")
	var result bson.D
	collection.FindOne(context.TODO(), bson.D{
		{"email", c.PostForm("email")},
		{"password", c.PostForm("password")},
	}).Decode(&result)

	if result != nil {
		c.JSON(200, gin.H{
			"status": "success",
		})
	} else {
		c.JSON(200, gin.H{
			"status": "fail",
		})
	}
}

func Register(c *gin.Context) {
	db, err := mongodb.GetDatabase()

	if err != nil {
		log.Fatal(err)
		return
	}
	
}
