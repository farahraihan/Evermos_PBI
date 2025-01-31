package repository

import (
	"evermos_pbi/internal/features/address"
	"evermos_pbi/internal/features/users"

	"gorm.io/gorm"
)

type Address struct {
	gorm.Model
	RcpName  string
	Phone    string
	Detail   string
	Province string
	Regency  string
	District string
	Village  string
	UserID   uint
	User     users.User `gorm:"foreignKey:UserID"`
}

func (a *Address) ToAddressEntity() address.Address {
	return address.Address{
		ID:        a.ID,
		RcpName:   a.RcpName,
		Phone:     a.Phone,
		Detail:    a.Detail,
		Province:  a.Province,
		Regency:   a.Regency,
		District:  a.District,
		Village:   a.Village,
		UserID:    a.UserID,
		CreatedAt: a.CreatedAt,
		UpdatedAt: a.UpdatedAt,
	}
}

func ToAddressQuery(input address.Address) Address {
	return Address{
		RcpName:  input.RcpName,
		Phone:    input.Phone,
		Detail:   input.Detail,
		Province: input.Province,
		Regency:  input.Regency,
		District: input.District,
		Village:  input.Village,
		UserID:   input.UserID,
	}
}
