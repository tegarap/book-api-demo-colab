package routes

import (
	"book-api-demo/controllers"
	"github.com/labstack/echo"
)

func New() *echo.Echo {
	e := echo.New()

	// User
	e.DELETE("/users/:id", controllers.DeleteUserController)
	//Book

	return e
}