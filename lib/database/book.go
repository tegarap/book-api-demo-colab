package database

import (
	"book-api-demo/config"
	"book-api-demo/models"
)

func GetOneBook(id int) (interface{}, int, error) {
	var book models.Book
	err := config.Db.Find(&book, id).Error

	if err != nil {
		return nil, 0, err
	}

	return book, 1, nil
}

func CreateNewBook(book *models.Book) (interface{}, error) {
	err := config.Db.Create(&book).Error

	if err != nil {
		return nil, err
	}

	return book, nil
}
