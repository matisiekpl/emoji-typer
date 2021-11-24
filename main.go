package main

import (
	"github.com/labstack/echo/v4"
	"os"
)

func main() {
	e := echo.New()
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	e.Logger.Fatal(e.Start(":" + port))
}
