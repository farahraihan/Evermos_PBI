package main

import (
	"net/http"

	"evermos_pbi/internal/factory"

	"github.com/labstack/echo/v4"
)

var e *echo.Echo

func init() {
	e = echo.New()
	factory.InitFactory(e)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	e.ServeHTTP(w, r)
}
