package main

// test suite to make sure you can successfully add a post, searching for a post, doesn't allow you to add a post for the wrong user
import (
	"bytes"
	"gatorrater/database"
	"gatorrater/router"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

// Function to check if the post is in the database
func checkDBForPost(db *gorm.DB, body string) bool {
	err := db.Table("posts").Where("body = ?", body).Error
	if err != nil {
		return false
	}
	return true
}

// Checks to see if post has been made correctly
func TestMakePost(t *testing.T) {
	router := router.SetupRouter()

	var jsonData = []byte(`{
		"Username": "morpheus",
		"Class": "CEN3031",
		"Body": " This class is amazing"

	}`)

	var userData = []byte(`{
		"Username": "morpheus",
		"Password": "leader"
	}`)

	var classData = []byte(`{
		"Class": "CEN3031"
	}`)

	w := httptest.NewRecorder()
	//Create a user
	userReq, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(userData))
	userReq.Header.Add("Content-Type", "application/json")
	router.ServeHTTP(w, userReq)

	//Create a class
	classReq, _ := http.NewRequest("POST", "/class", bytes.NewBuffer(classData))
	classReq.Header.Add("Content-Type", "application/json")
	router.ServeHTTP(w, classReq)

	//Making a post
	req, _ := http.NewRequest("POST", "/post", bytes.NewBuffer(jsonData))
	req.Header.Add("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, true, checkDBForPost(database.Db, "This class is amazing"))
}

// Searching for all posts from a specific user
func TestGetPostFromUser(t *testing.T) {
	router := router.SetupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/post/username/morpheus", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}

// Searching for all posts from a specific class
func TestGetPostFromClass(t *testing.T) {
	router := router.SetupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/post/class/CEN3031", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}

// Testing to make sure it will return an error if a post is made from an non-existing user
func TestIncorrectUserPost(t *testing.T) {
	router := router.SetupRouter()

	var jsonData = []byte(`{
		"Username": "pineapple",
		"Class": "CEN3031",
		"Body": " This class is amazing"

	}`)

	w := httptest.NewRecorder()
	userReq, _ := http.NewRequest("POST", "/post", bytes.NewBuffer(jsonData))
	userReq.Header.Add("Content-Type", "application/json")
	router.ServeHTTP(w, userReq)

	//500 is the fail code and we want to ensure that user doesn't exist
	assert.Equal(t, 500, w.Code)
}

// Testing to make sure it will return an error if a post is made from an non-existing class
func TestIncorrectClassPost(t *testing.T) {
	router := router.SetupRouter()

	var jsonData = []byte(`{
		"Username": "morepheus",
		"Class": "INT2021",
		"Body": " This class is amazing"

	}`)

	w := httptest.NewRecorder()
	userReq, _ := http.NewRequest("POST", "/post", bytes.NewBuffer(jsonData))
	userReq.Header.Add("Content-Type", "application/json")
	router.ServeHTTP(w, userReq)

	//500 is the fail code and we want to ensure that user doesn't exist
	assert.Equal(t, 500, w.Code)
}

// func TestEditPost(t *testing.T) {
// 	router := router.SetupRouter()

// 	var jsonData = []byte(`{
// 		"Username": "smith",
// 		"Password": "agent"
// 	}`)

// 	w := httptest.NewRecorder()
// 	req, _ := http.NewRequest("PUT", "/users/morpheus", bytes.NewBuffer(jsonData))
// 	req.Header.Add("Content-Type", "application/json")
// 	router.ServeHTTP(w, req)

// 	var user models.User

// 	assert.Equal(t, 200, w.Code)
// 	assert.Equal(t, false, checkDBForUser(database.Db, &user, "morpheus", "leader"))
// 	assert.Equal(t, true, checkDBForUser(database.Db, &user, "smith", "agent"))
// }

// func TestDeletePost(t *testing.T) {
// 	router := router.SetupRouter()

// 	w := httptest.NewRecorder()
// 	req, _ := http.NewRequest("DELETE", "/users/smith", nil)
// 	router.ServeHTTP(w, req)

// 	var user models.User

// 	assert.Equal(t, 200, w.Code)
// 	assert.Equal(t, false, checkDBForUser(database.Db, &user, "morpheus", "leader"))
// 	assert.Equal(t, false, checkDBForUser(database.Db, &user, "smith", "agent"))
// }
