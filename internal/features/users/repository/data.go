package repository

import (
	"evermos_pbi/internal/features/users"
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name      string
	Password  string
	Phone     string
	BirthDate time.Time
	Gender    string
	UserImage string
	About     string
	Job       string
	Email     string
	IsAdmin   bool
}

func (u *User) ToUserEntity() users.User {
	return users.User{
		ID:        u.ID,
		Name:      u.Name,
		Password:  u.Password,
		Phone:     u.Phone,
		BirthDate: u.BirthDate,
		Gender:    u.Gender,
		UserImage: u.UserImage,
		About:     u.About,
		Job:       u.About,
		Email:     u.Email,
		IsAdmin:   u.IsAdmin,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}

func ToUserQuery(input users.User) User {
	return User{
		Name:      input.Name,
		Password:  input.Password,
		Phone:     input.Phone,
		BirthDate: input.BirthDate,
		Gender:    input.Gender,
		UserImage: input.UserImage,
		About:     input.About,
		Job:       input.Job,
		Email:     input.Email,
		IsAdmin:   input.IsAdmin,
	}
}
