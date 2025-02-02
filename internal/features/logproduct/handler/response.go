package handler

import (
	"evermos_pbi/internal/features/logproduct"
	"time"
)

type LogProductResponse struct {
	ID            uint      `json:"id"`
	ProductName   string    `json:"product_name"`
	ProductImage  string    `json:"product_image"`
	Slug          string    `json:"slug"`
	ResellerPrice float32   `json:"reseller_price"`
	ConsumenPrice float32   `json:"consumen_price"`
	Stock         uint      `json:"stock"`
	Description   string    `json:"description"`
	ProductID     uint      `json:"product_id"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

func ToLogProductResponse(input logproduct.LogProduct) LogProductResponse {
	return LogProductResponse{
		ID:            input.ID,
		ProductName:   input.ProductName,
		ProductImage:  input.ProductImage,
		Slug:          input.Slug,
		ResellerPrice: input.ResellerPrice,
		ConsumenPrice: input.ConsumenPrice,
		Stock:         input.Stock,
		Description:   input.Description,
		ProductID:     input.ProductID,
		CreatedAt:     input.CreatedAt,
		UpdatedAt:     input.UpdatedAt,
	}
}

func ToLogProductResponses(logProduct []logproduct.LogProduct) []LogProductResponse {
	responses := make([]LogProductResponse, len(logProduct))
	for i, lProduct := range logProduct {
		responses[i] = ToLogProductResponse(lProduct)
	}
	return responses
}
