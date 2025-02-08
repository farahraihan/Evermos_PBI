package repository

import (
	"evermos_pbi/internal/features/address"
	"evermos_pbi/internal/features/transaction"
	"evermos_pbi/internal/features/users"

	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	Status    string
	UserID    uint
	AddressID uint
	User      users.User      `gorm:"foreignKey:UserID"`
	Address   address.Address `gorm:"foreignKey:AddressID"`
}

func (t *Transaction) ToTransactionEntity() transaction.Transaction {
	return transaction.Transaction{
		ID:        t.ID,
		Status:    t.Status,
		UserID:    t.UserID,
		AddressID: t.AddressID,
		CreatedAt: t.CreatedAt,
		UpdatedAt: t.UpdatedAt,
	}
}

func ToTransactionQuery(input transaction.Transaction) Transaction {
	return Transaction{
		Status:    input.Status,
		UserID:    input.UserID,
		AddressID: input.AddressID,
	}
}
