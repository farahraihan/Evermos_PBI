package logproduct

import (
	"evermos_pbi/internal/features/products"
	"time"

	"github.com/labstack/echo/v4"
)

type LogProduct struct {
	ID            uint
	ProductName   string
	ProductImage  string
	Slug          string
	ResellerPrice float32
	ConsumenPrice float32
	Stock         uint
	Description   string
	ProductID     uint
	Product       products.Product `gorm:"foreignKey:ProductID"`
	CreatedAt     time.Time        `gorm:"default:current_timestamp"`
	UpdatedAt     time.Time        `gorm:"default:current_timestamp"`
}

type LQuery interface {
	AddLogProduct(newLogProduct LogProduct) error
	GetLogProductByID(logProductID uint) (LogProduct, error)
	GetAllLogProduct(limit uint, page uint, search string) ([]LogProduct, uint, error)
}

type LService interface {
	AddLogProduct(newLogProduct LogProduct) error
	GetLogProductByID(logProductID uint) (LogProduct, error)
	GetAllLogProduct(limit uint, page uint, search string) ([]LogProduct, uint, error)
}

type LHandler interface {
	GetLogProductByID() echo.HandlerFunc
	GetAllLogProduct() echo.HandlerFunc
}
