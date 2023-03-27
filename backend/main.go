package main

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// user represents data about website users.
type user struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// users slice to seed record user data.
var users = []user{
	{ID: "1", Username: "Michael", Password: "password"},
	{ID: "2", Username: "Dimitri", Password: "password"},
	{ID: "3", Username: "Nishant", Password: "password"},
}

func main() {
	
	router := gin.Default()
	router.Use(cors.Default())
	router.GET("api/users", getusers)
	router.GET("api/login/:username/:password", getLogin)


	router.Run("localhost:8080")
}

// getusers responds with the list of all users as JSON.
func getusers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, users)
}

// getLogin responds with the data for that specifc user login details
func getLogin(c *gin.Context) {
	username := c.Param("username")
	//println(username)
	password := c.Param("password")
	//println(password)

	// Loop over the list of users, looking for
	// an user whose ID value matches the parameter.
	for _, a := range users {
		if a.Username == username && a.Password == password {
			c.IndentedJSON(http.StatusOK, a)
			return
		}

	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "user not found"})
}
