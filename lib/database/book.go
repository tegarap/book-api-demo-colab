package database

import (
	"book-api-demo/config"
	"book-api-demo/models"
)

func GetBooks() (interface{}, error) {
	var books []models.Book
	if e := config.Db.Find(&books).Error; e != nil {
		return nil, e
	}
	return books, nil
}

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

func DeleteBook(id int) (interface{}, error) {
	var book []models.Book
	err := config.Db.Delete(&book, id).Error
	if err != nil {
		return nil, err
	}
	return "deleted", nil
}

func UpdateBook(id int, newBook *models.Book) (interface{}, int, error) {
	var book models.Book
	tx1 := config.Db.Find(&book, id)

	if tx1.Error != nil {
		return nil, 0, tx1.Error
	}

	if tx1.RowsAffected > 0 {
		tx := config.Db.Model(&book).Updates(newBook)

		if tx.Error != nil {
			return nil, 0, tx.Error
		}

		return book, 1, nil
	}
	return nil, 0, nil
}
