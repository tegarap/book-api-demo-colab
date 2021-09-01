package controllers

import (
	"github.com/labstack/echo/v4"

	"book-api-demo/lib/database"
	"net/http"
)

func GetUsers(c echo.Context) error {
	users, err := database.GetUsers()

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all users",
		"users":   users,
	})

}
