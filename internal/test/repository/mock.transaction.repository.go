package repositorytest

import (
	"time"

	"github.com/savanyv/Golang-Findest/internal/models"
	"github.com/stretchr/testify/mock"
)

type MockTransactionRepository struct {
	mock.Mock
}

func (m *MockTransactionRepository) CreateTransaction(transaction *models.Transaction) (*models.Transaction, error) {
	args := m.Called(transaction)
	return args.Get(0).(*models.Transaction), args.Error(1)
}

func (m *MockTransactionRepository) GetTransaction(userID *uint, status *string) ([]models.Transaction, error) {
	args := m.Called(userID, status)
	return args.Get(0).([]models.Transaction), args.Error(1)
}

func (m *MockTransactionRepository) GetTransactionByID(transactionID uint) (*models.Transaction, error) {
	args := m.Called(transactionID)
	return args.Get(0).(*models.Transaction), args.Error(1)
}

func (m *MockTransactionRepository) UpdateStatusTransaction(transactionID uint, status string) (*models.Transaction, error) {
	args := m.Called(transactionID, status)
	return args.Get(0).(*models.Transaction), args.Error(1)
}

func (m *MockTransactionRepository) DeleteTransaction(transactionID uint) error {
	args := m.Called(transactionID)
	return args.Error(0)
}

func (m *MockTransactionRepository) GetTotalSuccessfullTransaction(startDate, endDate time.Time) (float64, error) {
	args := m.Called(startDate, endDate)
	return args.Get(0).(float64), args.Error(1)
}

func (m *MockTransactionRepository) GetAverageTransactionPerUser() (float64, error) {
	args := m.Called()
	return args.Get(0).(float64), args.Error(1)
}

func (m *MockTransactionRepository) GetLatestTransactions(limit int) ([]models.Transaction, error) {
	args := m.Called(limit)
	return args.Get(0).([]models.Transaction), args.Error(1)
}

