package handler

import "evermos_pbi/internal/features/transaction"

type AddTransactionRequest struct {
	Status    string `json:"status" form:"status"`
	UserID    uint   `json:"user_id" form:"user_id"`
	AddressID uint   `json:"address_id" form:"address_id"`
	Quantity  uint   `json:"quantity" form:"quantity"`
	StoreID   uint   `json:"store_id" form:"store_id"`
	ProductID uint   `json:"product_id" form:"product_id"`
}

func AddToTransaction(tr AddTransactionRequest) transaction.Transaction {
	return transaction.Transaction{
		Status:    tr.Status,
		UserID:    tr.UserID,
		AddressID: tr.AddressID,
	}
}

func AddToDetailTransaction(tr AddTransactionRequest) transaction.DetailTransaction2 {
	return transaction.DetailTransaction2{
		Quantity:  tr.Quantity,
		StoreID:   tr.StoreID,
		ProductID: tr.ProductID,
	}
}
