package main

import (
	"fmt"

	"book-api-demo/controllers"

	"github.com/labstack/echo/v4"
)

func main() {
	fmt.Println("hello")

	c := echo.New()

	c.GET("users", controllers.GetUsers)
	c.POST("user", controllers.CreateUser)
	c.GET("user/:id", controllers.GetUser)

	c.Logger.Fatal(c.Start(":1323"))
}
