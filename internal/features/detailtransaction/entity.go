package detailtransaction

import (
	"evermos_pbi/internal/features/products"
	"evermos_pbi/internal/features/stores"
	"evermos_pbi/internal/features/transaction"
	"time"
)

type DetailTransaction struct {
	ID            uint
	Quantity      uint
	StoreID       uint
	ProductID     uint
	TransactionID uint
	Store         stores.Store            `gorm:"foreignKey:StoreID"`
	Product       products.Product        `gorm:"foreignKey:ProductID"`
	Transaction   transaction.Transaction `gorm:"foreignKey:TransactionID"`
	CreatedAt     time.Time               `gorm:"default:current_timestamp"`
	UpdatedAt     time.Time               `gorm:"default:current_timestamp"`
}

type DQuery interface {
	AddDetailTransaction(newDetailTransaction DetailTransaction) error
	UpdateDetailTransaction(productID uint, transactionID uint, quantity uint) error
	DeleteDetailTransaction(productID uint, transactionID uint) error
	IsProductInDetail(productID uint, transactionID uint) (bool, error)
}

type DService interface {
	AddDetailTransaction(newDetailTransaction DetailTransaction) error
	UpdateDetailTransaction(productID uint, transactionID uint, quantity uint) error
	DeleteDetailTransaction(productID uint, transactionID uint) error
	IsProductInDetail(productID uint, transactionID uint) (bool, error)
}
