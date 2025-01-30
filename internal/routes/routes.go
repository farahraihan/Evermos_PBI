package routes

import (
	"evermos_pbi/config"
	"evermos_pbi/internal/features/stores"
	"evermos_pbi/internal/features/users"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func InitRoute(e *echo.Echo, uh users.UHandler, sh stores.SHandler) {
	e.POST("/login", uh.Login())
	e.POST("/register", uh.Register())
	e.GET("/stores/:id", sh.GetStoreByID())
	e.GET("/stores/all", sh.GetAllStores())

	UserRoute(e, uh)
	StoreRoute(e, sh)
}

func UserRoute(e *echo.Echo, uh users.UHandler) {
	u := e.Group("/users")
	u.Use(JWTConfig())
	u.PUT("", uh.UpdateUser())
	u.DELETE("", uh.DeleteUser())
	u.GET("", uh.GetUserByID())
}

func StoreRoute(e *echo.Echo, sh stores.SHandler) {
	s := e.Group("/stores")
	s.Use(JWTConfig())
	s.POST("", sh.AddStore())
	s.PUT("", sh.UpdateStore())
	s.DELETE("", sh.DeleteStore())
	s.GET("", sh.GetStoreByUserID())
}

func JWTConfig() echo.MiddlewareFunc {
	return echojwt.WithConfig(
		echojwt.Config{
			SigningKey:    []byte(config.ImportSetting().JWTSecret),
			SigningMethod: jwt.SigningMethodHS256.Name,
		},
	)
}
