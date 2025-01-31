package handler

import (
	"evermos_pbi/internal/features/categories"
	"time"
)

type CategoryResponse struct {
	ID           uint      `json:"id"`
	CategoryName string    `json:"category_name"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

func ToCategoryResponse(input categories.Category) CategoryResponse {
	return CategoryResponse{
		ID:           input.ID,
		CategoryName: input.CategoryName,
		CreatedAt:    input.CreatedAt,
		UpdatedAt:    input.UpdatedAt,
	}
}

func ToCategoryResponses(categories []categories.Category) []CategoryResponse {
	responses := make([]CategoryResponse, len(categories))
	for i, category := range categories {
		responses[i] = ToCategoryResponse(category)
	}
	return responses
}
