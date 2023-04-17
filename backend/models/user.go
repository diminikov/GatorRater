package models

import (
	"gorm.io/gorm"
)

type User struct {
	UID      uint   `gorm:"primaryKey"`
	Username string `gorm:"unique"`
	Password string
}

// create a user
func CreateUser(db *gorm.DB, User *User) (err error) {
	err = db.Create(User).Error
	if err != nil {
		return err
	}
	return nil
}

// get users
func GetUsers(db *gorm.DB, User *[]User) (err error) {
	err = db.Find(User).Error
	if err != nil {
		return err
	}
	return nil
}

// get user by username
func GetUser(db *gorm.DB, User *User, username string) (err error) {
	err = db.Where("username = ?", username).First(User).Error
	if err != nil {
		return err
	}
	return nil
}

// update user
func UpdateUser(db *gorm.DB, User *User) (err error) {
	db.Save(User)
	return nil
}

// delete user
func DeleteUser(db *gorm.DB, User *User, username string) (err error) {
	db.Where("username = ?", username).Delete(User)
	return nil
}
