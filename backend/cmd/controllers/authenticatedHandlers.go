package controllers

import (
	"github.com/gin-gonic/gin"
)

//LoginPageHandler handler for /login
func LoginPageHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"hello": "ga",
	})
}
