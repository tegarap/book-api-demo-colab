package routes

import (
	"book-api-demo/controllers"

	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	e := echo.New()
	// User
	e.GET("/users", controllers.GetUsers)
	e.POST("/users", controllers.CreateUser)
	e.GET("/users/:id", controllers.GetUser)
	e.DELETE("/users/:id", controllers.DeleteUserController)
	e.PUT("/users/:id", controllers.UpdateUserController)

	//Book
	e.GET("/books", controllers.GetBooksController)
	e.GET("/books/:id", controllers.GetOneBookController)
	e.POST("/books", controllers.CreateBookController)
	e.DELETE("/books/:id", controllers.DeleteBookController)
	e.PUT("/books/:id", controllers.UpdateBookController)

	return e
}
