package controllers

import (
	"errors"
	"gatorrater/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// create post
func (repository *GatorRaterRepo) CreatePost(c *gin.Context) {
	var post models.Post
	c.BindJSON(&post)
	err := models.CreatePost(repository.Db, &post)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, post)
}

// get all posts by the username
func (repository *GatorRaterRepo) GetPostFromUser(c *gin.Context) {
	name := c.Param("username")
	var post []models.Post
	err := models.GetPostFromUser(repository.Db, &post, name)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}

		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, post)
}
