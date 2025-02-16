package transactiontest

import (
	"testing"
	"time"

	dtos "github.com/savanyv/Golang-Findest/internal/dto"
	"github.com/savanyv/Golang-Findest/internal/models"
	repositorytest "github.com/savanyv/Golang-Findest/internal/test/repository"
	"github.com/savanyv/Golang-Findest/internal/usecase"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Test CreateTransaction
func TestCreateTransaction(t *testing.T) {
	mockRepo := new(repositorytest.MockTransactionRepository)
	mockUserRepo := new(repositorytest.MockUserRepository)
	usecase := usecase.NewTransactionUsecase(mockRepo, mockUserRepo)

	userID := uint(1)
	request := &dtos.CreateTransactionRequest{
		Amount: 1000,
		Status: "pending",
	}

	mockUserRepo.On("GetUserByID", userID).Return(&models.User{ID: userID}, nil)

	expectedTransaction := &models.Transaction{
		ID:        1,
		UserID:    userID,
		Amount:    request.Amount,
		Status:    request.Status,
		CreatedAt: time.Now(),
	}

	mockRepo.On("CreateTransaction", mock.Anything).Return(expectedTransaction, nil)

	result, err := usecase.CreateTransaction(request, userID)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, expectedTransaction.ID, result.ID)
}

// Test GetTransaction
func TestGetTransaction(t *testing.T) {
	mockRepo := new(repositorytest.MockTransactionRepository)
	usecase := usecase.NewTransactionUsecase(mockRepo, nil)

	userID := uint(1)
	status := "success"

	mockTransactions := []models.Transaction{
		{ID: 1, UserID: userID, Amount: 1000, Status: "success", CreatedAt: time.Now()},
	}

	mockRepo.On("GetTransaction", &userID, &status).Return(mockTransactions, nil)

	result, err := usecase.GetTransaction(&userID, &status)

	assert.NoError(t, err)
	assert.Len(t, result, 1)
	assert.Equal(t, mockTransactions[0].ID, result[0].ID)
}

// Test GetTransactionByID
func TestGetTransactionByID(t *testing.T) {
	mockRepo := new(repositorytest.MockTransactionRepository)
	mockUserRepo := new(repositorytest.MockUserRepository)
	usecase := usecase.NewTransactionUsecase(mockRepo, mockUserRepo)

	userID := uint(1)
	transactionID := uint(1)

	mockUserRepo.On("GetUserByID", userID).Return(&models.User{ID: userID}, nil)

	mockTransaction := &models.Transaction{ID: transactionID, UserID: userID, Amount: 1000, Status: "success", CreatedAt: time.Now()}

	mockRepo.On("GetTransactionByID", transactionID).Return(mockTransaction, nil)

	result, err := usecase.GetTransactionByID(userID, transactionID)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, mockTransaction.ID, result.ID)
}

// Test UpdateStatusTransaction
func TestUpdateStatusTransaction(t *testing.T) {
	mockRepo := new(repositorytest.MockTransactionRepository)
	usecase := usecase.NewTransactionUsecase(mockRepo, nil)

	transactionID := uint(1)
	updateRequest := &dtos.UpdateTranscationRequest{Status: "success"}

	mockTransaction := &models.Transaction{ID: transactionID, Status: "pending"}

	mockRepo.On("GetTransactionByID", transactionID).Return(mockTransaction, nil)
	mockRepo.On("UpdateStatusTransaction", transactionID, updateRequest.Status).Return(mockTransaction, nil)

	result, err := usecase.UpdateStatusTransaction(transactionID, updateRequest)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, updateRequest.Status, result.Status)
}

// Test DeleteTransaction
func TestDeleteTransaction(t *testing.T) {
	mockRepo := new(repositorytest.MockTransactionRepository)
	usecase := usecase.NewTransactionUsecase(mockRepo, nil)

	userID := uint(1)
	transactionID := uint(1)

	mockTransaction := &models.Transaction{ID: transactionID, UserID: userID}

	mockRepo.On("GetTransactionByID", transactionID).Return(mockTransaction, nil)
	mockRepo.On("DeleteTransaction", transactionID).Return(nil)

	err := usecase.DeleteTransaction(userID, transactionID)

	assert.NoError(t, err)
}

// Test GetDashboardSummary
func TestGetDashboardSummary(t *testing.T) {
	mockRepo := new(repositorytest.MockTransactionRepository)
	usecase := usecase.NewTransactionUsecase(mockRepo, nil)

	startDate := time.Now().Truncate(24 * time.Hour)
	endDate := startDate.Add(24 * time.Hour)

	mockRepo.On("GetTotalSuccessfullTransaction", startDate, endDate).Return(5000.0, nil)
	mockRepo.On("GetAverageTransactionPerUser").Return(250.0, nil)
	mockRepo.On("GetLatestTransactions", 10).Return([]models.Transaction{
		{ID: 1, UserID: 2, Amount: 1000, Status: "success", CreatedAt: time.Now()},
	}, nil)

	result, err := usecase.GetDashboardSummary()

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, 5000.0, result[0].TotalSuccessTransactions)
	assert.Equal(t, 250.0, result[0].AverageTransactionPerUser)
	assert.Len(t, result[0].LatestTransactions, 1)
}