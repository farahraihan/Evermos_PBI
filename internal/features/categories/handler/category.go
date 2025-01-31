package handler

import (
	"evermos_pbi/internal/features/categories"
	"evermos_pbi/internal/helpers"
	"evermos_pbi/internal/utils"
	"log"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type CategoryHandler struct {
	srv categories.CService
	tu  utils.JwtUtilityInterface
}

func NewCategoryHandler(s categories.CService, t utils.JwtUtilityInterface) categories.CHandler {
	return &CategoryHandler{
		srv: s,
		tu:  t,
	}
}

func (ch *CategoryHandler) AddCategory() echo.HandlerFunc {
	return func(c echo.Context) error {
		userID := ch.tu.DecodToken(c.Get("user").(*jwt.Token))
		if userID == 0 {
			log.Println("unauthorized attempt to update user")
			return c.JSON(http.StatusBadRequest, helpers.ResponseFormat(http.StatusUnauthorized, "failed", "unauthorized", nil))
		}

		var input AddOrUpdateCategoryRequest

		err := c.Bind(&input)
		if err != nil {
			log.Println("failed to bind add category request")
			return c.JSON(http.StatusBadRequest, helpers.ResponseFormat(http.StatusBadRequest, "failed", "invalid input", nil))
		}

		newCategory := AddToCategory(input)
		err = ch.srv.AddCategory(uint(userID), newCategory)
		if err != nil {
			log.Println("failed to add category")
			return c.JSON(http.StatusInternalServerError, helpers.ResponseFormat(http.StatusInternalServerError, "failed", "failed to add category", nil))
		}

		return c.JSON(http.StatusOK, helpers.ResponseFormat(http.StatusOK, "success", "success create a category", nil))

	}
}

func (ch *CategoryHandler) UpdateCategory() echo.HandlerFunc {
	return func(c echo.Context) error {
		userID := ch.tu.DecodToken(c.Get("user").(*jwt.Token))
		if userID == 0 {
			log.Println("unauthorized attempt to update user")
			return c.JSON(http.StatusBadRequest, helpers.ResponseFormat(http.StatusUnauthorized, "failed", "unauthorized", nil))
		}

		categoryIDStr := c.Param("id")
		categoryID, err := strconv.ParseUint(categoryIDStr, 10, 32)
		if err != nil {
			log.Println("invalid category id")
			return c.JSON(http.StatusBadRequest, helpers.ResponseFormat(http.StatusBadRequest, "failed", "invalid category ID", nil))
		}

		var input AddOrUpdateCategoryRequest

		err = c.Bind(&input)
		if err != nil {
			log.Println("failed to bind update category request")
			return c.JSON(http.StatusBadRequest, helpers.ResponseFormat(http.StatusBadRequest, "failed", "invalid input", nil))
		}

		updateCategory := AddToCategory(input)
		err = ch.srv.UpdateCategory(uint(userID), uint(categoryID), updateCategory)
		if err != nil {
			log.Println("failed to update category")
			return c.JSON(http.StatusInternalServerError, helpers.ResponseFormat(http.StatusInternalServerError, "failed", "failed to update category", nil))
		}

		return c.JSON(http.StatusOK, helpers.ResponseFormat(http.StatusOK, "success", "success update category", nil))

	}
}

func (ch *CategoryHandler) DeleteCategory() echo.HandlerFunc {
	return func(c echo.Context) error {
		userID := ch.tu.DecodToken(c.Get("user").(*jwt.Token))
		if userID == 0 {
			log.Println("unauthorized attempt to update user")
			return c.JSON(http.StatusBadRequest, helpers.ResponseFormat(http.StatusUnauthorized, "failed", "unauthorized", nil))
		}

		categoryIDStr := c.Param("id")
		categoryID, err := strconv.ParseUint(categoryIDStr, 10, 32)
		if err != nil {
			log.Println("invalid category id")
			return c.JSON(http.StatusBadRequest, helpers.ResponseFormat(http.StatusBadRequest, "failed", "invalid category ID", nil))
		}

		err = ch.srv.DeleteCategory(uint(userID), uint(categoryID))
		if err != nil {
			log.Println("failed to delete category")
			return c.JSON(http.StatusInternalServerError, helpers.ResponseFormat(http.StatusInternalServerError, "failed", "failed to delete category", nil))
		}

		return c.JSON(http.StatusOK, helpers.ResponseFormat(http.StatusOK, "success", "success delete category", nil))
	}
}

func (ch *CategoryHandler) GetCategoryByID() echo.HandlerFunc {
	return func(c echo.Context) error {
		categoryIDStr := c.Param("id")
		categoryID, err := strconv.ParseUint(categoryIDStr, 10, 32)
		if err != nil {
			log.Println("invalid category id")
			return c.JSON(http.StatusBadRequest, helpers.ResponseFormat(http.StatusBadRequest, "failed", "invalid category ID", nil))
		}

		category, err := ch.srv.GetCategoryByID(uint(categoryID))
		if err != nil {
			log.Println("failed to get category by ID")
			return c.JSON(http.StatusInternalServerError, helpers.ResponseFormat(http.StatusInternalServerError, "failed", "failed to retrieve category", nil))
		}

		CategoryResponse := ToCategoryResponse(category)
		return c.JSON(http.StatusOK, helpers.ResponseFormat(http.StatusOK, "success", "category was successfully retrieved", CategoryResponse))
	}
}

func (ch *CategoryHandler) GetAllCategories() echo.HandlerFunc {
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

		categories, totalItems, err := ch.srv.GetAllCategories(uint(limit), uint(page), search)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helpers.ResponseFormat(http.StatusInternalServerError, "failed", "failed to retrieve categories", nil))
		}

		responseData := ToCategoryResponses(categories)
		meta := helpers.Meta{
			TotalItems:   int(totalItems),
			ItemsPerPage: limit,
			CurrentPage:  page,
			TotalPages:   (int(totalItems) + limit - 1) / limit,
		}

		return c.JSON(http.StatusOK, helpers.ResponseWithMetaFormat(http.StatusOK, "success", "successfully retrieved all categories", responseData, meta))

	}
}
