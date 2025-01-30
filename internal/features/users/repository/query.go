package repository

import (
	"evermos_pbi/internal/features/users"

	"gorm.io/gorm"
)

type UserQuery struct {
	db *gorm.DB
}

func NewUserQuery(connect *gorm.DB) users.UQuery {
	return &UserQuery{
		db: connect,
	}
}

func (uq *UserQuery) Login(email string) (users.User, error) {
	var result User
	err := uq.db.Where("email = ?", email).First(&result).Error

	if err != nil {
		return users.User{}, err
	}

	return result.ToUserEntity(), nil
}

func (uq *UserQuery) Register(newUsers *users.User) error {
	cnvData := ToUserQuery(*newUsers)
	err := uq.db.Create(&cnvData).Error

	if err != nil {
		return err
	}

	if cnvData.ID != 0 {
		newUsers.ID = cnvData.ID
	}

	return nil
}

func (uq *UserQuery) UpdateUser(userID uint, updatedUser users.User) error {
	cnvData := ToUserQuery(updatedUser)

	qry := uq.db.Where("id = ?", userID).Updates(&cnvData)

	if qry.Error != nil {
		return qry.Error
	}

	if qry.RowsAffected < 1 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func (uq *UserQuery) DeleteUser(userID uint) error {
	qry := uq.db.Where("id = ?", userID).Delete(&User{})

	if qry.Error != nil {
		return qry.Error
	}

	if qry.RowsAffected < 1 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func (uq *UserQuery) GetUserByID(userID uint) (users.User, error) {
	var user users.User

	err := uq.db.First(&user, userID).Error
	if err != nil {
		return users.User{}, err
	}
	return user, nil
}

func (uq *UserQuery) IsAdmin(userID uint) (bool, error) {
	var user users.User
	if err := uq.db.Where("id = ?", userID).First(&user).Error; err != nil {
		return false, err
	}

	return user.IsAdmin, nil
}

func (uq *UserQuery) IsEmailExist(email string) (bool, error) {
	var count int64
	if err := uq.db.Model(&users.User{}).Where("email = ?", email).Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}

func (uq *UserQuery) IsPhoneExist(phone string) (bool, error) {
	var count int64
	if err := uq.db.Model(&users.User{}).Where("phone = ?", phone).Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}
