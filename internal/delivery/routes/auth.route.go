package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/savanyv/Golang-Findest/internal/config/database"
	"github.com/savanyv/Golang-Findest/internal/delivery/handlers"
	"github.com/savanyv/Golang-Findest/internal/repository"
	"github.com/savanyv/Golang-Findest/internal/usecase"
)

func authRoutes(e *echo.Group) {
	repo := repository.NewUserRepository(database.DB)
	usecase := usecase.NewUserUsecase(repo)
	handler := handlers.NewAuthHandler(usecase)

	e.POST("/register", handler.Register)
	e.POST("/login", handler.Login)
}