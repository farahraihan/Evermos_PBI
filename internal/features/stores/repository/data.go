package repository

import (
	"evermos_pbi/internal/features/stores"
	"evermos_pbi/internal/features/users"

	"gorm.io/gorm"
)

type Store struct {
	gorm.Model
	StoreName  string
	StoreImage string
	UserID     uint
	User       users.User `gorm:"foreignKey:UserID"`
}

func (s *Store) ToStoreEntity() stores.Store {
	return stores.Store{
		ID:         s.ID,
		StoreName:  s.StoreName,
		StoreImage: s.StoreImage,
		UserID:     s.UserID,
		OwnerName:  s.User.Name,
		CreatedAt:  s.CreatedAt,
		UpdatedAt:  s.UpdatedAt,
	}
}

func ToStoreQuery(input stores.Store) Store {
	return Store{
		StoreName:  input.StoreName,
		StoreImage: input.StoreImage,
		UserID:     input.UserID,
	}
}
