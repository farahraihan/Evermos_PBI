package products

import (
	"evermos_pbi/internal/features/categories"
	"evermos_pbi/internal/features/stores"
	"mime/multipart"
	"time"

	"github.com/labstack/echo/v4"
)

type Product struct {
	ID            uint
	ProductName   string
	ProductImage  string
	Slug          string
	ResellerPrice float32
	ConsumenPrice float32
	Stock         uint
	Description   string
	StoreID       uint
	StoreName     string
	CategoryID    uint
	CategoryName  string
	Store         stores.Store        `gorm:"foreignKey:StoreID"`
	Category      categories.Category `gorm:"foreignKey:CategoryID"`
	CreatedAt     time.Time           `gorm:"default:current_timestamp"`
	UpdatedAt     time.Time           `gorm:"default:current_timestamp"`
}

type PQuery interface {
	AddProduct(newProduct *Product) error
	UpdateProduct(productID uint, UpdateProduct Product) error
	DeleteProduct(productID uint) error
	GetProductByID(productID uint) (Product, error)
	GetProductsByStoreID(storeID uint, limit uint, page uint, search string) ([]Product, uint, error)
	GetAllProducts(limit uint, page uint, search string) ([]Product, uint, error)
	IsProductOwnedByUser(productID uint, userID uint) (bool, error)
	IsStock(productID uint, n uint) (bool, error)
	IncreaseStock(productID uint, n uint) error
	DecreaseStock(productID uint, n uint) error
}

type PService interface {
	AddProduct(userID uint, newProduct Product, src multipart.File, filename string) error
	UpdateProduct(userID uint, productID uint, newProduct Product, src multipart.File, filename string) error
	DeleteProduct(userID uint, productID uint) error
	GetProductByID(productID uint) (Product, error)
	GetProductsByStoreID(storeID uint, limit uint, page uint, search string) ([]Product, uint, error)
	GetAllProducts(limit uint, page uint, search string) ([]Product, uint, error)
	IsStock(productID uint, n uint) (bool, error)
	IncreaseStock(productID uint, n uint) error
	DecreaseStock(productID uint, n uint) error
}

type PHandler interface {
	AddProduct() echo.HandlerFunc
	UpdateProduct() echo.HandlerFunc
	DeleteProduct() echo.HandlerFunc
	GetProductByID() echo.HandlerFunc
	GetProductsByStoreID() echo.HandlerFunc
	GetAllProducts() echo.HandlerFunc
}
