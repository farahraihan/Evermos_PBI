package handler

import "evermos_pbi/internal/features/categories"

type AddOrUpdateCategoryRequest struct {
	CategoryName string `json:"category_name" form:"category_name"`
}

func AddToCategory(cr AddOrUpdateCategoryRequest) categories.Category {
	return categories.Category{
		CategoryName: cr.CategoryName,
	}
}
