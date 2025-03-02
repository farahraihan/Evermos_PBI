// package handler

// import (
// 	"evermos_pbi/internal/factory"
// 	"net/http"

// 	"github.com/labstack/echo/v4"
// )

// func Handler(w http.ResponseWriter, r *http.Request) {
// 	e := echo.New()
// 	factory.InitFactory(e)
// 	e.ServeHTTP(w, r)
// }

package main

import (
	"evermos_pbi/internal/factory"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	factory.InitFactory(e)
	e.Logger.Error(e.Start(":8000"))

}
