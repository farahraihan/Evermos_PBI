package users

import (
	"mime/multipart"
	"time"

	"github.com/labstack/echo/v4"
)

type User struct {
	ID        uint
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
	CreatedAt time.Time `gorm:"default:current_timestamp"`
	UpdatedAt time.Time `gorm:"default:current_timestamp"`
}

type UQuery interface {
	Login(email string) (User, error)
	Register(newUser *User) error
	UpdateUser(userID uint, updateUser User) error
	DeleteUser(userID uint) error
	GetUserByID(userID uint) (User, error)
	IsAdmin(userID uint) (bool, error)
	IsEmailExist(email string) (bool, error)
	IsPhoneExist(phone string) (bool, error)
}

type UService interface {
	Login(email string, password string) (string, error)
	Register(newUser User, src multipart.File, filename string) error
	UpdateUser(userID uint, updatedUser User, src multipart.File, filename string) error
	DeleteUser(userID uint) error
	GetUserByID(userID uint) (User, error)
	IsAdmin(userID uint) (bool, error)
	IsEmailExist(email string) (bool, error)
	IsPhoneExist(phone string) (bool, error)
}

type UHandler interface {
	Login() echo.HandlerFunc
	Register() echo.HandlerFunc
	UpdateUser() echo.HandlerFunc
	DeleteUser() echo.HandlerFunc
	GetUserByID() echo.HandlerFunc
}
