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

func DeleteUser(id int) (interface{}, error) {
	user := models.Users{}
	if e := config.Db.Delete(&user, id).Error; e != nil {
		return nil, e
	}
	return user, nil
}

func UpdateUser(id int, usr *models.Users) (interface{}, error) {
	user := models.Users{}
	e := config.Db.First(&user, id)
	if e.Error != nil {
		return nil, e.Error
	}
	if e.RowsAffected > 0 {
		if usr.Name != "" {
			user.Name = usr.Name
		}
		if usr.Email != "" {
			user.Email = usr.Email
		}
		if usr.Password != "" {
			user.Password = usr.Password
		}
		config.Db.Save(&user)
		return &user, nil
	}
	return nil, e.Error
}