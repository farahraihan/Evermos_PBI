package handler

import (
	"evermos_pbi/internal/features/users"
	"evermos_pbi/internal/helpers"
	"evermos_pbi/internal/utils"
	"log"
	"mime/multipart"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	srv users.UService
	tu  utils.JwtUtilityInterface
}

func NewUserHandler(s users.UService, t utils.JwtUtilityInterface) users.UHandler {
	return &UserHandler{
		srv: s,
		tu:  t,
	}
}

func (uh *UserHandler) Register() echo.HandlerFunc {
	return func(c echo.Context) error {
		image, err := c.FormFile("user_image")
		if err != nil {
			log.Println("failed to get image file")
			return c.JSON(http.StatusBadRequest, helpers.ResponseFormat(http.StatusBadRequest, "failed", "invalid image file", nil))
		}

		src, err := image.Open()
		if err != nil {
			log.Println("failed to open image file")
			return c.JSON(http.StatusInternalServerError, helpers.ResponseFormat(http.StatusInternalServerError, "failed", "unable to process the image", nil))
		}
		defer src.Close()

		var input RegisterOrUpdateRequest

		err = c.Bind(&input)
		if err != nil {
			log.Println("failed to bind register request")
			return c.JSON(http.StatusBadRequest, helpers.ResponseFormat(http.StatusBadRequest, "failed", "invalid input", nil))
		}

		parsedDate, err := time.Parse("2006-01-02", input.BirthDate)
		if err != nil {
			log.Println("failed to parse birth date:", err)
			return c.JSON(http.StatusBadRequest, helpers.ResponseFormat(http.StatusBadRequest, "failed", "invalid date format, use YYYY-MM-DD", nil))
		}

		exist, _ := uh.srv.IsEmailExist(input.Email)
		if exist {
			log.Println("email has been registered")
			return c.JSON(http.StatusBadRequest, helpers.ResponseFormat(http.StatusBadRequest, "failed", "email has been registered", nil))
		}

		exist, _ = uh.srv.IsPhoneExist(input.Phone)
		if exist {
			log.Println("phone has been registered")
			return c.JSON(http.StatusBadRequest, helpers.ResponseFormat(http.StatusBadRequest, "failed", "phone has been registered", nil))
		}

		newUser := RegisterToUser(input, parsedDate)

		err = uh.srv.Register(newUser, src, image.Filename)
		if err != nil {
			log.Print("failed to register user")
			return c.JSON(http.StatusInternalServerError, helpers.ResponseFormat(http.StatusInternalServerError, "failed", "registration unsuccessful", nil))
		}

		return c.JSON(http.StatusCreated, helpers.ResponseFormat(http.StatusCreated, "success", "user registration successful", nil))

	}
}

func (uh *UserHandler) Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input LoginRequest

		err := c.Bind(&input)
		if err != nil {
			log.Print("failed to bind login request")
			return c.JSON(http.StatusBadRequest, helpers.ResponseFormat(http.StatusBadRequest, "failed", "invalid login input", nil))
		}

		token, err := uh.srv.Login(input.Email, input.Password)
		if err != nil {
			log.Print("login attempt failed")
			return c.JSON(http.StatusInternalServerError, helpers.ResponseFormat(http.StatusInternalServerError, "failed", "login unsuccessful", nil))
		}

		return c.JSON(http.StatusOK, helpers.ResponseFormat(http.StatusOK, "success", "login successful", ToLoginResponse(token)))
	}
}

func (uh *UserHandler) UpdateUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		userID := uh.tu.DecodToken(c.Get("user").(*jwt.Token))
		if userID == 0 {
			log.Println("unauthorized attempt to update user")
			return c.JSON(http.StatusBadRequest, helpers.ResponseFormat(http.StatusUnauthorized, "failed", "unauthorized", nil))
		}

		var src multipart.File
		var filename string
		image, err := c.FormFile("image")
		if err == nil {
			src, err = image.Open()
			if err != nil {
				log.Println("failed to open image file")
				return c.JSON(http.StatusInternalServerError, helpers.ResponseFormat(http.StatusInternalServerError, "failed", "unable to process image", nil))
			}
			defer src.Close()
			filename = image.Filename
		}

		var req RegisterOrUpdateRequest
		err = c.Bind(&req)
		if err != nil {
			log.Println("failed to bind update request")
			return c.JSON(http.StatusBadRequest, helpers.ResponseFormat(http.StatusBadRequest, "failed", "invalid input", nil))
		}

		parsedDate, err := time.Parse("2006-01-02", req.BirthDate)
		if err != nil {
			log.Println("failed to parse birth date:", err)
			return c.JSON(http.StatusBadRequest, helpers.ResponseFormat(
				http.StatusBadRequest, "failed", "invalid date format, use YYYY-MM-DD", nil))
		}

		updateUser := RegisterToUser(req, parsedDate)

		exist, _ := uh.srv.IsEmailExist(updateUser.Email)
		if exist {
			log.Println("email has been registered")
			return c.JSON(http.StatusBadRequest, helpers.ResponseFormat(http.StatusBadRequest, "failed", "email has been registered", nil))
		}

		exist, _ = uh.srv.IsPhoneExist(updateUser.Phone)
		if exist {
			log.Println("phone has been registered")
			return c.JSON(http.StatusBadRequest, helpers.ResponseFormat(http.StatusBadRequest, "failed", "phone has been registered", nil))
		}

		err = uh.srv.UpdateUser(uint(userID), updateUser, src, filename)
		if err != nil {
			log.Println("failed to update user")
			return c.JSON(http.StatusInternalServerError, helpers.ResponseFormat(http.StatusInternalServerError, "failed", "update unsuccessful", nil))
		}

		return c.JSON(http.StatusOK, helpers.ResponseFormat(http.StatusOK, "success", "user was successfully updated", nil))
	}
}

func (uh *UserHandler) DeleteUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		userID := uh.tu.DecodToken(c.Get("user").(*jwt.Token))
		if userID == 0 {
			log.Println("unauthorized attempt to delete user")
			return c.JSON(http.StatusUnauthorized, helpers.ResponseFormat(http.StatusUnauthorized, "failed", "unauthorized", nil))
		}

		err := uh.srv.DeleteUser(uint(userID))
		if err != nil {
			log.Println("failed to delete user")
			return c.JSON(http.StatusInternalServerError, helpers.ResponseFormat(http.StatusInternalServerError, "failed", "deletion unsuccessful", nil))
		}

		return c.JSON(http.StatusOK, helpers.ResponseFormat(http.StatusOK, "success", "your account has been deleted", nil))
	}
}

func (uh *UserHandler) GetUserByID() echo.HandlerFunc {
	return func(c echo.Context) error {

		userID := uh.tu.DecodToken(c.Get("user").(*jwt.Token))
		if userID == 0 {
			log.Println("unauthorized attempt to get user")
			return c.JSON(http.StatusUnauthorized, helpers.ResponseFormat(http.StatusUnauthorized, "failed", "unauthorized", nil))
		}

		user, err := uh.srv.GetUserByID(uint(userID))
		if err != nil {
			log.Println("failed to get user by ID")
			return c.JSON(http.StatusInternalServerError, helpers.ResponseFormat(http.StatusInternalServerError, "failed", "failed to retrieve user", nil))
		}

		UserResponse := ToUserResponse(user)

		return c.JSON(http.StatusOK, helpers.ResponseFormat(http.StatusOK, "success", "user was successfully retrieved", UserResponse))

	}
}
