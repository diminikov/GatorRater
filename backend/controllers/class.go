package controllers

import (
	"errors"
	"gatorrater/database"
	"gatorrater/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ClassRepo struct {
	Db *gorm.DB
}

func NewClassRepo() *ClassRepo {
	db := database.InitDb()
	db.AutoMigrate(&models.Class{})
	return &ClassRepo{Db: db}
}

// create class, intercepting the put request with a post
func (repository *ClassRepo) CreateClass(c *gin.Context) {
	var class models.Class
	c.BindJSON(&class)
	err := models.CreateClass(repository.Db, &class)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, class)
}

// get all classes
func (repository *ClassRepo) GetClasses(c *gin.Context) {
	var class []models.Class
	err := models.GetClasses(repository.Db, &class)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, class)
}

// get specific class by the classname
func (repository *ClassRepo) GetClass(c *gin.Context) {
	name := c.Param("name")
	var class models.Class
	err := models.GetClass(repository.Db, &class, name)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}

		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, class)
}

// update Class
func (repository *ClassRepo) UpdateClass(c *gin.Context) {
	var class models.Class
	name := c.Param("name")
	err := models.GetClass(repository.Db, &class, name)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}

		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.BindJSON(&class)
	err = models.UpdateClass(repository.Db, &class)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, class)
}

// delete a class
func (repository *ClassRepo) DeleteClass(c *gin.Context) {
	var class models.Class
	name := c.Param("name")
	err := models.DeleteClass(repository.Db, &class, name)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Class deleted successfully"})
}
