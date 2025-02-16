package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/savanyv/Golang-Findest/internal/config/database"
	"github.com/savanyv/Golang-Findest/internal/delivery/handlers"
	"github.com/savanyv/Golang-Findest/internal/helpers"
	"github.com/savanyv/Golang-Findest/internal/middlewares"
	"github.com/savanyv/Golang-Findest/internal/repository"
	"github.com/savanyv/Golang-Findest/internal/usecase"
)

func transactionRoutes(e *echo.Group) {
	jwtService := helpers.NewJWTService()

	repo := repository.NewTransactionRepository(database.DB)
	userRepo := repository.NewUserRepository(database.DB)
	usecase := usecase.NewTransactionUsecase(repo, userRepo)
	handler := handlers.NewTransactionHandler(usecase)

	e.POST("/transactions", handler.CreateTransaction, middlewares.AuthMiddlewares(jwtService))
	e.GET("/transactions", handler.GetTransaction)
	e.GET("/transactions/:id", handler.GetTransactionByID, middlewares.AuthMiddlewares(jwtService))
	e.PUT("/transactions/:id", handler.UpdateStatusTransaction)
	e.DELETE("/transactions/:id", handler.DeleteTransaction, middlewares.AuthMiddlewares(jwtService))

	e.GET("/dashboard/summary", handler.GetDashboardSummary)
}