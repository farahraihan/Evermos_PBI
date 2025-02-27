package handler

import (
	"evermos_pbi/internal/features/transaction"
	"evermos_pbi/internal/helpers"
	"evermos_pbi/internal/utils"
	"log"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type TransactionHandler struct {
	srv transaction.TService
	tu  utils.JwtUtilityInterface
}

func NewTransactionHandler(s transaction.TService, t utils.JwtUtilityInterface) transaction.THandler {
	return &TransactionHandler{
		srv: s,
		tu:  t,
	}
}

func (th *TransactionHandler) AddTransaction() echo.HandlerFunc {
	return func(c echo.Context) error {
		userID := th.tu.DecodToken(c.Get("user").(*jwt.Token))
		if userID == 0 {
			log.Println("unauthorized attempt to update user")
			return c.JSON(http.StatusBadRequest, helpers.ResponseFormat(http.StatusUnauthorized, "failed", "unauthorized", nil))
		}

		var input AddTransactionRequest

		err := c.Bind(&input)
		if err != nil {
			log.Println("failed to bind add transaction request", err)
			return c.JSON(http.StatusBadRequest, helpers.ResponseFormat(http.StatusBadRequest, "failed", "failed add product to cart", nil))
		}

		input.UserID = uint(userID)

		newTransaction := AddToTransaction(input)
		newDetailTransaction := AddToDetailTransaction(input)
		err = th.srv.AddTransaction(newTransaction, newDetailTransaction)
		if err != nil {
			if err.Error() == "insufficient stock: the requested quantity exceeds available stock" {
				log.Println("failed add product to cart", err)
				return c.JSON(http.StatusInternalServerError, helpers.ResponseFormat(http.StatusInternalServerError, "failed", "insufficient stock", nil))
			}
			log.Println("failed add product to cart", err)
			return c.JSON(http.StatusInternalServerError, helpers.ResponseFormat(http.StatusInternalServerError, "failed", "failed add product to cart", nil))
		}

		return c.JSON(http.StatusOK, helpers.ResponseFormat(http.StatusOK, "success", "success add product to cart", nil))

	}
}

func (th *TransactionHandler) UpdateDetailTransaction() echo.HandlerFunc {
	return func(c echo.Context) error {
		userID := th.tu.DecodToken(c.Get("user").(*jwt.Token))
		if userID == 0 {
			log.Println("unauthorized attempt to update user")
			return c.JSON(http.StatusBadRequest, helpers.ResponseFormat(http.StatusUnauthorized, "failed", "unauthorized", nil))
		}

		productIDStr := c.QueryParam("product_id")
		productID, err := strconv.ParseUint(productIDStr, 10, 32)
		if err != nil {
			log.Println("invalid product id", err)
			return c.JSON(http.StatusBadRequest, helpers.ResponseFormat(http.StatusBadRequest, "failed", "invalid product ID", nil))
		}

		transactionIDStr := c.QueryParam("transaction_id")
		transactionID, err := strconv.ParseUint(transactionIDStr, 10, 32)
		if err != nil {
			log.Println("invalid transaction id", err)
			return c.JSON(http.StatusBadRequest, helpers.ResponseFormat(http.StatusBadRequest, "failed", "invalid transaction ID", nil))
		}

		quantity := uint(1)

		err = th.srv.UpdateDetailTransaction(uint(userID), uint(transactionID), uint(productID), quantity)
		if err != nil {
			log.Println("failed add quantity", err)
			return c.JSON(http.StatusInternalServerError, helpers.ResponseFormat(http.StatusInternalServerError, "failed", "failed add quantity", nil))
		}

		return c.JSON(http.StatusOK, helpers.ResponseFormat(http.StatusOK, "success", "success add quantity", nil))

	}
}

func (th *TransactionHandler) DeleteTransaction() echo.HandlerFunc {
	return func(c echo.Context) error {
		userID := th.tu.DecodToken(c.Get("user").(*jwt.Token))
		if userID == 0 {
			log.Println("unauthorized attempt to update user")
			return c.JSON(http.StatusBadRequest, helpers.ResponseFormat(http.StatusUnauthorized, "failed", "unauthorized", nil))
		}

		productIDStr := c.QueryParam("product_id")
		productID, err := strconv.ParseUint(productIDStr, 10, 32)
		if err != nil {
			log.Println("invalid product id", err)
			return c.JSON(http.StatusBadRequest, helpers.ResponseFormat(http.StatusBadRequest, "failed", "invalid product ID", nil))
		}

		transactionIDStr := c.QueryParam("transaction_id")
		transactionID, err := strconv.ParseUint(transactionIDStr, 10, 32)
		if err != nil {
			log.Println("invalid transaction id", err)
			return c.JSON(http.StatusBadRequest, helpers.ResponseFormat(http.StatusBadRequest, "failed", "invalid transaction ID", nil))
		}

		err = th.srv.DeleteTransaction(uint(userID), uint(transactionID), uint(productID))
		if err != nil {
			log.Println("failed to delete detail transaction", err)
			return c.JSON(http.StatusInternalServerError, helpers.ResponseFormat(http.StatusInternalServerError, "failed", "failed to delete product from cart", nil))
		}

		return c.JSON(http.StatusOK, helpers.ResponseFormat(http.StatusOK, "success", "success delete product from cart", nil))

	}
}

func (th *TransactionHandler) UpdateTransactionStatusCanceled() echo.HandlerFunc {
	return func(c echo.Context) error {
		userID := th.tu.DecodToken(c.Get("user").(*jwt.Token))
		if userID == 0 {
			log.Println("unauthorized attempt to update user")
			return c.JSON(http.StatusBadRequest, helpers.ResponseFormat(http.StatusUnauthorized, "failed", "unauthorized", nil))
		}

		transactionIDStr := c.Param("transaction_id")
		transactionID, err := strconv.ParseUint(transactionIDStr, 10, 32)
		if err != nil {
			log.Println("invalid transaction id", err)
			return c.JSON(http.StatusBadRequest, helpers.ResponseFormat(http.StatusBadRequest, "failed", "invalid transaction ID", nil))
		}

		status := "canceled"

		err = th.srv.UpdateTransaction(uint(userID), uint(transactionID), status)
		if err != nil {
			log.Println("failed to delete transaction", err)
			return c.JSON(http.StatusInternalServerError, helpers.ResponseFormat(http.StatusInternalServerError, "failed", "failed to cancel transaction", nil))
		}

		return c.JSON(http.StatusOK, helpers.ResponseFormat(http.StatusOK, "success", "success cancel transaction (delete cart)", nil))

	}
}

func (th *TransactionHandler) UpdateTransactionStatusCompleted() echo.HandlerFunc {
	return func(c echo.Context) error {
		userID := th.tu.DecodToken(c.Get("user").(*jwt.Token))
		if userID == 0 {
			log.Println("unauthorized attempt to update user")
			return c.JSON(http.StatusBadRequest, helpers.ResponseFormat(http.StatusUnauthorized, "failed", "unauthorized", nil))
		}

		transactionIDStr := c.Param("transaction_id")
		transactionID, err := strconv.ParseUint(transactionIDStr, 10, 32)
		if err != nil {
			log.Println("invalid transaction id", err)
			return c.JSON(http.StatusBadRequest, helpers.ResponseFormat(http.StatusBadRequest, "failed", "invalid transaction ID", nil))
		}

		status := "completed"

		err = th.srv.UpdateTransaction(uint(userID), uint(transactionID), status)
		if err != nil {
			log.Println("failed to update transaction status", err)
			return c.JSON(http.StatusInternalServerError, helpers.ResponseFormat(http.StatusInternalServerError, "failed", "failed to checkout", nil))
		}

		return c.JSON(http.StatusOK, helpers.ResponseFormat(http.StatusOK, "success", "checkout successful", nil))

	}
}

func (th *TransactionHandler) GetTransactionByStatusCart() echo.HandlerFunc {
	return func(c echo.Context) error {
		userID := th.tu.DecodToken(c.Get("user").(*jwt.Token))
		if userID == 0 {
			log.Println("unauthorized attempt to update user")
			return c.JSON(http.StatusBadRequest, helpers.ResponseFormat(http.StatusUnauthorized, "failed", "unauthorized", nil))
		}

		transactionIDStr := c.Param("transaction_id")
		transactionID, err := strconv.ParseUint(transactionIDStr, 10, 32)
		if err != nil {
			log.Println("invalid transaction id", err)
			return c.JSON(http.StatusBadRequest, helpers.ResponseFormat(http.StatusBadRequest, "failed", "invalid transaction ID", nil))
		}

		cart, err := th.srv.GetTransactionByStatusCart(uint(userID), uint(transactionID))
		if err != nil {
			log.Println("failed to retrieve transaction", err)
			return c.JSON(http.StatusInternalServerError, helpers.ResponseFormat(http.StatusInternalServerError, "failed", "failed to retrieve cart", nil))
		}

		CartResponse := ToTransactionResponse(cart)

		return c.JSON(http.StatusOK, helpers.ResponseFormat(http.StatusOK, "success", "cart was successfully retrieved", CartResponse))

	}
}

func (th *TransactionHandler) GetTransactionHistory() echo.HandlerFunc {
	return func(c echo.Context) error {
		userID := th.tu.DecodToken(c.Get("user").(*jwt.Token))
		if userID == 0 {
			log.Println("unauthorized attempt to update user")
			return c.JSON(http.StatusBadRequest, helpers.ResponseFormat(http.StatusUnauthorized, "failed", "unauthorized", nil))
		}

		limit, err := strconv.Atoi(c.QueryParam("limit"))
		if err != nil || limit <= 0 {
			limit = 10
		}

		page, err := strconv.Atoi(c.QueryParam("page"))
		if err != nil || page < 0 {
			page = 1
		}

		transactionHistory, totalItems, err := th.srv.GetTransactionHistory(uint(userID), uint(limit), uint(page))
		if err != nil {
			log.Println("failed to retrieve transaction history", err)
			return c.JSON(http.StatusInternalServerError, helpers.ResponseFormat(http.StatusInternalServerError, "failed", "failed to retrieve transaction history", nil))
		}

		responseData := ToTransactionHistoryResponses(transactionHistory)
		meta := helpers.Meta{
			TotalItems:   int(totalItems),
			ItemsPerPage: limit,
			CurrentPage:  page,
			TotalPages:   (int(totalItems) + limit - 1) / limit,
		}

		return c.JSON(http.StatusOK, helpers.ResponseWithMetaFormat(http.StatusOK, "success", "successfully retrieved transaction history", responseData, meta))

	}
}

func (th *TransactionHandler) GetTransactionByID() echo.HandlerFunc {
	return func(c echo.Context) error {
		userID := th.tu.DecodToken(c.Get("user").(*jwt.Token))
		if userID == 0 {
			log.Println("unauthorized attempt to update user")
			return c.JSON(http.StatusBadRequest, helpers.ResponseFormat(http.StatusUnauthorized, "failed", "unauthorized", nil))
		}

		transactionIDStr := c.Param("transaction_id")
		transactionID, err := strconv.ParseUint(transactionIDStr, 10, 32)
		if err != nil {
			log.Println("invalid transaction id", err)
			return c.JSON(http.StatusBadRequest, helpers.ResponseFormat(http.StatusBadRequest, "failed", "invalid transaction ID", nil))
		}

		transaction, err := th.srv.GetTransactionByID(uint(userID), uint(transactionID))
		if err != nil {
			log.Println("failed to retrieve transaction by ID", err)
			return c.JSON(http.StatusInternalServerError, helpers.ResponseFormat(http.StatusInternalServerError, "failed", "failed to retrieve transaction by ID", nil))
		}

		TransactionResponse := ToTransactionResponse(transaction)
		return c.JSON(http.StatusOK, helpers.ResponseFormat(http.StatusOK, "success", "cart was successfully retrieved", TransactionResponse))

	}
}
