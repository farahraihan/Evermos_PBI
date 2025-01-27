package handler

import (
	"evermos_pbi/internal/features/users"
	"time"
)

type LoginResponse struct {
	Token string `json:"token"`
}

type UserResponse struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Password  string    `json:"password"`
	Phone     string    `json:"phone"`
	BirthDate time.Time `json:"birth_date"`
	Gender    string    `json:"gender"`
	UserImage string    `json:"user_image"`
	About     string    `json:"about"`
	Job       string    `json:"job"`
	Email     string    `json:"email"`
	IsAdmin   bool      `json:"is_admin"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func ToLoginResponse(token string) LoginResponse {
	return LoginResponse{
		Token: token,
	}
}

func ToUserResponse(input users.User) UserResponse {
	return UserResponse{
		ID:        input.ID,
		Name:      input.Name,
		Phone:     input.Phone,
		BirthDate: input.BirthDate,
		Gender:    input.Gender,
		UserImage: input.UserImage,
		About:     input.About,
		Job:       input.Job,
		Email:     input.Email,
		IsAdmin:   input.IsAdmin,
		CreatedAt: input.CreatedAt,
		UpdatedAt: input.UpdatedAt,
	}
}
