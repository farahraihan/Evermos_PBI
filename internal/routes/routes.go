package routes

import (
	"evermos_pbi/config"
	"evermos_pbi/internal/features/users"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func InitRoute(e *echo.Echo, uh users.UHandler) {
	e.POST("/login", uh.Login())
	e.POST("/register", uh.Register())

	UserRoute(e, uh)
}

func UserRoute(e *echo.Echo, uh users.UHandler) {
	u := e.Group("/users")
	u.Use(JWTConfig())
	u.PUT("", uh.UpdateUser())
	u.DELETE("/:id", uh.DeleteUser())
	u.GET("", uh.GetUserByID())
}

func JWTConfig() echo.MiddlewareFunc {
	return echojwt.WithConfig(
		echojwt.Config{
			SigningKey:    []byte(config.ImportSetting().JWTSecret),
			SigningMethod: jwt.SigningMethodHS256.Name,
		},
	)
}
