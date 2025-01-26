package handler

import (
	"evermos_pbi/internal/features/users"
	"time"
)

type LoginRequest struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type RegisterOrUpdateRequest struct {
	Name      string `json:"name" form:"name"`
	Password  string `json:"password" form:"password"`
	Phone     string `json:"phone" form:"phone"`
	BirthDate string `json:"birth_date" form:"birth_date"`
	Gender    string `json:"gender" form:"gender"`
	UserImage string `json:"user_image" form:"user_image"`
	About     string `json:"about" form:"about"`
	Job       string `json:"job" form:"job"`
	Email     string `json:"email" form:"email"`
}

func RegisterToUser(ur RegisterOrUpdateRequest, BirthDate time.Time) users.User {
	return users.User{
		Name:      ur.Name,
		Password:  ur.Password,
		Phone:     ur.Phone,
		BirthDate: BirthDate,
		Gender:    ur.Gender,
		UserImage: ur.UserImage,
		About:     ur.About,
		Job:       ur.Job,
		Email:     ur.Email,
	}
}
