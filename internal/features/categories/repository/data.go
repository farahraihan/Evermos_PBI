package repository

import (
	"evermos_pbi/internal/features/categories"

	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	CategoryName string
}

func (c *Category) ToCategoryEntity() categories.Category {
	return categories.Category{
		ID:           c.ID,
		CategoryName: c.CategoryName,
		CreatedAt:    c.CreatedAt,
		UpdatedAt:    c.UpdatedAt,
	}
}

func ToCategoryQuery(input categories.Category) Category {
	return Category{
		CategoryName: input.CategoryName,
	}
}
