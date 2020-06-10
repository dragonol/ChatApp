package controllers

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"log"

	database "../db"
	models "../models"

	"github.com/gin-gonic/gin"
)

func UpdateUserDetail(c *gin.Context) {
	db, err := database.GetDatabase()
	if err != nil {
		log.Fatal(err)
		return
	}
	collection := db.Collection("user_details")

	// check if email exist in database
	var result models.UserAuthenticate
	collection.FindOne(context.TODO(), bson.D{
		{Key: "email", Value: c.PostForm("email")},
	}).Decode(&result)
}
