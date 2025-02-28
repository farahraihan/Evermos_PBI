package main

import (
	"evermos_pbi/internal/factory"
	"os"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	factory.InitFactory(e)

	// Gunakan port dari environment (default ke 3000 untuk Vercel)
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	e.Logger.Fatal(e.Start(":" + port))
}
