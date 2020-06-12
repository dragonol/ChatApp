package controllers

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"net/http"

	database "../db"
	"github.com/gin-gonic/gin"
)

func UpdateUserDetail(c *gin.Context) {
	// user information
	userEmail := c.PostForm("email")
	userDisplayName := c.PostForm("displayName")

	// get database
	db, err := database.GetDatabase()
	if err != nil {
		log.Fatal(err)
		return
	}

	// get collection
	userDetails := db.Collection("user_details")

	// update database
	_, err = userDetails.UpdateOne(context.TODO(),
		bson.D{
			{Key: "email", Value: userEmail},
		},
		bson.D{{"$set", bson.D{
			{Key: "displayname", Value: userDisplayName},
		}}})

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusOK, gin.H{
			"status": "fail",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": "success",
		})
	}
}
