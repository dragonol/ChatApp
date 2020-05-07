package router

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/static"
	controllers "../controllers"
)

//InitRouter initialize router
func InitRouter() *gin.Engine {
	router := gin.Default()

	router.Use(static.Serve("/", static.LocalFile("../../frontend/dist", false)))

	router.GET("/login", controllers.LoginPageHandler)
	
	return router
}
