package routes

import (
	"evermos_pbi/config"
	"evermos_pbi/internal/features/address"
	"evermos_pbi/internal/features/categories"
	"evermos_pbi/internal/features/stores"
	"evermos_pbi/internal/features/users"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func InitRoute(e *echo.Echo, uh users.UHandler, sh stores.SHandler, ah address.AHandler, ch categories.CHandler) {
	e.POST("/login", uh.Login())
	e.POST("/register", uh.Register())
	e.GET("/stores/:id", sh.GetStoreByID())
	e.GET("/stores/all", sh.GetAllStores())
	e.GET("/provinces", ah.GetProvince())
	e.GET("/regency/:province_id", ah.GetRegency())
	e.GET("/district/:regency_id", ah.GetDistrict())
	e.GET("/village/:district_id", ah.GetVillage())
	e.GET("/category/:id", ch.GetCategoryByID())
	e.GET("/category", ch.GetAllCategories())

	UserRoute(e, uh)
	StoreRoute(e, sh)
	AddressRoute(e, ah)
	CategoryRoute(e, ch)
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

func AddressRoute(e *echo.Echo, ah address.AHandler) {
	a := e.Group("/address")
	a.Use(JWTConfig())
	a.POST("", ah.AddAddress())
	a.PUT("/:id", ah.UpdateAddress())
	a.DELETE("/:id", ah.DeleteAddress())
	a.GET("", ah.GetAddressByUserID())
}

func CategoryRoute(e *echo.Echo, ch categories.CHandler) {
	c := e.Group("/category")
	c.Use(JWTConfig())
	c.POST("", ch.AddCategory())
	c.PUT("/:id", ch.UpdateCategory())
	c.DELETE("/:id", ch.DeleteCategory())
}

func JWTConfig() echo.MiddlewareFunc {
	return echojwt.WithConfig(
		echojwt.Config{
			SigningKey:    []byte(config.ImportSetting().JWTSecret),
			SigningMethod: jwt.SigningMethodHS256.Name,
		},
	)
}
