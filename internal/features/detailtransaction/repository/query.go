package repository

import (
	"evermos_pbi/internal/features/detailtransaction"

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

func (dq *DetailTransactionQuery) UpdateDetailTransaction(ProductID uint, TransactionID uint, updateDetailTransaction detailtransaction.DetailTransaction) error {
	cnvData := ToDetailTransactionQuery(updateDetailTransaction)
	qry := dq.db.Where("product_id = ? AND transaction_id = ?", ProductID, TransactionID).Updates(&cnvData)

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

func (dq *DetailTransactionQuery) IsProductInDetail(ProductID uint) (bool, error) {
	var count int64
	err := dq.db.Model(&DetailTransaction{}).Where("product_id = ?", ProductID).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}
