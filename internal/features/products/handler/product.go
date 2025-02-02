package handler

import (
	"evermos_pbi/internal/features/products"
	"evermos_pbi/internal/helpers"
	"evermos_pbi/internal/utils"
	"log"
	"mime/multipart"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type ProductHandler struct {
	srv products.PService
	tu  utils.JwtUtilityInterface
}

func NewProductHandler(s products.PService, t utils.JwtUtilityInterface) products.PHandler {
	return &ProductHandler{
		srv: s,
		tu:  t,
	}
}

func (ph *ProductHandler) AddProduct() echo.HandlerFunc {
	return func(c echo.Context) error {
		userID := ph.tu.DecodToken(c.Get("user").(*jwt.Token))
		if userID == 0 {
			log.Println("unauthorized attempt to update user")
			return c.JSON(http.StatusBadRequest, helpers.ResponseFormat(http.StatusUnauthorized, "failed", "unauthorized", nil))
		}

		var src multipart.File
		var filename string
		image, err := c.FormFile("product_image")
		if err == nil {
			src, err = image.Open()
			if err != nil {
				log.Println("failed to open image file")
				return c.JSON(http.StatusInternalServerError, helpers.ResponseFormat(http.StatusInternalServerError, "failed", "unable to process image", nil))
			}
			defer src.Close()
			filename = image.Filename
		}

		var input AddOrUpdateProductRequest

		err = c.Bind(&input)
		if err != nil {
			log.Println("failed to bind add product request")
			return c.JSON(http.StatusBadRequest, helpers.ResponseFormat(http.StatusBadRequest, "failed", "invalid input", nil))
		}

		newProduct := AddToProduct(input)
		err = ph.srv.AddProduct(uint(userID), newProduct, src, filename)
		if err != nil {
			log.Println("failed to add a product")
			return c.JSON(http.StatusInternalServerError, helpers.ResponseFormat(http.StatusInternalServerError, "failed", "failed to add a product", nil))
		}

		return c.JSON(http.StatusOK, helpers.ResponseFormat(http.StatusOK, "success", "success create a product", nil))
	}
}

func (ph *ProductHandler) UpdateProduct() echo.HandlerFunc {
	return func(c echo.Context) error {
		userID := ph.tu.DecodToken(c.Get("user").(*jwt.Token))
		if userID == 0 {
			log.Println("unauthorized attempt to update user")
			return c.JSON(http.StatusBadRequest, helpers.ResponseFormat(http.StatusUnauthorized, "failed", "unauthorized", nil))
		}

		productIDStr := c.Param("id")
		productID, err := strconv.ParseUint(productIDStr, 10, 32)
		if err != nil {
			log.Println("invalid product id")
			return c.JSON(http.StatusBadRequest, helpers.ResponseFormat(http.StatusBadRequest, "failed", "invalid product ID", nil))
		}

		var src multipart.File
		var filename string
		image, err := c.FormFile("product_image")
		if err == nil {
			src, err = image.Open()
			if err != nil {
				log.Println("failed to open image file")
				return c.JSON(http.StatusInternalServerError, helpers.ResponseFormat(http.StatusInternalServerError, "failed", "unable to process image", nil))
			}
			defer src.Close()
			filename = image.Filename
		}

		var input AddOrUpdateProductRequest

		err = c.Bind(&input)
		if err != nil {
			log.Println("failed to bind update product request")
			return c.JSON(http.StatusBadRequest, helpers.ResponseFormat(http.StatusBadRequest, "failed", "invalid input", nil))
		}

		updateProduct := AddToProduct(input)
		err = ph.srv.UpdateProduct(uint(userID), uint(productID), updateProduct, src, filename)
		if err != nil {
			log.Println("failed to update product")
			return c.JSON(http.StatusInternalServerError, helpers.ResponseFormat(http.StatusInternalServerError, "failed", "failed to update product", nil))
		}

		return c.JSON(http.StatusOK, helpers.ResponseFormat(http.StatusOK, "success", "success update product", nil))

	}
}

func (ph *ProductHandler) DeleteProduct() echo.HandlerFunc {
	return func(c echo.Context) error {
		userID := ph.tu.DecodToken(c.Get("user").(*jwt.Token))
		if userID == 0 {
			log.Println("unauthorized attempt to update user")
			return c.JSON(http.StatusBadRequest, helpers.ResponseFormat(http.StatusUnauthorized, "failed", "unauthorized", nil))
		}

		productIDStr := c.Param("id")
		productID, err := strconv.ParseUint(productIDStr, 10, 32)
		if err != nil {
			log.Println("invalid product id")
			return c.JSON(http.StatusBadRequest, helpers.ResponseFormat(http.StatusBadRequest, "failed", "invalid product ID", nil))
		}

		err = ph.srv.DeleteProduct(uint(userID), uint(productID))
		if err != nil {
			log.Println("failed to delete product")
			return c.JSON(http.StatusInternalServerError, helpers.ResponseFormat(http.StatusInternalServerError, "failed", "failed to delete product", nil))
		}

		return c.JSON(http.StatusOK, helpers.ResponseFormat(http.StatusOK, "success", "success delete product", nil))

	}
}

func (ph *ProductHandler) GetProductByID() echo.HandlerFunc {
	return func(c echo.Context) error {
		productIDStr := c.Param("id")
		productID, err := strconv.ParseUint(productIDStr, 10, 32)
		if err != nil {
			log.Println("invalid product id")
			return c.JSON(http.StatusBadRequest, helpers.ResponseFormat(http.StatusBadRequest, "failed", "invalid product ID", nil))
		}

		product, err := ph.srv.GetProductByID(uint(productID))
		if err != nil {
			log.Println("failed to get product by ID")
			return c.JSON(http.StatusInternalServerError, helpers.ResponseFormat(http.StatusInternalServerError, "failed", "failed to retrieve product", nil))
		}

		ProductResponse := ToProductResponse(product)

		return c.JSON(http.StatusOK, helpers.ResponseFormat(http.StatusOK, "success", "product was successfully retrieved", ProductResponse))
	}
}

func (ph *ProductHandler) GetProductsByStoreID() echo.HandlerFunc {
	return func(c echo.Context) error {
		storeIDStr := c.Param("store_id")
		storeID, err := strconv.ParseUint(storeIDStr, 10, 32)
		if err != nil {
			log.Println("invalid store id")
			return c.JSON(http.StatusBadRequest, helpers.ResponseFormat(http.StatusBadRequest, "failed", "invalid store ID", nil))
		}

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

		products, totalItems, err := ph.srv.GetProductsByStoreID(uint(storeID), uint(limit), uint(page), search)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helpers.ResponseFormat(http.StatusInternalServerError, "failed", "failed to retrieve products", nil))
		}

		responseData := ToProductResponses(products)
		meta := helpers.Meta{
			TotalItems:   int(totalItems),
			ItemsPerPage: limit,
			CurrentPage:  page,
			TotalPages:   (int(totalItems) + limit - 1) / limit,
		}

		return c.JSON(http.StatusOK, helpers.ResponseWithMetaFormat(http.StatusOK, "success", "successfully retrieved all products", responseData, meta))

	}
}

func (ph *ProductHandler) GetAllProducts() echo.HandlerFunc {
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

		products, totalItems, err := ph.srv.GetAllProducts(uint(limit), uint(page), search)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helpers.ResponseFormat(http.StatusInternalServerError, "failed", "failed to retrieve products", nil))
		}

		responseData := ToProductResponses(products)
		meta := helpers.Meta{
			TotalItems:   int(totalItems),
			ItemsPerPage: limit,
			CurrentPage:  page,
			TotalPages:   (int(totalItems) + limit - 1) / limit,
		}

		return c.JSON(http.StatusOK, helpers.ResponseWithMetaFormat(http.StatusOK, "success", "successfully retrieved all products", responseData, meta))

	}
}
