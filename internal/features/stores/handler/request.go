package handler

import "evermos_pbi/internal/features/stores"

type AddOrUpdateStoreRequest struct {
	StoreName  string `json:"name" form:"store_name"`
	StoreImage string `json:"store_image" form:"store_image"`
	UserID     uint   `json:"userID" form:"user_id"`
}

func AddToStore(sr AddOrUpdateStoreRequest) stores.Store {
	return stores.Store{
		StoreName:  sr.StoreName,
		StoreImage: sr.StoreImage,
		UserID:     sr.UserID,
	}
}
