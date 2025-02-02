package stores

import (
	"evermos_pbi/internal/features/users"
	"mime/multipart"
	"time"

	"github.com/labstack/echo/v4"
)

type Store struct {
	ID         uint
	StoreName  string
	StoreImage string
	UserID     uint
	OwnerName  string
	User       users.User `gorm:"foreignKey:UserID"`
	CreatedAt  time.Time  `gorm:"default:current_timestamp"`
	UpdatedAt  time.Time  `gorm:"default:current_timestamp"`
}

type SQuery interface {
	AddStore(newStore Store) error
	UpdateStore(userID uint, updateStore Store) error
	DeleteStore(userID uint) error
	GetStoreByUserID(userID uint) (Store, error)
	GetStoreByID(storeID uint) (Store, error)
	GetAllStores(limit uint, page uint, search string) ([]Store, uint, error)
	IsOwnerExist(userID uint) (bool, error)
	IsStoreOwnedByUser(storeID uint, userID uint) (bool, error)
}

type SService interface {
	AddStore(newStore Store, src multipart.File, filename string) error
	UpdateStore(userID uint, updateStore Store, src multipart.File, filename string) error
	DeleteStore(userID uint) error
	GetStoreByUserID(userID uint) (Store, error)
	GetStoreByID(storeID uint) (Store, error)
	GetAllStores(limit uint, page uint, search string) ([]Store, uint, error)
	IsStoreOwnedByUser(storeID uint, userID uint) (bool, error)
}

type SHandler interface {
	AddStore() echo.HandlerFunc
	UpdateStore() echo.HandlerFunc
	DeleteStore() echo.HandlerFunc
	GetStoreByUserID() echo.HandlerFunc
	GetStoreByID() echo.HandlerFunc
	GetAllStores() echo.HandlerFunc
}
