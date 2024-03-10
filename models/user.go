package models

import (
	"mochi/db"
	"github.com/labstack/echo/v4"
)

func GetAllUsers() ([]db.User, error) {
	users := []db.User{}

	if db.DB.Find(&users).Error != nil {
		return nil, echo.ErrNotFound
	}
	return users, nil
}
func GetUserById(id string) (*db.User, error) {
	user := db.User{}

	if db.DB.Where("id = ?", id).First(&user).Error != nil {
		return nil, echo.ErrNotFound
	}
	return &user, nil
}

func CreateUser(name string) (*db.User, error) {
	user := db.User{Name: name}
	if db.DB.Create(&user).Error != nil {
		return nil, echo.ErrInternalServerError
	}
	return &user, nil
}