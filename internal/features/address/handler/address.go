package handler

import (
	"evermos_pbi/internal/features/address"
	"evermos_pbi/internal/helpers"
	"evermos_pbi/internal/utils"
	"log"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type AddressHandler struct {
	srv address.AService
	tu  utils.JwtUtilityInterface
}

func NewAddressHandler(s address.AService, t utils.JwtUtilityInterface) address.AHandler {
	return &AddressHandler{
		srv: s,
		tu:  t,
	}
}

func (ah *AddressHandler) AddAddress() echo.HandlerFunc {
	return func(c echo.Context) error {
		userID := ah.tu.DecodToken(c.Get("user").(*jwt.Token))
		if userID == 0 {
			log.Println("unauthorized attempt to update user")
			return c.JSON(http.StatusBadRequest, helpers.ResponseFormat(http.StatusUnauthorized, "failed", "unauthorized", nil))
		}

		var input AddOrUpdateAddressRequest

		err := c.Bind(&input)
		if err != nil {
			log.Println("failed to bind add address request")
			return c.JSON(http.StatusBadRequest, helpers.ResponseFormat(http.StatusBadRequest, "failed", "invalid input", nil))
		}

		newAddress := AddToAddress(input)
		newAddress.UserID = uint(userID)
		err = ah.srv.AddAddress(newAddress)
		if err != nil {
			log.Println("failed to add address")
			return c.JSON(http.StatusInternalServerError, helpers.ResponseFormat(http.StatusInternalServerError, "failed", "failed to add address", nil))
		}

		return c.JSON(http.StatusOK, helpers.ResponseFormat(http.StatusOK, "success", "success create address", nil))

	}
}

func (ah *AddressHandler) UpdateAddress() echo.HandlerFunc {
	return func(c echo.Context) error {
		userID := ah.tu.DecodToken(c.Get("user").(*jwt.Token))
		if userID == 0 {
			log.Println("unauthorized attempt to update user")
			return c.JSON(http.StatusBadRequest, helpers.ResponseFormat(http.StatusUnauthorized, "failed", "unauthorized", nil))
		}

		addressIDStr := c.Param("id")
		addressID, err := strconv.ParseUint(addressIDStr, 10, 32)
		if err != nil {
			log.Println("invalid address id")
			return c.JSON(http.StatusBadRequest, helpers.ResponseFormat(http.StatusBadRequest, "failed", "invalid store ID", nil))
		}

		var input AddOrUpdateAddressRequest

		err = c.Bind(&input)
		if err != nil {
			log.Println("failed to bind add address request")
			return c.JSON(http.StatusBadRequest, helpers.ResponseFormat(http.StatusBadRequest, "failed", "invalid input", nil))
		}

		newAddress := AddToAddress(input)
		newAddress.UserID = uint(userID)
		err = ah.srv.UpdateAddress(uint(userID), uint(addressID), newAddress)
		if err != nil {
			log.Println("failed to update address")
			return c.JSON(http.StatusInternalServerError, helpers.ResponseFormat(http.StatusInternalServerError, "failed", "failed to update address", nil))
		}

		return c.JSON(http.StatusOK, helpers.ResponseFormat(http.StatusOK, "success", "success update address", nil))

	}
}

func (ah *AddressHandler) DeleteAddress() echo.HandlerFunc {
	return func(c echo.Context) error {
		userID := ah.tu.DecodToken(c.Get("user").(*jwt.Token))
		if userID == 0 {
			log.Println("unauthorized attempt to update user")
			return c.JSON(http.StatusBadRequest, helpers.ResponseFormat(http.StatusUnauthorized, "failed", "unauthorized", nil))
		}

		addressIDStr := c.Param("id")
		addressID, err := strconv.ParseUint(addressIDStr, 10, 32)
		if err != nil {
			log.Println("invalid address id")
			return c.JSON(http.StatusBadRequest, helpers.ResponseFormat(http.StatusBadRequest, "failed", "invalid store ID", nil))
		}

		err = ah.srv.DeleteAddress(uint(userID), uint(addressID))
		if err != nil {
			log.Println("failed to delete address")
			return c.JSON(http.StatusInternalServerError, helpers.ResponseFormat(http.StatusInternalServerError, "failed", "failed to delete address", nil))
		}

		return c.JSON(http.StatusOK, helpers.ResponseFormat(http.StatusOK, "success", "success delete address", nil))
	}
}

func (ah *AddressHandler) GetAddressByUserID() echo.HandlerFunc {
	return func(c echo.Context) error {
		userID := ah.tu.DecodToken(c.Get("user").(*jwt.Token))
		if userID == 0 {
			log.Println("unauthorized attempt to update user")
			return c.JSON(http.StatusBadRequest, helpers.ResponseFormat(http.StatusUnauthorized, "failed", "unauthorized", nil))
		}

		address, err := ah.srv.GetAddressByUserID(uint(userID))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helpers.ResponseFormat(http.StatusInternalServerError, "failed", "failed to retrieve address", nil))
		}

		responseData := ToAddressResponses(address)
		return c.JSON(http.StatusOK, helpers.ResponseFormat(http.StatusOK, "success", "address was successfully retrieved", responseData))
	}
}

func (ah *AddressHandler) GetProvince() echo.HandlerFunc {
	return func(c echo.Context) error {
		provinces, err := ah.srv.GetProvince()
		if err != nil {
			log.Println("Error fetching provinces:", err)
			return c.JSON(http.StatusInternalServerError, helpers.ResponseFormat(http.StatusInternalServerError, "failed", "failed to retrieve provinces", nil))
		}

		return c.JSON(http.StatusOK, helpers.ResponseFormat(http.StatusOK, "success", "provinces successfully retrieved", provinces))
	}
}

func (ah *AddressHandler) GetRegency() echo.HandlerFunc {
	return func(c echo.Context) error {
		provinceIDStr := c.Param("province_id")
		provinceID, err := strconv.ParseUint(provinceIDStr, 10, 32)
		if err != nil {
			log.Println("invalid province id")
			return c.JSON(http.StatusBadRequest, helpers.ResponseFormat(http.StatusBadRequest, "failed", "invalid province ID", nil))
		}

		regencies, err := ah.srv.GetRegency(uint(provinceID))
		if err != nil {
			log.Println("Error fetching regencies:", err)
			return c.JSON(http.StatusInternalServerError, helpers.ResponseFormat(http.StatusInternalServerError, "failed", "failed to retrieve regencies", nil))
		}

		return c.JSON(http.StatusOK, helpers.ResponseFormat(http.StatusOK, "success", "regencies successfully retrieved", regencies))
	}
}

func (ah *AddressHandler) GetDistrict() echo.HandlerFunc {
	return func(c echo.Context) error {
		regencyIDStr := c.Param("regency_id")
		regencyID, err := strconv.ParseUint(regencyIDStr, 10, 32)
		if err != nil {
			log.Println("invalid regency id")
			return c.JSON(http.StatusBadRequest, helpers.ResponseFormat(http.StatusBadRequest, "failed", "invalid regency ID", nil))
		}

		districts, err := ah.srv.GetDistrict(uint(regencyID))
		if err != nil {
			log.Println("Error fetching regencies:", err)
			return c.JSON(http.StatusInternalServerError, helpers.ResponseFormat(http.StatusInternalServerError, "failed", "failed to retrieve districts", nil))
		}

		return c.JSON(http.StatusOK, helpers.ResponseFormat(http.StatusOK, "success", "regencies successfully retrieved", districts))
	}
}

func (ah *AddressHandler) GetVillage() echo.HandlerFunc {
	return func(c echo.Context) error {
		districtIDStr := c.Param("district_id")
		districtID, err := strconv.ParseUint(districtIDStr, 10, 32)
		if err != nil {
			log.Println("invalid district id")
			return c.JSON(http.StatusBadRequest, helpers.ResponseFormat(http.StatusBadRequest, "failed", "invalid district ID", nil))
		}

		villages, err := ah.srv.GetVillage(uint(districtID))
		if err != nil {
			log.Println("Error fetching regencies:", err)
			return c.JSON(http.StatusInternalServerError, helpers.ResponseFormat(http.StatusInternalServerError, "failed", "failed to retrieve villages", nil))
		}

		return c.JSON(http.StatusOK, helpers.ResponseFormat(http.StatusOK, "success", "villages successfully retrieved", villages))
	}
}
