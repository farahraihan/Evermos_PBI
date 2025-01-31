package repository

import (
	"evermos_pbi/internal/features/address"
	"fmt"

	"gorm.io/gorm"
)

type AddressQuery struct {
	db *gorm.DB
}

func NewAddressQuery(connect *gorm.DB) address.AQuery {
	return &AddressQuery{
		db: connect,
	}
}

func (aq *AddressQuery) AddAddress(newAddress address.Address) error {
	cnvData := ToAddressQuery(newAddress)
	qry := aq.db.Create(&cnvData)

	if qry.Error != nil {
		return qry.Error
	}

	return nil
}

func (aq *AddressQuery) UpdateAddress(userID uint, addressID uint, updateAddress address.Address) error {
	var existingAddress address.Address

	err := aq.db.Select("user_id").Where("id = ?", addressID).First(&existingAddress).Error
	if err != nil {
		return err
	}

	if existingAddress.UserID != userID {
		return fmt.Errorf("user does not own this address")
	}

	cnvData := ToAddressQuery(updateAddress)

	qry := aq.db.Where("id = ?", addressID).Updates(&cnvData)

	if qry.Error != nil {
		return qry.Error
	}

	if qry.RowsAffected < 1 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func (aq *AddressQuery) DeleteAddress(userID uint, addressID uint) error {
	var existingAddress address.Address

	err := aq.db.Select("user_id").Where("id = ?", addressID).First(&existingAddress).Error
	if err != nil {
		return err
	}

	if existingAddress.UserID != userID {
		return fmt.Errorf("user does not own this address")
	}

	qry := aq.db.Where("id = ?", addressID).Delete(&Address{})

	if qry.Error != nil {
		return err
	}

	if qry.RowsAffected < 1 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func (aq *AddressQuery) GetAddressByUserID(userID uint) ([]address.Address, error) {
	var addressList []Address

	//get address list from database by user id
	err := aq.db.Where("user_id = ?", userID).Find(&addressList).Error
	if err != nil {
		return nil, err
	}

	addressEntities := make([]address.Address, len(addressList))
	for i, address := range addressList {
		addressEntities[i] = address.ToAddressEntity()
	}

	return addressEntities, nil

}
