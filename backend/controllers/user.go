package controllers

import (
	"errors"
	"gatorrater/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// create user
func (repository *GatorRaterRepo) CreateUser(c *gin.Context) {
	var user models.User
	c.BindJSON(&user)
	err := models.CreateUser(repository.Db, &user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, user)
}

// get users
func (repository *GatorRaterRepo) GetUsers(c *gin.Context) {
	var user []models.User
	err := models.GetUsers(repository.Db, &user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, user)
}

// get user by username
func (repository *GatorRaterRepo) GetUser(c *gin.Context) {
	username := c.Param("username")
	var user models.User
	err := models.GetUser(repository.Db, &user, username)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}

		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, user)
}

// update user
func (repository *GatorRaterRepo) UpdateUser(c *gin.Context) {
	var user models.User
	username := c.Param("username")
	err := models.GetUser(repository.Db, &user, username)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}

		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.BindJSON(&user)
	err = models.UpdateUser(repository.Db, &user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, user)
}

// delete user
func (repository *GatorRaterRepo) DeleteUser(c *gin.Context) {
	var user models.User
	username := c.Param("username")
	err := models.DeleteUser(repository.Db, &user, username)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
