package controllers

import (
	"context"
	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"net/http"
	"os"
	"time"

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
		c.JSON(http.StatusServiceUnavailable, err.Error())
		log.Fatal(err.Error())
		return
	}

	// check if email exist in database
	userAuthenticate := db.Collection("user_authenticate")
	var user models.UserAuthenticate
	err = userAuthenticate.FindOne(context.TODO(), bson.M{
		"email": userEmail,
	}).Decode(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		log.Fatal(err.Error())
		return
	}

	// compare password hash
	compareErr := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userPassword))
	if compareErr != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status": "fail",
		})
		return
	}

	// create token
	token, err := CreateToken(user.Id)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		log.Fatal(err.Error())
		return
	}

	// response on success
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"token":  token,
		"id":     user.Id,
	})
}

func CreateToken(userid primitive.ObjectID) (string, error) {
	var err error
	//Creating Access Token
	os.Setenv("ACCESS_SECRET", "jdnfksdmfksd") //this should be in an env file
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_id"] = userid
	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		return "", err
	}
	return token, nil
}
