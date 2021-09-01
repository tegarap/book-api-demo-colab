package controllers

import (
	"book-api-demo/lib/database"
	"book-api-demo/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetBooksController(c echo.Context) error {
	books, err := database.GetBooks()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success get one book",
		"book":     books,
	})
}

func GetOneBookController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	book, _, err := database.GetOneBook(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success get one book",
		"book":     book,
	})
}

func CreateBookController(c echo.Context) error {
	var books models.Book
	c.Bind(&books)

	book, err := database.CreateNewBook(&books)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success create new book",
		"book":     book,
	})

}

func DeleteBookController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	message, err := database.DeleteBook(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": message,
	})
}
