package repository

import (
	"evermos_pbi/internal/features/logproduct"
	"fmt"

	"gorm.io/gorm"
)

type LogProductQuery struct {
	db *gorm.DB
}

func NewLogProductQuery(connect *gorm.DB) logproduct.LQuery {
	return &LogProductQuery{
		db: connect,
	}
}

func (lq *LogProductQuery) AddLogProduct(newLogProduct logproduct.LogProduct) error {
	cnvData := ToLogProductQuery(newLogProduct)
	qry := lq.db.Create(&cnvData)

	if qry.Error != nil {
		return qry.Error
	}

	return nil
}

func (lq *LogProductQuery) GetLogProductByID(logProductID uint) (logproduct.LogProduct, error) {
	var logProduct logproduct.LogProduct

	err := lq.db.First(&logProduct, logProductID).Error
	if err != nil {
		return logproduct.LogProduct{}, err
	}
	return logProduct, nil
}

func (pq *LogProductQuery) GetAllLogProduct(limit uint, page uint, search string) ([]logproduct.LogProduct, uint, error) {
	var logProductsList []LogProduct
	var totalItems int64

	offset := (page - 1) * limit

	qry := pq.db.Model(&LogProduct{})

	if search != "" {
		qry = qry.Where("product_name ILIKE ? OR slug ILIKE ?", "%"+search+"%", "%"+search+"%")
	}

	if err := qry.Count(&totalItems).Error; err != nil {
		return nil, 0, fmt.Errorf("failed to count log products: %w", err)
	}

	if err := qry.Limit(int(limit)).Offset(int(offset)).Find(&logProductsList).Error; err != nil {
		return nil, 0, err
	}

	logProductsEntities := make([]logproduct.LogProduct, len(logProductsList))
	for i, logProduct := range logProductsList {
		logProductsEntities[i] = logProduct.ToLogProductEntity()
	}

	return logProductsEntities, uint(totalItems), nil
}
