package categories

import (
	"time"

	"github.com/labstack/echo/v4"
)

type Category struct {
	ID           uint
	CategoryName string
	CreatedAt    time.Time `gorm:"default:current_timestamp"`
	UpdatedAt    time.Time `gorm:"default:current_timestamp"`
}

type CQuery interface {
	AddCategory(newCategory Category) error
	UpdateCategory(categoryID uint, updateCategory Category) error
	DeleteCategory(categoryID uint) error
	GetCategoryByID(categoryID uint) (Category, error)
	GetAllCategories(limit uint, page uint, search string) ([]Category, uint, error)
}

type CService interface {
	AddCategory(userID uint, newCategory Category) error
	UpdateCategory(userID uint, categoryID uint, updateCategory Category) error
	DeleteCategory(userID uint, categoryID uint) error
	GetCategoryByID(categoryID uint) (Category, error)
	GetAllCategories(limit uint, page uint, search string) ([]Category, uint, error)
}

type CHandler interface {
	AddCategory() echo.HandlerFunc
	UpdateCategory() echo.HandlerFunc
	DeleteCategory() echo.HandlerFunc
	GetCategoryByID() echo.HandlerFunc
	GetAllCategories() echo.HandlerFunc
}
