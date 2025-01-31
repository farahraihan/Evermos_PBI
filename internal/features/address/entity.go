package address

import (
	"evermos_pbi/internal/features/users"
	"time"

	"github.com/labstack/echo/v4"
)

type Address struct {
	ID        uint
	RcpName   string
	Phone     string
	Detail    string
	Province  string
	Regency   string
	District  string
	Village   string
	UserID    uint
	User      users.User `gorm:"foreignKey:UserID"`
	CreatedAt time.Time  `gorm:"default:current_timestamp"`
	UpdatedAt time.Time  `gorm:"default:current_timestamp"`
}

type AQuery interface {
	AddAddress(newAddress Address) error
	UpdateAddress(userID uint, addressID uint, updateAddress Address) error
	DeleteAddress(userID uint, addressID uint) error
	GetAddressByUserID(userID uint) ([]Address, error)
}

type AService interface {
	GetProvince() ([]map[string]interface{}, error)
	GetRegency(provinceID uint) ([]map[string]interface{}, error)
	GetDistrict(regenciesID uint) ([]map[string]interface{}, error)
	GetVillage(districtsID uint) ([]map[string]interface{}, error)
	AddAddress(newAddress Address) error
	UpdateAddress(userID uint, addressID uint, updateAddress Address) error
	DeleteAddress(userID uint, addressID uint) error
	GetAddressByUserID(userID uint) ([]Address, error)
}

type AHandler interface {
	GetProvince() echo.HandlerFunc
	GetRegency() echo.HandlerFunc
	GetDistrict() echo.HandlerFunc
	GetVillage() echo.HandlerFunc
	AddAddress() echo.HandlerFunc
	UpdateAddress() echo.HandlerFunc
	DeleteAddress() echo.HandlerFunc
	GetAddressByUserID() echo.HandlerFunc
}
