package repository

import (
	"evermos_pbi/internal/features/detailtransaction"
	"fmt"

	"gorm.io/gorm"
)

type DetailTransactionQuery struct {
	db *gorm.DB
}

func NewDetailTransactionQuery(connect *gorm.DB) detailtransaction.DQuery {
	return &DetailTransactionQuery{
		db: connect,
	}
}

func (dq *DetailTransactionQuery) AddDetailTransaction(newDetailTransaction detailtransaction.DetailTransaction) error {
	cnvData := ToDetailTransactionQuery(newDetailTransaction)
	qry := dq.db.Create(&cnvData)

	if qry.Error != nil {
		return qry.Error
	}

	return nil
}

func (dq *DetailTransactionQuery) UpdateDetailTransaction(productID uint, transactionID uint, quantity uint) error {
	var existingDetailTransaction DetailTransaction

	err := dq.db.Where("product_id = ? AND transaction_id = ?", productID, transactionID).First(&existingDetailTransaction).Error
	if err != nil {
		return err
	}

	newQuantity := existingDetailTransaction.Quantity + quantity

	qry := dq.db.Model(&DetailTransaction{}).Where("product_id = ? AND transaction_id = ?", productID, transactionID).Update("quantity", newQuantity)

	if qry.Error != nil {
		return qry.Error
	}

	if qry.RowsAffected < 1 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func (dq *DetailTransactionQuery) DeleteDetailTransaction(ProductID uint, TransactionID uint) error {
	qry := dq.db.Where("product_id = ? AND transaction_id = ?", ProductID, TransactionID).Delete(&DetailTransaction{})

	if qry.Error != nil {
		return qry.Error
	}

	if qry.RowsAffected < 1 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func (dq *DetailTransactionQuery) IsProductInDetail(productID uint, transactionID uint) (bool, error) {
	var count int64
	err := dq.db.Model(&DetailTransaction{}).Where("product_id = ? AND transaction_id = ?", productID, transactionID).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (dq *DetailTransactionQuery) GetDetailTransactions(transactionID uint) ([]detailtransaction.DetailTransaction, error) {
	var detailTransactionList []DetailTransaction

	if err := dq.db.Where("transaction_id = ?", transactionID).Find(&detailTransactionList).Error; err != nil {
		return nil, fmt.Errorf("failed to get detail transactions: %w", err)
	}

	detailEntities := make([]detailtransaction.DetailTransaction, len(detailTransactionList))
	for i, detailtransaction := range detailTransactionList {
		detailEntities[i] = detailtransaction.ToDetailTransactionEntity()
	}

	return detailEntities, nil
}
