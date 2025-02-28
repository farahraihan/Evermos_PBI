package main

import (
	"evermos_pbi/internal/factory"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

// Handler untuk Vercel
func Handler(w http.ResponseWriter, r *http.Request) {
	e := echo.New()
	factory.InitFactory(e)
	e.ServeHTTP(w, r)
}

func main() {
	e := echo.New()
	factory.InitFactory(e)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	e.Logger.Fatal(e.Start(":" + port))
}
