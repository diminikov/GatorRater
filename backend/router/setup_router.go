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

	userRepo := controllers.NewUserRepo()
	r.POST("/users", userRepo.CreateUser)
	r.GET("/users", userRepo.GetUsers)
	r.GET("/users/:username", userRepo.GetUser)
	r.PUT("/users/:username", userRepo.UpdateUser)
	r.DELETE("/users/:username", userRepo.DeleteUser)

	classRepo := controllers.NewClassRepo()
	r.POST("/class", classRepo.CreateClass)
	r.GET("/class", classRepo.GetClasses)
	r.GET("/class/:name", classRepo.GetClass)
	r.PUT("/class/:name", classRepo.UpdateClass)
	r.DELETE("/class/:name", classRepo.DeleteClass)

	return r
}
