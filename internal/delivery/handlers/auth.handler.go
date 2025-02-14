package handlers

import (
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
		return c.JSON(400, map[string]string{
			"error": err.Error(),
		})
	}

	if err := h.validator.Validate(&req); err != nil {
		return c.JSON(400, map[string]string{
			"error": err.Error(),
		})
	}

	response, err := h.usecase.Register(&req)
	if err != nil {
		return c.JSON(400, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(200, map[string]interface{}{
		"message": "success",
		"data": response,
	})
}

func (h *AuthHandler) Login(c echo.Context) error {
	var req dtos.LoginRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(400, map[string]string{
			"error": err.Error(),
		})
	}

	if err := h.validator.Validate(&req); err != nil {
		return c.JSON(400, map[string]string{
			"error": err.Error(),
		})
	}

	response, err := h.usecase.Login(&req)
	if err != nil {
		return c.JSON(400, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(200, map[string]interface{}{
		"message": "success",
		"data": response,
	})
}