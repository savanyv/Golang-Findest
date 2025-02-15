package usecase

import (
	"errors"
	"time"

	dtos "github.com/savanyv/Golang-Findest/internal/dto"
	"github.com/savanyv/Golang-Findest/internal/models"
	"github.com/savanyv/Golang-Findest/internal/repository"
)

type TransactionUsecase interface {
	CreateTransaction(req *dtos.CreateTransactionRequest, userID uint) (*dtos.TransactionResponse, error)
	GetTransaction(userID *uint, status *string) ([]dtos.TransactionResponse, error)
	GetTransactionByID(transactionID uint) (*dtos.TransactionResponse, error)
}

type transactionUsecase struct {
	repo     repository.TransactionRepository
	userRepo repository.UserRepository
}

func NewTransactionUsecase(repo repository.TransactionRepository, userRepo repository.UserRepository) TransactionUsecase {
	return &transactionUsecase{
		repo:     repo,
		userRepo: userRepo,
	}
}

func (u *transactionUsecase) CreateTransaction(req *dtos.CreateTransactionRequest, userID uint) (*dtos.TransactionResponse, error) {
	// check if user exists
	_, err := u.userRepo.GetUserByID(userID)
	if err != nil {
		return nil, errors.New("user not found")
	}

	// set status
	status := req.Status
	if status == "" {
		status = "pending"
	}

	// create transaction
	createdTransaction := &models.Transaction{
		UserID: userID,
		Amount: req.Amount,
		Status: status,
		CreatedAt: time.Now(),
	}

	// save transaction
	transaction, err := u.repo.CreateTransaction(createdTransaction)
	if err != nil {
		return nil, errors.New("failed to create transaction")
	}

	// response
	response := &dtos.TransactionResponse{
		ID:        transaction.ID,
		UserID:    transaction.UserID,
		Amount:    transaction.Amount,
		Status:    transaction.Status,
		CreatedAt: transaction.CreatedAt,
	}

	return response, nil
}

func (u *transactionUsecase) GetTransaction(userID *uint, status *string) ([]dtos.TransactionResponse, error) {
	// get transactions
	transaction, err := u.repo.GetTransaction(userID, status)
	if err != nil {
		return nil, errors.New("failed to get transactions")
	}

	// response
	var response []dtos.TransactionResponse
	for _, t := range transaction {
		response = append(response, dtos.TransactionResponse{
			ID:        t.ID,
			UserID:    t.UserID,
			Amount:    t.Amount,
			Status:    t.Status,
			CreatedAt: t.CreatedAt,
		})
	}

	return response, nil
}

func (u *transactionUsecase) GetTransactionByID(transactionID uint) (*dtos.TransactionResponse, error) {
	// check if user logged in
	_, err := u.userRepo.GetUserByID(transactionID)
	if err != nil {
		return nil, errors.New("user not logged in")
	}
	
	// get transaction
	transaction, err := u.repo.GetTransactionByID(transactionID)
	if err != nil {
		return nil, errors.New("failed to get transaction")
	}

	// return response
	response := &dtos.TransactionResponse{
		ID:        transaction.ID,
		UserID:    transaction.UserID,
		Amount:    transaction.Amount,
		Status:    transaction.Status,
		CreatedAt: transaction.CreatedAt,
	}

	return response, nil
}