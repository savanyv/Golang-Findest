package handlers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	dtos "github.com/savanyv/Golang-Findest/internal/dto"
	"github.com/savanyv/Golang-Findest/internal/helpers"
	"github.com/savanyv/Golang-Findest/internal/usecase"
)

type TransactionHandler struct {
	usecase usecase.TransactionUsecase
	validator *helpers.CustomValidator
}

func NewTransactionHandler(usecase usecase.TransactionUsecase) *TransactionHandler {
	return &TransactionHandler{
		usecase: usecase,
		validator: helpers.NewValidator(),
	}
}

func (h *TransactionHandler) CreateTransaction(c echo.Context) error {
	userIDStr, ok := c.Get("userID").(string)
	if !ok {
		return c.JSON(http.StatusUnauthorized, echo.Map{
			"error": "failed to get user id",
		})
	}
	userID, err := strconv.ParseUint(userIDStr, 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "failed to parse user id",
		})
	}

	var req dtos.CreateTransactionRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": err.Error(),
		})
	}

	if err := h.validator.Validate(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "failed to validate request",
		})
	}

	response, err := h.usecase.CreateTransaction(&req, uint(userID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, echo.Map{
		"message": "successfully created transaction",
		"data":    response,
	})
}

func (h *TransactionHandler) GetTransaction(c echo.Context) error {
	var userID *uint
	var status *string

	userIDStr := c.QueryParam("user_id")
	if userIDStr != "" {
		parsedUserID, err := strconv.ParseUint(userIDStr, 10, 32)
		if err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{
				"error": "failed to parse user id",
			})
		}
		convertedUserID := uint(parsedUserID)
		userID = &convertedUserID
	}

	statusStr := c.QueryParam("status")
	if statusStr != "" {
		status = &statusStr
	}

	response, err := h.usecase.GetTransaction(userID, status)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "successfully get transactions",
		"data":    response,
	})
}

func (h *TransactionHandler) GetTransactionByID(c echo.Context) error {
	userIDStr, ok := c.Get("userID").(string)
	if !ok {
		return c.JSON(http.StatusUnauthorized, echo.Map{
			"error": "failed to get user id",
		})
	}

	userID, err := strconv.ParseUint(userIDStr, 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "failed to parse user id",
		})
	}

	transactionID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "failed to parse transaction id",
		})
	}

	response, err := h.usecase.GetTransactionByID(uint(userID), uint(transactionID))
	if err != nil {
		return c.JSON(http.StatusForbidden, echo.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "successfully get transaction",
		"data":    response,
	})
}

func (h *TransactionHandler) UpdateStatusTransaction(c echo.Context) error {
	transactionID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "failed to parse transaction id",
		})
	}

	var req dtos.UpdateTranscationRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": err.Error(),
		})
	}

	if err := h.validator.Validate(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": err.Error(),
		})
	}

	response, err := h.usecase.UpdateStatusTransaction(uint(transactionID), &req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "successfully update status transaction",
		"data": response,
	})
}