package repository

import (
	"evermos_pbi/internal/features/categories"

	"gorm.io/gorm"
)

type CategoryQuery struct {
	db *gorm.DB
}

func NewCategoryQuery(connect *gorm.DB) categories.CQuery {
	return &CategoryQuery{
		db: connect,
	}
}

func (cq *CategoryQuery) AddCategory(newCategory categories.Category) error {
	cnvData := ToCategoryQuery(newCategory)
	qry := cq.db.Create(&cnvData)

	if qry.Error != nil {
		return qry.Error
	}

	return nil
}

func (cq *CategoryQuery) UpdateCategory(categoryID uint, updateCategory categories.Category) error {
	cnvData := ToCategoryQuery(updateCategory)

	qry := cq.db.Where("id = ?", categoryID).Updates(&cnvData)

	if qry.Error != nil {
		return qry.Error
	}

	if qry.RowsAffected < 1 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func (cq *CategoryQuery) DeleteCategory(categoryID uint) error {
	qry := cq.db.Where("id = ?", categoryID).Delete(&Category{})

	if qry.Error != nil {
		return qry.Error
	}

	if qry.RowsAffected < 1 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func (cq *CategoryQuery) GetCategoryByID(categoryID uint) (categories.Category, error) {
	var category categories.Category

	err := cq.db.First(&category, categoryID).Error
	if err != nil {
		return categories.Category{}, err
	}
	return category, nil
}

func (cq *CategoryQuery) GetAllCategories(limit uint, page uint, search string) ([]categories.Category, uint, error) {
	var categoryList []Category
	var totalItems int64

	offset := (page - 1) * limit

	qry := cq.db.Model(&Category{})
	if search != "" {
		qry = qry.Where("category_name ILIKE ?", "%"+search+"%")
	}
	if err := qry.Count(&totalItems).Error; err != nil {
		return nil, 0, err
	}

	qry = cq.db.Limit(int(limit)).Offset(int(offset)).Find(&categoryList)
	if qry.Error != nil {
		return nil, 0, qry.Error
	}

	categoryEntities := make([]categories.Category, len(categoryList))
	for i, category := range categoryList {
		categoryEntities[i] = category.ToCategoryEntity()
	}

	return categoryEntities, uint(totalItems), nil
}
