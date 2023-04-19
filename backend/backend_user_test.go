package main

import (
	"bytes"
	"gatorrater/database"
	"gatorrater/models"
	"gatorrater/router"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

// helper function to check if user is in database
func checkDBForUser(db *gorm.DB, user *models.User, u string, p string) bool {
	err := db.Where("username = ?", u).Where("password = ?", p).First(user).Error
	if err != nil {
		return false
	}
	return true
}

// testing the connection to the backend
func TestPingRoute(t *testing.T) {
	router := router.SetupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "\"pong\"", w.Body.String())
}

func TestMakeUser(t *testing.T) {
	router := router.SetupRouter()

	var jsonData = []byte(`{
		"Username": "morpheus",
		"Password": "leader"
	}`)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(jsonData))
	req.Header.Add("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	var user models.User

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, true, checkDBForUser(database.Db, &user, "morpheus", "leader"))
}

func TestGetUser(t *testing.T) {
	router := router.SetupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/users/morpheus", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}

func TestEditUser(t *testing.T) {
	router := router.SetupRouter()

	var jsonData = []byte(`{
		"Username": "smith",
		"Password": "agent"
	}`)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", "/users/morpheus", bytes.NewBuffer(jsonData))
	req.Header.Add("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	var user models.User

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, false, checkDBForUser(database.Db, &user, "morpheus", "leader"))
	assert.Equal(t, true, checkDBForUser(database.Db, &user, "smith", "agent"))
}

func TestDeleteUser(t *testing.T) {
	router := router.SetupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/users/smith", nil)
	router.ServeHTTP(w, req)

	var user models.User

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, false, checkDBForUser(database.Db, &user, "morpheus", "leader"))
	assert.Equal(t, false, checkDBForUser(database.Db, &user, "smith", "agent"))
}
