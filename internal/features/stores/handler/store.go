package handler

import (
	"evermos_pbi/internal/features/stores"
	"evermos_pbi/internal/helpers"
	"evermos_pbi/internal/utils"
	"log"
	"mime/multipart"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type StoreHandler struct {
	srv stores.SService
	tu  utils.JwtUtilityInterface
}

func NewStoreHandler(s stores.SService, t utils.JwtUtilityInterface) stores.SHandler {
	return &StoreHandler{
		srv: s,
		tu:  t,
	}
}

func (sh *StoreHandler) AddStore() echo.HandlerFunc {
	return func(c echo.Context) error {
		userID := sh.tu.DecodToken(c.Get("user").(*jwt.Token))
		if userID == 0 {
			log.Println("unauthorized attempt to update user")
			return c.JSON(http.StatusBadRequest, helpers.ResponseFormat(http.StatusUnauthorized, "failed", "unauthorized", nil))
		}

		var src multipart.File
		var filename string
		image, err := c.FormFile("store_image")
		if err == nil {
			src, err = image.Open()
			if err != nil {
				log.Println("failed to open image file")
				return c.JSON(http.StatusInternalServerError, helpers.ResponseFormat(http.StatusInternalServerError, "failed", "unable to process image", nil))
			}
			defer src.Close()
			filename = image.Filename
		}

		var input AddOrUpdateStoreRequest

		err = c.Bind(&input)
		if err != nil {
			log.Println("failed to bind add store request")
			return c.JSON(http.StatusBadRequest, helpers.ResponseFormat(http.StatusBadRequest, "failed", "invalid input", nil))
		}

		newStore := AddToStore(input)
		newStore.UserID = uint(userID)
		err = sh.srv.AddStore(newStore, src, filename)
		if err != nil {
			log.Println("failed to add a store")
			return c.JSON(http.StatusInternalServerError, helpers.ResponseFormat(http.StatusInternalServerError, "failed", "failed to add a store", nil))
		}

		return c.JSON(http.StatusOK, helpers.ResponseFormat(http.StatusOK, "success", "success create a store", nil))
	}
}

func (sh *StoreHandler) UpdateStore() echo.HandlerFunc {
	return func(c echo.Context) error {
		userID := sh.tu.DecodToken(c.Get("user").(*jwt.Token))
		if userID == 0 {
			log.Println("unauthorized attempt to update user")
			return c.JSON(http.StatusBadRequest, helpers.ResponseFormat(http.StatusUnauthorized, "failed", "unauthorized", nil))
		}

		var src multipart.File
		var filename string
		image, err := c.FormFile("store_image")
		if err == nil {
			src, err = image.Open()
			if err != nil {
				log.Println("failed to open image file")
				return c.JSON(http.StatusInternalServerError, helpers.ResponseFormat(http.StatusInternalServerError, "failed", "unable to process image", nil))
			}
			defer src.Close()
			filename = image.Filename
		}

		var input AddOrUpdateStoreRequest

		err = c.Bind(&input)
		if err != nil {
			log.Println("failed to bind add store request")
			return c.JSON(http.StatusBadRequest, helpers.ResponseFormat(http.StatusBadRequest, "failed", "invalid input", nil))
		}

		updateStore := AddToStore(input)
		err = sh.srv.UpdateStore(uint(userID), updateStore, src, filename)
		if err != nil {
			log.Println("failed to update a store")
			return c.JSON(http.StatusInternalServerError, helpers.ResponseFormat(http.StatusInternalServerError, "failed", "failed to update a store", nil))
		}

		return c.JSON(http.StatusOK, helpers.ResponseFormat(http.StatusOK, "success", "success update store", nil))
	}
}

func (sh *StoreHandler) DeleteStore() echo.HandlerFunc {
	return func(c echo.Context) error {
		userID := sh.tu.DecodToken(c.Get("user").(*jwt.Token))
		if userID == 0 {
			log.Println("unauthorized attempt to update user")
			return c.JSON(http.StatusBadRequest, helpers.ResponseFormat(http.StatusUnauthorized, "failed", "unauthorized", nil))
		}

		err := sh.srv.DeleteStore(uint(userID))
		if err != nil {
			log.Println("failed to delete store")
			return c.JSON(http.StatusInternalServerError, helpers.ResponseFormat(http.StatusInternalServerError, "failed", "failed to delete store", nil))
		}

		return c.JSON(http.StatusOK, helpers.ResponseFormat(http.StatusOK, "success", "success delete store", nil))

	}
}

func (sh *StoreHandler) GetStoreByID() echo.HandlerFunc {
	return func(c echo.Context) error {
		storeIDStr := c.Param("id")
		storeID, err := strconv.ParseUint(storeIDStr, 10, 32)
		if err != nil {
			log.Println("invalid store id")
			return c.JSON(http.StatusBadRequest, helpers.ResponseFormat(http.StatusBadRequest, "failed", "invalid store ID", nil))
		}

		store, err := sh.srv.GetStoreByID(uint(storeID))
		if err != nil {
			log.Println("failed to get store by ID")
			return c.JSON(http.StatusInternalServerError, helpers.ResponseFormat(http.StatusInternalServerError, "failed", "failed to retrieve store", nil))
		}

		StoreResponse := ToStoreResponse(store)

		return c.JSON(http.StatusOK, helpers.ResponseFormat(http.StatusOK, "success", "store was successfully retrieved", StoreResponse))
	}
}

func (sh *StoreHandler) GetStoreByUserID() echo.HandlerFunc {
	return func(c echo.Context) error {
		userID := sh.tu.DecodToken(c.Get("user").(*jwt.Token))
		if userID == 0 {
			log.Println("unauthorized attempt to update user")
			return c.JSON(http.StatusBadRequest, helpers.ResponseFormat(http.StatusUnauthorized, "failed", "unauthorized", nil))
		}

		store, err := sh.srv.GetStoreByUserID(uint(userID))
		if err != nil {
			log.Println("failed to get store by user ID")
			return c.JSON(http.StatusInternalServerError, helpers.ResponseFormat(http.StatusInternalServerError, "failed", "failed to retrieve store", nil))
		}

		StoreResponse := ToStoreResponse(store)

		return c.JSON(http.StatusOK, helpers.ResponseFormat(http.StatusOK, "success", "store was successfully retrieved", StoreResponse))
	}
}

func (sh *StoreHandler) GetAllStores() echo.HandlerFunc {
	return func(c echo.Context) error {

		limit, err := strconv.Atoi(c.QueryParam("limit"))
		if err != nil || limit <= 0 {
			limit = 10
		}

		page, err := strconv.Atoi(c.QueryParam("page"))
		if err != nil || page < 0 {
			page = 1
		}

		search := c.QueryParam("search")
		if search == "" {
			search = ""
		}

		stores, totalItems, err := sh.srv.GetAllStores(uint(limit), uint(page), search)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helpers.ResponseFormat(http.StatusInternalServerError, "failed", "failed to retrieve books", nil))
		}

		responseData := ToStoreResponses(stores)
		meta := helpers.Meta{
			TotalItems:   int(totalItems),
			ItemsPerPage: limit,
			CurrentPage:  page,
			TotalPages:   (int(totalItems) + limit - 1) / limit,
		}

		return c.JSON(http.StatusOK, helpers.ResponseWithMetaFormat(http.StatusOK, "success", "successfully retrieved all stores", responseData, meta))

	}
}
