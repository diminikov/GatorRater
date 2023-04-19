package controllers

import (
	"errors"
	"gatorrater/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type PostInput struct {
	Class    string
	Username string
	Body     string
}

// create post
func (repository *GatorRaterRepo) CreatePost(c *gin.Context) {
	var post_input PostInput
	c.BindJSON(&post_input)

	var post models.Post
	repository.Db.Table("users").Joins("NATURAL JOIN classes").Select("c_id, uid").Where("username = ?", post_input.Username).Where("name = ?", post_input.Class).Find(&post)

	if post.CID == 0 || post.UID == 0 {
		c.AbortWithStatusJSON(http.StatusInternalServerError, `{error : invalid class or user}`)
		return
	}

	post.Body = post_input.Body

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

// get all posts in the class
func (repository *GatorRaterRepo) GetPostFromClass(c *gin.Context) {
	name := c.Param("classname")
	var post []models.PostJsonUser
	err := models.GetPostFromClass(repository.Db, &post, name)
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
