package controllers

import (
	"fmt"
	"strconv"

	"github.com/labstack/echo/v4"

	"book-api-demo/lib/database"
	"book-api-demo/models"
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

func CreateUser(c echo.Context) error {
	var user models.Users
	c.Bind(&user)

	user, err := database.CreateUser(user)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new user",
		"user":    user,
	})
}

func GetUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}

	user, err := database.GetUser(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "failed get user in handler",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": fmt.Sprintf("success get one user with id %d", id),
		"users":    user,
	})

}
