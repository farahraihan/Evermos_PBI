package repository

import (
	"evermos_pbi/internal/features/detailtransaction"
	"evermos_pbi/internal/features/transaction"
	"fmt"

	"gorm.io/gorm"
)

type TransactionQuery struct {
	db *gorm.DB
}

func NewTransactionQuery(connect *gorm.DB) transaction.TQuery {
	return &TransactionQuery{
		db: connect,
	}
}

func (tq *TransactionQuery) AddTransaction(newTransaction *transaction.Transaction) error {
	cnvData := ToTransactionQuery(*newTransaction)
	qry := tq.db.Create(&cnvData)

	if qry.Error != nil {
		return qry.Error
	}

	if cnvData.ID != 0 {
		newTransaction.ID = cnvData.ID
	}

	return nil
}

func (tq *TransactionQuery) UpdateTransaction(transactionID uint, status string) error {
	qry := tq.db.Model(&Transaction{}).Where("id = ?", transactionID).Update("status", status)

	if qry.Error != nil {
		return qry.Error
	}

	if qry.RowsAffected < 1 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func (tq *TransactionQuery) GetTransactionByStatusCart(transactionID uint) (transaction.TransactionWithDetail, error) {
	var trx transaction.Transaction
	var details []detailtransaction.DetailTransaction
	var convertedDetails []transaction.DetailTransaction2

	// Ambil transaksi dengan status "cart"
	qry := tq.db.Where("id = ? AND status = ?", transactionID, "cart").Preload("User").Preload("Address").First(&trx)
	if qry.Error != nil {
		return transaction.TransactionWithDetail{}, qry.Error
	}

	detailQry := tq.db.Where("transaction_id = ?", transactionID).Find(&details)
	if detailQry.Error != nil {
		return transaction.TransactionWithDetail{}, detailQry.Error
	}

	for _, d := range details {
		totalPrice := (d.Product.ConsumenPrice) * float32(d.Quantity)
		convertedDetails = append(convertedDetails, transaction.DetailTransaction2{
			Quantity:    d.Quantity,
			StoreID:     d.StoreID,
			StoreName:   d.Store.StoreName,
			ProductID:   d.ProductID,
			ProductName: d.Product.ProductName,
			TotalPrice:  totalPrice,
		})
	}

	return transaction.TransactionWithDetail{
		Trx:               trx,
		DetailTransaction: convertedDetails,
	}, nil
}

func (tq *TransactionQuery) GetTransactionHistory(userID uint, limit uint, page uint) ([]transaction.TransactionWithDetail, uint, error) {
	var transactions []transaction.Transaction
	var details []detailtransaction.DetailTransaction
	var convertedDetails []transaction.DetailTransaction2
	var totalCount int64

	countQry := tq.db.Model(&transaction.Transaction{}).Where("user_id = ? AND (status = ? OR status = ?)", userID, "canceled", "completed").Count(&totalCount)
	if countQry.Error != nil {
		return nil, 0, countQry.Error
	}

	offset := (page - 1) * limit
	qry := tq.db.Where("user_id = ? AND (status = ? OR status = ?)", userID, "canceled", "completed").
		Limit(int(limit)).
		Offset(int(offset)).
		Preload("User").
		Preload("Address").
		Find(&transactions)

	if qry.Error != nil {
		return nil, 0, qry.Error
	}

	var transactionsWithDetails []transaction.TransactionWithDetail

	for _, trx := range transactions {
		detailQry := tq.db.Where("transaction_id = ?", trx.ID).Preload("Store").Preload("Product").Find(&details)
		if detailQry.Error != nil {
			return nil, 0, detailQry.Error
		}

		for _, d := range details {
			totalPrice := (d.Product.ConsumenPrice) * float32(d.Quantity)
			convertedDetails = append(convertedDetails, transaction.DetailTransaction2{
				Quantity:    d.Quantity,
				StoreID:     d.StoreID,
				StoreName:   d.Store.StoreName,
				ProductID:   d.ProductID,
				ProductName: d.Product.ProductName,
				TotalPrice:  totalPrice,
			})
		}

		transactionsWithDetails = append(transactionsWithDetails, transaction.TransactionWithDetail{
			Trx:               trx,
			DetailTransaction: convertedDetails,
		})
	}

	return transactionsWithDetails, uint(totalCount), nil
}

func (tq *TransactionQuery) GetTransactionByID(transactionID uint) (transaction.TransactionWithDetail, error) {
	var trx transaction.Transaction
	var details []detailtransaction.DetailTransaction
	var convertedDetails []transaction.DetailTransaction2

	qry := tq.db.Where("id = ?", transactionID).Preload("User").Preload("Address").First(&trx)
	if qry.Error != nil {
		if qry.Error == gorm.ErrRecordNotFound {
			return transaction.TransactionWithDetail{}, fmt.Errorf("transaction not found")
		}
		return transaction.TransactionWithDetail{}, qry.Error
	}

	detailQry := tq.db.Where("transaction_id = ?", transactionID).Preload("Store").Preload("Product").Find(&details)
	if detailQry.Error != nil {
		return transaction.TransactionWithDetail{}, detailQry.Error
	}

	for _, d := range details {
		totalPrice := (d.Product.ConsumenPrice * float32(d.Quantity))
		convertedDetails = append(convertedDetails, transaction.DetailTransaction2{
			Quantity:    d.Quantity,
			StoreID:     d.StoreID,
			StoreName:   d.Store.StoreName,
			ProductID:   d.ProductID,
			ProductName: d.Product.ProductName,
			TotalPrice:  totalPrice,
		})
	}

	return transaction.TransactionWithDetail{
		Trx:               trx,
		DetailTransaction: convertedDetails,
	}, nil
}

func (tq *TransactionQuery) CheckTransactionInCart(userID uint) (uint, error) {
	var trx transaction.Transaction

	qry := tq.db.Where("user_id = ? AND status = ?", userID, "cart").First(&trx)

	if qry.Error != nil {
		if qry.Error == gorm.ErrRecordNotFound {
			return 0, nil
		}
		return 0, qry.Error
	}

	return trx.ID, nil
}

func (tq *TransactionQuery) IsTransactionOwner(userID uint, transactionID uint) (bool, error) {
	var count int64
	err := tq.db.Model(&Transaction{}).Where("id = ? AND user_id = ?", transactionID, userID).Count(&count).Error
	if err != nil {
		return false, err
	}

	return count > 0, nil
}
