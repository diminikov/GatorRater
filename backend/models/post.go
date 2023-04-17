package models

import (
	"gorm.io/gorm"
)

type Post struct {
	PID  int `gorm:"primaryKey"`
	CID  int
	UID  int
	Body string
}

// create a post
func CreatePost(db *gorm.DB, Post *Post) (err error) {
	err = db.Create(Post).Error
	if err != nil {
		return err
	}
	return nil
}

// get posts by username
func GetPostFromUser(db *gorm.DB, Post *[]Post, username string) (err error) {
	err = db.Joins("NATURAL JOIN users").Where("username = ?", username).Find(Post).Error
	if err != nil {
		return err
	}
	return nil
}
