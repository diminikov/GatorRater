package main

import (
	"gatorrater/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
)

// test with curl -X POST http://localhost:8080/users -H "Content-Type: application/json" -d "{\"username\": \"diminikov\",\"password\": \"1234\"}"
func main() {
	r := setupRouter()
	_ = r.Run(":8080")
}

func setupRouter() *gin.Engine {
	r := gin.Default()

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
