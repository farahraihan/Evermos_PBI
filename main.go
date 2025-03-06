package main

import (
	"evermos_pbi/internal/factory"
	"os"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000" // default port
	}

	factory.InitFactory(e)
	e.Logger.Fatal(e.Start(":" + port))
}
