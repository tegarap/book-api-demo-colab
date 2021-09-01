package database

import (
	"book-api-demo/config"
	"book-api-demo/models"
	"errors"
)

func GetUsers() (interface{}, error) {
	var users []models.Users

	err := config.Db.Find(&users).Error
	if err != nil {
		return nil, err
	}

	return users, nil

}

func GetUser(id int) (models.Users, error) {
	var user models.Users

	if err := config.Db.Find(&user, id).Error; err != nil {
		return models.Users{}, errors.New("cant find user with id")
	}

	return user, nil
}

func CreateUser(user models.Users) (models.Users, error) {

	if err := config.Db.Save(&user).Error; err != nil {
		return models.Users{}, err
	}

	return user, nil

}
