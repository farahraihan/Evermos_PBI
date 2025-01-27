package service

import (
	"errors"
	"evermos_pbi/internal/features/users"
	"evermos_pbi/internal/utils"
	"log"
	"mime/multipart"
)

type UserServices struct {
	qry        users.UQuery
	pwd        utils.PassUtilInterface
	jwt        utils.JwtUtilityInterface
	cloudinary utils.CloudinaryUtilityInterface
}

func NewUserServices(q users.UQuery, p utils.PassUtilInterface, j utils.JwtUtilityInterface, c utils.CloudinaryUtilityInterface) users.UService {
	return &UserServices{
		qry:        q,
		pwd:        p,
		jwt:        j,
		cloudinary: c,
	}
}

func (us *UserServices) Login(email string, password string) (string, error) {
	result, err := us.qry.Login(email)

	if err != nil {
		log.Println("login query error: ", err.Error())
		return "", errors.New("login failed, please try again later")
	}

	err = us.pwd.ComparePassword([]byte(result.Password), []byte(password))
	if err != nil {
		log.Println("invalid password", err)
		return "", errors.New("invalid credentials")
	}

	token, err := us.jwt.GenerateJwt(result.ID)
	if err != nil {
		log.Println("error generating jwt", err)
		return "", errors.New("login failed, please try again later")
	}

	return token, nil
}

func (us *UserServices) Register(newUsers users.User, src multipart.File, filename string) error {
	hashPw, err := us.pwd.GeneratePassword(newUsers.Password)
	if err != nil {
		log.Println("register password generation error: ", err)
		return errors.New("registration failed, please try again later")
	}

	newUsers.Password = string(hashPw)
	newUsers.IsAdmin = false

	imageURL, err := us.cloudinary.UploadToCloudinary(src, filename)
	if err != nil {
		log.Println("image upload failed: ", err)
		return errors.New("failed to upload image, please try again later")
	}
	newUsers.UserImage = imageURL

	err = us.qry.Register(newUsers)
	if err != nil {
		log.Println("register query error: ", err)
		return errors.New("registration failed, please try again later")
	}

	return nil
}

func (us *UserServices) UpdateUser(userID uint, updatedUser users.User, src multipart.File, filename string) error {
	if updatedUser.Password != "" {
		hashPassword, err := us.pwd.GeneratePassword(updatedUser.Password)
		if err != nil {
			log.Println("update password generation error: ", err)
			return errors.New("update failed, please try again later")
		}

		updatedUser.Password = string(hashPassword)
	}

	if src != nil && filename != "" {
		imageURL, err := us.cloudinary.UploadToCloudinary(src, filename)
		if err != nil {
			log.Println("image upload failed: ", err)
			return errors.New("failed to upload image, please try again later")
		}
		updatedUser.UserImage = imageURL
	}

	err := us.qry.UpdateUser(userID, updatedUser)
	if err != nil {
		log.Println("update user query error: ", err)
		return errors.New("update failed, please try again later")
	}
	return nil
}

func (us *UserServices) DeleteUser(userID uint) error {
	err := us.qry.DeleteUser(userID)
	if err != nil {
		log.Println("delete user query error: ", err)
		return errors.New("delete failed, please try again later")
	}

	return nil
}

func (us *UserServices) GetUserByID(userID uint) (users.User, error) {
	user, err := us.qry.GetUserByID(userID)
	if err != nil {
		log.Println("get user by ID query error: ", err)
		return users.User{}, errors.New("failed to retrieve user, please try again later")
	}

	return user, nil
}

func (us *UserServices) IsAdmin(userID uint) (bool, error) {
	return us.qry.IsAdmin(userID)
}

func (us *UserServices) IsEmailExist(email string) (bool, error) {
	exist, err := us.qry.IsEmailExist(email)

	return exist, err
}

func (us *UserServices) IsPhoneExist(phone string) (bool, error) {
	exist, err := us.qry.IsPhoneExist(phone)

	return exist, err
}
