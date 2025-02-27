package handler

import (
	"evermos_pbi/internal/features/transaction"
	"time"
)

type TransactionResponse struct {
	ID        uint      `json:"transaction_id"`
	UserID    uint      `json:"user_id"`
	AddressID uint      `json:"address_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type DetailTransactionResponse struct {
	Quantity    uint    `json:"quantity"`
	StoreID     uint    `json:"store_id"`
	StoreName   string  `json:"store_name"`
	ProductID   uint    `json:"product_id"`
	ProductName string  `json:"product_name"`
	TotalPrice  float32 `json:"total_price"`
}

type TransactionWithDetailResponse struct {
	Trx               TransactionResponse         `json:"transaction"`
	DetailTransaction []DetailTransactionResponse `json:"detail_transaction"`
}

func ToTransactionResponse(input transaction.TransactionWithDetail) TransactionWithDetailResponse {
	trxResponse := TransactionResponse{
		ID:        input.Trx.ID,
		UserID:    input.Trx.UserID,
		AddressID: input.Trx.AddressID,
		CreatedAt: input.Trx.CreatedAt,
		UpdatedAt: input.Trx.UpdatedAt,
	}

	var detailResponses []DetailTransactionResponse
	for _, detail := range input.DetailTransaction {
		detailResponse := DetailTransactionResponse{
			Quantity:    detail.Quantity,
			StoreID:     detail.StoreID,
			StoreName:   detail.StoreName,
			ProductID:   detail.ProductID,
			ProductName: detail.ProductName,
			TotalPrice:  detail.TotalPrice,
		}
		detailResponses = append(detailResponses, detailResponse)
	}

	// Mengembalikan response yang sudah dikonversi
	return TransactionWithDetailResponse{
		Trx:               trxResponse,
		DetailTransaction: detailResponses,
	}
}

func ToTransactionHistoryResponses(input []transaction.TransactionWithDetail) []TransactionWithDetailResponse {
	responses := make([]TransactionWithDetailResponse, len(input))
	for i, transaction := range input {
		responses[i] = ToTransactionResponse(transaction)
	}
	return responses
}
