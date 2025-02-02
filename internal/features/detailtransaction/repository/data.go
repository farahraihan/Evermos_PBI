package repository

import (
	"evermos_pbi/internal/features/detailtransaction"
	"evermos_pbi/internal/features/products"
	"evermos_pbi/internal/features/stores"
	"evermos_pbi/internal/features/transaction"

	"gorm.io/gorm"
)

type DetailTransaction struct {
	gorm.Model
	Quantity      uint
	StoreID       uint
	ProductID     uint
	TransactionID uint
	Store         stores.Store            `gorm:"foreignKey:StoreID"`
	Product       products.Product        `gorm:"foreignKey:ProductID"`
	Transaction   transaction.Transaction `gorm:"foreignKey:TransactionID"`
}

func (d *DetailTransaction) ToDetailTransactionEntity() detailtransaction.DetailTransaction {
	return detailtransaction.DetailTransaction{
		ID:            d.ID,
		Quantity:      d.Quantity,
		StoreID:       d.StoreID,
		ProductID:     d.ProductID,
		TransactionID: d.TransactionID,
		CreatedAt:     d.CreatedAt,
		UpdatedAt:     d.UpdatedAt,
	}
}

func ToDetailTransactionQuery(input detailtransaction.DetailTransaction) DetailTransaction {
	return DetailTransaction{
		Quantity:      input.Quantity,
		StoreID:       input.StoreID,
		ProductID:     input.ProductID,
		TransactionID: input.TransactionID,
	}
}
