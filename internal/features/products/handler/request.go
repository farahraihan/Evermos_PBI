package handler

import "evermos_pbi/internal/features/products"

type AddOrUpdateProductRequest struct {
	ProductName   string  `json:"product_name" form:"product_name"`
	ProductImage  string  `json:"product_image" form:"product_image"`
	Slug          string  `json:"slug" form:"slug"`
	ResellerPrice float32 `json:"reseller_price" form:"reseller_price"`
	ConsumenPrice float32 `json:"consumen_price" form:"consumen_price"`
	Stock         uint    `json:"stock" form:"stock"`
	Description   string  `json:"description" form:"description"`
	StoreID       uint    `json:"store_id" form:"store_id"`
	CategoryID    uint    `json:"category_id" form:"category_id"`
}

func AddToProduct(pr AddOrUpdateProductRequest) products.Product {
	return products.Product{
		ProductName:   pr.ProductName,
		ProductImage:  pr.ProductImage,
		Slug:          pr.Slug,
		ResellerPrice: pr.ResellerPrice,
		ConsumenPrice: pr.ConsumenPrice,
		Stock:         pr.Stock,
		Description:   pr.Description,
		StoreID:       pr.StoreID,
		CategoryID:    pr.CategoryID,
	}
}
