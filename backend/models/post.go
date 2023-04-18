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

type PostInput struct {
	class    string
	username string
	body     string
}

type PostJsonUser struct {
	Username string
	Body     string
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

// get posts by classname
func GetPostFromClass(db *gorm.DB, PostJsonUser *[]PostJsonUser, classname string) (err error) {
	err = db.Table("posts").Select("users.username, body").Joins("NATURAL JOIN classes NATURAL JOIN users").Where("name = ?", classname).Find(PostJsonUser).Error
	if err != nil {
		return err
	}
	return nil
}
