package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	dtos "github.com/savanyv/Golang-Findest/internal/dto"
	"github.com/savanyv/Golang-Findest/internal/helpers"
	"github.com/savanyv/Golang-Findest/internal/usecase"
)

type AuthHandler struct {
	usecase usecase.UserUsecase
	validator *helpers.CustomValidator
}

func NewAuthHandler(usecase usecase.UserUsecase) *AuthHandler {
	return &AuthHandler{
		usecase: usecase,
		validator: helpers.NewValidator(),
	}
}

func (h *AuthHandler) Register(c echo.Context) error {
	var req dtos.RegisterRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "failed to bind request",
		})
	}

	if err := h.validator.Validate(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "failed to validate request",
		})
	}

	response, err := h.usecase.Register(&req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, echo.Map{
		"message": "successfully registered",
		"data": response,
	})
}

func (h *AuthHandler) Login(c echo.Context) error {
	var req dtos.LoginRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "failed to bind request",
		})
	}

	if err := h.validator.Validate(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "failed to validate request",
		})
	}

	response, err := h.usecase.Login(&req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "successfully logged in",
		"data": response,
	})
}