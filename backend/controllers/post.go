package controllers

import (
	"gatorrater/database"
	"gatorrater/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type PostRepo struct {
	Db *gorm.DB
}

func NewPostRepo() *PostRepo {
	db := database.InitDb()
	db.AutoMigrate(&models.Post{})
	return &PostRepo{Db: db}
}

// create post
func (repository *PostRepo) CreatePost(c *gin.Context) {
	var post models.Post
	c.BindJSON(&post)
	err := models.CreatePost(repository.Db, &post)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, post)
}
