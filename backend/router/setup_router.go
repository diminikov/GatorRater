package router

import (
	"gatorrater/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:4200") // Replace with the origin domain of your client-side application
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, PATCH, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
		c.Next()
	})

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, "pong")
	})

	grRepo := controllers.NewGRRepo()

	//Create HTTP Requests for Users
	r.POST("/users", grRepo.CreateUser)
	r.GET("/users", grRepo.GetUsers)
	r.GET("/users/:username", grRepo.GetUser)
	r.PUT("/users/:username", grRepo.UpdateUser)
	r.DELETE("/users/:username", grRepo.DeleteUser)

	//Create HTTP Requests for Classes
	r.POST("/class", grRepo.CreateClass)
	r.GET("/class", grRepo.GetClasses)
	r.GET("/class/:name", grRepo.GetClass)
	r.PUT("/class/:name", grRepo.UpdateClass)
	r.DELETE("/class/:name", grRepo.DeleteClass)

	//Create HTTP Requests for Posts
	r.POST("/post", grRepo.CreatePost)
	r.GET("/post/username/:username", grRepo.GetPostFromUser)
	r.GET("/post/class/:classname", grRepo.GetPostFromClass)

	return r
}
