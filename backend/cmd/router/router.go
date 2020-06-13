package router

import (
	controllers "../controllers"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

//InitRouter initialize router
func InitRouter() *gin.Engine {
	router := gin.Default()

	router.POST("/api/login", controllers.Login)
	router.POST("/api/register", controllers.Register)
	router.POST("/api/update", controllers.UpdateUserDetail)
	router.GET("/test", controllers.Filter)
	router.Use(static.Serve("/", static.LocalFile("./frontend/dist", false)))

	return router
}
