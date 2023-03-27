package main

import (
	"bytes"
	"gatorrater/database"
	"gatorrater/models"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

// helper function to check if user is in database
func checkDBForClass(db *gorm.DB, class *models.Class, u string) bool {
	err := db.Where("name = ?", u).First(class).Error
	if err != nil {
		return false
	}
	return true
}

func TestMakeClass(t *testing.T) {
	router := setupRouter()

	var jsonData = []byte(`{
		"name": "morpheus"
	}`)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/class", bytes.NewBuffer(jsonData))
	req.Header.Add("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	var class models.Class

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, true, checkDBForClass(database.Db, &class, "morpheus"))
}

func TestGetClass(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/class/morpheus", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}

func TestEditClass(t *testing.T) {
	router := setupRouter()

	var jsonData = []byte(`{
		"name": "smith"
	}`)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", "/class/morpheus", bytes.NewBuffer(jsonData))
	req.Header.Add("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	var class models.Class

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, false, checkDBForClass(database.Db, &class, "morpheus"))
	assert.Equal(t, true, checkDBForClass(database.Db, &class, "smith"))
}

func TestDeleteClass(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/class/smith", nil)
	router.ServeHTTP(w, req)

	var class models.Class

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, false, checkDBForClass(database.Db, &class, "morpheus"))
	assert.Equal(t, false, checkDBForClass(database.Db, &class, "smith"))
}
