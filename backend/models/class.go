package models

import (
	"gorm.io/gorm"
)

type Class struct {
	gorm.Model
	ID   int
	Name string
}

// create a class
func CreateClass(db *gorm.DB, Class *Class) (err error) {
	err = db.Create(Class).Error
	if err != nil {
		return err
	}
	return nil
}

// get all the Classes
func GetClasses(db *gorm.DB, Class *[]Class) (err error) {
	err = db.Find(Class).Error
	if err != nil {
		return err
	}
	return nil
}

// get specfic Class by classname
func GetClass(db *gorm.DB, Class *Class, name string) (err error) {
	err = db.Where("name = ?", name).First(Class).Error
	if err != nil {
		return err
	}
	return nil
}

// update Class
func UpdateClass(db *gorm.DB, Class *Class) (err error) {
	db.Save(Class)
	return nil
}

// delete Class
func DeleteClass(db *gorm.DB, Class *Class, name string) (err error) {
	db.Where("name = ?", name).Delete(Class)
	return nil
}
