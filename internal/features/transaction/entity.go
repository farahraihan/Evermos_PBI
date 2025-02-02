package transaction

import (
	"evermos_pbi/internal/features/address"
	"evermos_pbi/internal/features/products"
	"evermos_pbi/internal/features/stores"
	"evermos_pbi/internal/features/users"
	"time"

	"github.com/labstack/echo/v4"
)

type Transaction struct {
	ID        uint
	Status    string
	UserID    uint
	AddressID uint
	User      users.User      `gorm:"foreignKey:UserID"`
	Address   address.Address `gorm:"foreignKey:AddressID"`
	CreatedAt time.Time       `gorm:"default:current_timestamp"`
	UpdatedAt time.Time       `gorm:"default:current_timestamp"`
}

type DetailTransaction2 struct {
	ID          uint
	Quantity    uint
	StoreID     uint
	StoreName   string
	ProductID   uint
	ProductName string
	TotalPrice  string
	Store       stores.Store     `gorm:"foreignKey:StoreID"`
	Product     products.Product `gorm:"foreignKey:ProductID"`
	CreatedAt   time.Time        `gorm:"default:current_timestamp"`
	UpdatedAt   time.Time        `gorm:"default:current_timestamp"`
}

type TransactionWithDetail struct {
	Transaction       Transaction
	DetailTransaction []DetailTransaction2
}

type TQuery interface {
	AddTransaction(newTransaction *Transaction) error
	UpdateDetailTransaction(transactionID uint, productID uint) error
	DeleteTransaction(transactionID uint, productID uint) error
	UpdateTransaction(transactionID uint) error
	GetTransactionByStatusCart(transactionID uint) (TransactionWithDetail, error)
	GetTransactionHistory(userID uint, limit uint, page uint) ([]TransactionWithDetail, uint, error)
	GetTransactionByID(transactionID uint) (TransactionWithDetail, error)
}

type TService interface {
	AddTransaction(newTransaction Transaction, TransactionID uint, newDetailTransaction DetailTransaction2) error
	UpdateDetailTransaction(userID uint, transactionID uint, productID uint) error
	DeleteTransaction(userID uint, transactionID uint, productID uint) error
	UpdateTransaction(userID uint, transactionID uint) error
	GetTransactionByStatusCart(userID uint, transactionID uint) (TransactionWithDetail, error)
	GetTransactionHistory(userID uint, limit uint, page uint) ([]TransactionWithDetail, uint, error)
	GetTransactionByID(userID uint, transactionID uint) (TransactionWithDetail, error)
}

type THandler interface {
	AddTransaction() echo.HandlerFunc
	UpdateDetailTransaction() echo.HandlerFunc
	DeleteTransaction() echo.HandlerFunc
	UpdateTransactionStatusCanceled() echo.HandlerFunc
	UpdateTransactionStatusCompleted() echo.HandlerFunc
	GetTransactionByStatusCart() echo.HandlerFunc
	GetTransactionHistory() echo.HandlerFunc
	GetTransactionByID() echo.HandlerFunc
}
