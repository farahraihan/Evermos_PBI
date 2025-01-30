package handler

import (
	"evermos_pbi/internal/features/stores"
	"time"
)

type StoreResponse struct {
	ID         uint      `json:"id"`
	StoreName  string    `json:"store_name"`
	StoreImage string    `json:"store_image"`
	UserID     uint      `json:"user_id"`
	OwnerName  string    `json:"owner_name"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func ToStoreResponse(input stores.Store) StoreResponse {
	return StoreResponse{
		ID:         input.ID,
		StoreName:  input.StoreName,
		StoreImage: input.StoreImage,
		UserID:     input.UserID,
		OwnerName:  input.OwnerName,
		CreatedAt:  input.CreatedAt,
		UpdatedAt:  input.UpdatedAt,
	}
}

func ToStoreResponses(stores []stores.Store) []StoreResponse {
	responses := make([]StoreResponse, len(stores))
	for i, store := range stores {
		responses[i] = ToStoreResponse(store)
	}
	return responses
}
