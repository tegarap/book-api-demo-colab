package main

import (
	"book-api-demo/routes"
)

func main() {
	e := routes.New()

	e.Logger.Fatal(e.Start(":1323"))
}
