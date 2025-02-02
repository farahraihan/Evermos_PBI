package repository

import (
	"evermos_pbi/internal/features/products"
	"fmt"

	"gorm.io/gorm"
)

type ProductQuery struct {
	db *gorm.DB
}

func NewProductQuery(connect *gorm.DB) products.PQuery {
	return &ProductQuery{
		db: connect,
	}
}

func (pq *ProductQuery) AddProduct(newProduct *products.Product) error {
	cnvData := ToProductQuery(*newProduct)
	qry := pq.db.Create(&cnvData)

	if qry.Error != nil {
		return qry.Error
	}

	if cnvData.ID != 0 {
		newProduct.ID = cnvData.ID
	}

	return nil
}

func (pq *ProductQuery) UpdateProduct(productID uint, updateProduct products.Product) error {
	cnvData := ToProductQuery(updateProduct)

	qry := pq.db.Where("id = ?", productID).Updates(&cnvData)

	if qry.Error != nil {
		return qry.Error
	}

	if qry.RowsAffected < 1 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func (pq *ProductQuery) DeleteProduct(productID uint) error {
	qry := pq.db.Where("id = ?", productID).Delete(&Product{})

	if qry.Error != nil {
		return qry.Error
	}

	if qry.RowsAffected < 1 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func (pq *ProductQuery) GetProductByID(productID uint) (products.Product, error) {
	var product products.Product

	err := pq.db.Preload("Category").Preload("Store").Find(&product, productID).Error
	if err != nil {
		return products.Product{}, err
	}

	product.CategoryName = product.Category.CategoryName
	product.StoreName = product.Store.StoreName

	return product, nil
}

func (pq *ProductQuery) GetProductsByStoreID(storeID uint, limit uint, page uint, search string) ([]products.Product, uint, error) {
	var productsList []Product
	var totalItems int64

	offset := (page - 1) * limit

	qry := pq.db.Model(&Product{}).Where("store_id = ?", storeID)

	if search != "" {
		qry = qry.Where("product_name ILIKE ? OR slug ILIKE ?", "%"+search+"%", "%"+search+"%")
	}

	if err := qry.Count(&totalItems).Error; err != nil {
		return nil, 0, fmt.Errorf("failed to count products: %w", err)
	}

	if err := qry.Preload("Store").Preload("Category").Limit(int(limit)).Offset(int(offset)).Find(&productsList).Error; err != nil {
		return nil, 0, err
	}

	productsEntities := make([]products.Product, len(productsList))
	for i, product := range productsList {
		productsEntities[i] = product.ToProductEntity()
	}

	return productsEntities, uint(totalItems), nil
}

func (pq *ProductQuery) GetAllProducts(limit uint, page uint, search string) ([]products.Product, uint, error) {
	var productsList []Product
	var totalItems int64

	offset := (page - 1) * limit

	qry := pq.db.Model(&Product{}).Preload("Store").Preload("Category")

	if search != "" {
		qry = qry.Where("product_name ILIKE ? OR slug ILIKE ?", "%"+search+"%", "%"+search+"%")
	}

	if err := qry.Count(&totalItems).Error; err != nil {
		return nil, 0, fmt.Errorf("failed to count products: %w", err)
	}

	if err := qry.Limit(int(limit)).Offset(int(offset)).Find(&productsList).Error; err != nil {
		return nil, 0, err
	}

	productsEntities := make([]products.Product, len(productsList))
	for i, product := range productsList {
		productsEntities[i] = product.ToProductEntity()
	}

	return productsEntities, uint(totalItems), nil
}

func (pq *ProductQuery) IsProductOwnedByUser(productID uint, userID uint) (bool, error) {
	var product Product

	err := pq.db.Preload("Store").First(&product, productID).Error
	if err != nil {
		return false, err
	}

	return product.Store.UserID == userID, nil
}
