package routes

import (
	"book-api-demo/controllers"

	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	e := echo.New()
	// User
	e.DELETE("/users/:id", controllers.DeleteUserController)

	//Book
	e.GET("/books/:id", controllers.GetOneBookController)

	return e
}
