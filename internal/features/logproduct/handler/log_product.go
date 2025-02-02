package handler

import (
	"evermos_pbi/internal/features/logproduct"
	"evermos_pbi/internal/helpers"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type LogProductHandler struct {
	srv logproduct.LService
}

func NewLogProductHandler(s logproduct.LService) logproduct.LHandler {
	return &LogProductHandler{
		srv: s,
	}
}

func (lh *LogProductHandler) GetLogProductByID() echo.HandlerFunc {
	return func(c echo.Context) error {
		logProductIDStr := c.Param("id")
		logProductID, err := strconv.ParseUint(logProductIDStr, 10, 32)
		if err != nil {
			log.Println("invalid log product id")
			return c.JSON(http.StatusBadRequest, helpers.ResponseFormat(http.StatusBadRequest, "failed", "invalid log product ID", nil))
		}

		logProduct, err := lh.srv.GetLogProductByID(uint(logProductID))
		if err != nil {
			log.Println("failed to get log product by ID")
			return c.JSON(http.StatusInternalServerError, helpers.ResponseFormat(http.StatusInternalServerError, "failed", "failed to retrieve log product", nil))
		}

		LogProductResponse := ToLogProductResponse(logProduct)
		return c.JSON(http.StatusOK, helpers.ResponseFormat(http.StatusOK, "success", "log product was successfully retrieved", LogProductResponse))

	}
}

func (lh *LogProductHandler) GetAllLogProduct() echo.HandlerFunc {
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

		logProduct, totalItems, err := lh.srv.GetAllLogProduct(uint(limit), uint(page), search)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helpers.ResponseFormat(http.StatusInternalServerError, "failed", "failed to retrieve log product", nil))
		}

		responseData := ToLogProductResponses(logProduct)
		meta := helpers.Meta{
			TotalItems:   int(totalItems),
			ItemsPerPage: limit,
			CurrentPage:  page,
			TotalPages:   (int(totalItems) + limit - 1) / limit,
		}

		return c.JSON(http.StatusOK, helpers.ResponseWithMetaFormat(http.StatusOK, "success", "successfully retrieved all log product", responseData, meta))
	}
}
