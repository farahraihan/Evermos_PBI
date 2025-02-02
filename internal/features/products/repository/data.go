package repository

import (
	"evermos_pbi/internal/features/categories"
	"evermos_pbi/internal/features/products"
	"evermos_pbi/internal/features/stores"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	ProductName   string
	ProductImage  string
	Slug          string
	ResellerPrice float32
	ConsumenPrice float32
	Stock         uint
	Description   string
	StoreID       uint
	CategoryID    uint
	Store         stores.Store        `gorm:"foreignKey:StoreID"`
	Category      categories.Category `gorm:"foreignKey:CategoryID"`
}

func (p *Product) ToProductEntity() products.Product {
	return products.Product{
		ID:            p.ID,
		ProductName:   p.ProductName,
		ProductImage:  p.ProductImage,
		Slug:          p.Slug,
		ResellerPrice: p.ResellerPrice,
		ConsumenPrice: p.ConsumenPrice,
		Stock:         p.Stock,
		Description:   p.Description,
		StoreID:       p.StoreID,
		StoreName:     p.Store.StoreName,
		CategoryID:    p.CategoryID,
		CategoryName:  p.Category.CategoryName,
		CreatedAt:     p.CreatedAt,
		UpdatedAt:     p.UpdatedAt,
	}
}

func ToProductQuery(input products.Product) Product {
	return Product{
		ProductName:   input.ProductName,
		ProductImage:  input.ProductImage,
		Slug:          input.Slug,
		ResellerPrice: input.ResellerPrice,
		ConsumenPrice: input.ConsumenPrice,
		Stock:         input.Stock,
		Description:   input.Description,
		StoreID:       input.StoreID,
		CategoryID:    input.CategoryID,
	}
}
