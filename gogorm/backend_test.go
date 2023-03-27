package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	//"gatorrater/database"

	"github.com/stretchr/testify/assert"
)

// testing the connection to the backend
func TestPingRoute(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "\"pong\"", w.Body.String())
}

func TestGetUser(t *testing.T) {
	router := setupRouter()

	var jsonData = []byte(`{
		"Username": "morpheus",
		"Password": "leader"
	}`)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(jsonData))
	req.Header.Add("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)

}
