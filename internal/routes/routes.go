package routes

import (
	"evermos_pbi/config"
	"evermos_pbi/internal/features/address"
	"evermos_pbi/internal/features/categories"
	"evermos_pbi/internal/features/logproduct"
	"evermos_pbi/internal/features/products"
	"evermos_pbi/internal/features/stores"
	"evermos_pbi/internal/features/transaction"
	"evermos_pbi/internal/features/users"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func InitRoute(e *echo.Echo, uh users.UHandler, sh stores.SHandler, ah address.AHandler, ch categories.CHandler, lh logproduct.LHandler, ph products.PHandler, th transaction.THandler) {
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
	e.GET("/product/:id", ph.GetProductByID())
	e.GET("/product/store/:store_id", ph.GetProductsByStoreID())
	e.GET("/product", ph.GetAllProducts())
	e.GET("/logProduct/:id", lh.GetLogProductByID())
	e.GET("/logProduct", lh.GetAllLogProduct())

	UserRoute(e, uh)
	StoreRoute(e, sh)
	AddressRoute(e, ah)
	CategoryRoute(e, ch)
	ProductRoute(e, ph)
	TransactionRoute(e, th)
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

func ProductRoute(e *echo.Echo, ph products.PHandler) {
	p := e.Group("/product")
	p.Use(JWTConfig())
	p.POST("", ph.AddProduct())
	p.PUT("/:id", ph.UpdateProduct())
	p.DELETE("/:id", ph.DeleteProduct())
}

func TransactionRoute(e *echo.Echo, th transaction.THandler) {
	t := e.Group("/transaction")
	t.Use(JWTConfig())
	t.POST("/cart", th.AddTransaction())
	t.PUT("/cart", th.UpdateDetailTransaction())
	t.DELETE("/cart", th.DeleteTransaction())
	t.DELETE("/cart/:transaction_id", th.UpdateTransactionStatusCanceled())
	t.POST("/checkout/:transaction_id", th.UpdateTransactionStatusCompleted())
	t.GET("/cart/:transaction_id", th.GetTransactionByStatusCart())
	t.GET("/history", th.GetTransactionHistory())
	t.GET("/:transaction_id", th.GetTransactionByID())
}

func JWTConfig() echo.MiddlewareFunc {
	return echojwt.WithConfig(
		echojwt.Config{
			SigningKey:    []byte(config.ImportSetting().JWTSecret),
			SigningMethod: jwt.SigningMethodHS256.Name,
		},
	)
}
