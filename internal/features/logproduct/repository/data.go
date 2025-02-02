package repository

import (
	"evermos_pbi/internal/features/logproduct"
	"evermos_pbi/internal/features/products"

	"gorm.io/gorm"
)

type LogProduct struct {
	gorm.Model
	ProductName   string
	ProductImage  string
	Slug          string
	ResellerPrice float32
	ConsumenPrice float32
	Stock         uint
	Description   string
	ProductID     uint
	Product       products.Product `gorm:"foreignKey:ProductID"`
}

func (l *LogProduct) ToLogProductEntity() logproduct.LogProduct {
	return logproduct.LogProduct{
		ID:            l.ID,
		ProductName:   l.ProductName,
		ProductImage:  l.ProductImage,
		Slug:          l.Slug,
		ResellerPrice: l.ResellerPrice,
		ConsumenPrice: l.ConsumenPrice,
		Stock:         l.Stock,
		Description:   l.Description,
		ProductID:     l.ProductID,
		CreatedAt:     l.CreatedAt,
		UpdatedAt:     l.UpdatedAt,
	}
}

func ToLogProductQuery(input logproduct.LogProduct) LogProduct {
	return LogProduct{
		ProductName:   input.ProductName,
		ProductImage:  input.ProductImage,
		Slug:          input.Slug,
		ResellerPrice: input.ResellerPrice,
		ConsumenPrice: input.ConsumenPrice,
		Stock:         input.Stock,
		Description:   input.Description,
		ProductID:     input.ProductID,
	}
}
