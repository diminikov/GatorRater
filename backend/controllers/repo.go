package controllers

import (
	"gatorrater/database"
	"gatorrater/models"

	"gorm.io/gorm"
)

type GatorRaterRepo struct {
	Db *gorm.DB
}

func NewGRRepo() *GatorRaterRepo {
	db := database.InitDb()

	//Creating Tables for All Models
	db.AutoMigrate(&models.Class{})
	db.AutoMigrate(&models.Post{})
	db.AutoMigrate(&models.User{})

	return &GatorRaterRepo{Db: db}
}
