package handler

import (
	"evermos_pbi/internal/features/products"
	"time"
)

type ProductResponse struct {
	ID            uint      `json:"id"`
	ProductName   string    `json:"product_name"`
	ProductImage  string    `json:"product_image"`
	Slug          string    `json:"slug"`
	ResellerPrice float32   `json:"reseller_price"`
	ConsumenPrice float32   `json:"consumen_price"`
	Stock         uint      `json:"stock"`
	Description   string    `json:"description"`
	StoreID       uint      `json:"store_id"`
	StoreName     string    `json:"store_name"`
	CategoryID    uint      `json:"category_id"`
	CategoryName  string    `json:"category_name"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

func ToProductResponse(input products.Product) ProductResponse {
	return ProductResponse{
		ID:            input.ID,
		ProductName:   input.ProductName,
		ProductImage:  input.ProductImage,
		Slug:          input.Slug,
		ResellerPrice: input.ResellerPrice,
		ConsumenPrice: input.ConsumenPrice,
		Stock:         input.Stock,
		Description:   input.Description,
		StoreID:       input.StoreID,
		StoreName:     input.StoreName,
		CategoryID:    input.CategoryID,
		CategoryName:  input.CategoryName,
		CreatedAt:     input.CreatedAt,
		UpdatedAt:     input.UpdatedAt,
	}
}

func ToProductResponses(products []products.Product) []ProductResponse {
	responses := make([]ProductResponse, len(products))
	for i, product := range products {
		responses[i] = ToProductResponse(product)
	}
	return responses
}
