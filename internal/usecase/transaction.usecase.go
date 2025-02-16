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
	GetTransactionByID(userID, transactionID uint) (*dtos.TransactionResponse, error)
	UpdateStatusTransaction(transactionID uint, req *dtos.UpdateTranscationRequest) (*dtos.TransactionResponse, error)
	DeleteTransaction(userID, transactionID uint) error
	GetDashboardSummary() ([]*dtos.DashboardSummaryResponse, error)
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

func (u *transactionUsecase) GetTransactionByID(userID, transactionID uint) (*dtos.TransactionResponse, error) {
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

	// check if user is the owner of the transaction
	if transaction.UserID != userID {
		return nil, errors.New("you are not authorized to view this transaction")
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

func (u *transactionUsecase) UpdateStatusTransaction(transactionID uint, req *dtos.UpdateTranscationRequest) (*dtos.TransactionResponse, error) {
	// check if transaction exists
	transaction, err := u.repo.GetTransactionByID(transactionID)
	if err != nil {
		return nil, errors.New("failed get tranasaction by id")
	}

	transaction.Status = req.Status

	// update transaction
	transaction, err = u.repo.UpdateStatusTransaction(transactionID, req.Status)
	if err != nil {
		return nil, errors.New("failed to update transaction")
	}

	// resposne
	response := dtos.TransactionResponse{
		ID: transaction.ID,
		UserID: transaction.UserID,
		Amount: transaction.Amount,
		Status: transaction.Status,
		CreatedAt: transaction.CreatedAt,
	}

	return &response, nil
}

func (u *transactionUsecase) DeleteTransaction(userID, transactionID uint) error {
	// get transaction
	transaction, err := u.repo.GetTransactionByID(transactionID)
	if err != nil {
		return errors.New("failed to get transaction by id")
	}

	// check if user is the owner of the transaction
	if transaction.UserID != userID {
		return errors.New("unauthorized to delete this transaction")
	}

	// delete transaction
	err = u.repo.DeleteTransaction(transactionID)
	if err != nil {
		return errors.New("failed to delete transaction")
	}

	return nil
}

func (u *transactionUsecase) GetDashboardSummary() ([]*dtos.DashboardSummaryResponse, error) {
	// total transaction
	startDate := time.Now().Truncate(24 * time.Hour)
	endDate := startDate.Add(24 * time.Hour)
	totalTransaction, err := u.repo.GetTotalSuccessfullTransaction(startDate, endDate)
	if err != nil {
		return nil, errors.New("failed to get total transaction")
	}

	// avarage transaction
	avarageTransaction, err := u.repo.GetAverageTransactionPerUser()
	if err != nil {
		return nil, errors.New("failed to get avarage transaction")
	}

	// lastest transaction
	lastestTransaction, err := u.repo.GetLatestTransactions(10)
	if err != nil {
		return nil, errors.New("failed to get lastest transaction")
	}

	// convert to dto
	var lastestTransactionDTO []dtos.TransactionResponse
	for _, t := range lastestTransaction {
		lastestTransactionDTO = append(lastestTransactionDTO, dtos.TransactionResponse{
			ID:        t.ID,
			UserID:    t.UserID,
			Amount:    t.Amount,
			Status:    t.Status,
			CreatedAt: t.CreatedAt,
		})
	}

	// response
	response := &dtos.DashboardSummaryResponse{
		TotalSuccessTransactions: totalTransaction,
		AverageTransactionPerUser: avarageTransaction,
		LatestTransactions:        lastestTransactionDTO,
	}

	return []*dtos.DashboardSummaryResponse{response}, nil
}