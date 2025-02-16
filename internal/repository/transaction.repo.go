package repository

import (
	"time"

	"github.com/savanyv/Golang-Findest/internal/models"
	"gorm.io/gorm"
)

type TransactionRepository interface {
	CreateTransaction(transaction *models.Transaction) (*models.Transaction, error)
	GetTransaction(userID *uint, status *string) ([]models.Transaction, error)
	GetTransactionByID(transactionID uint) (*models.Transaction, error)
	GetTotalSuccessfullTransaction(startDate, endDate time.Time) (float64, error)
	UpdateStatusTransaction(transactionID uint, status string) (*models.Transaction, error)
	DeleteTransaction(transactionID uint) error
}

type transactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) TransactionRepository {
	return &transactionRepository{
		db: db,
	}
}

func (r *transactionRepository) CreateTransaction(transaction *models.Transaction) (*models.Transaction, error) {
	if err := r.db.Create(transaction).Error; err != nil {
		return nil, err
	}

	return transaction, nil
}

func (r *transactionRepository) GetTransaction(userID *uint, status *string) ([]models.Transaction, error) {
	var transactions []models.Transaction
	query := r.db.Model(&models.Transaction{})

	if userID != nil {
		query = query.Where("user_id = ?", *userID)
	}

	if status != nil {
		query = query.Where("status = ?", status)
	}

	if err := query.Find(&transactions).Error; err != nil {
		return nil, err
	}

	return transactions, nil
}

func (r *transactionRepository) GetTransactionByID(transactionID uint) (*models.Transaction, error) {
	var transaction models.Transaction
	if err := r.db.Where("id = ?", transactionID).First(&transaction).Error; err != nil {
		return nil, err
	}

	return &transaction, nil
}

func (r *transactionRepository) GetTotalSuccessfullTransaction(startDate, endDate time.Time) (float64, error) {
	var total float64
	if err := r.db.Model(&models.Transaction{}).
		Where("status = ? AND created_at BETWEEN ? AND ?", "success", startDate, endDate).
		Select("SUM(amount)").Scan(&total).Error; err != nil {
			return 0, err
		}

		return total, nil
}

func (r *transactionRepository) UpdateStatusTransaction(transactionID uint, status string) (*models.Transaction, error) {
	var transaction models.Transaction
	if err := r.db.Where("id = ?", transactionID).First(&transaction).Error; err != nil {
		return nil, err
	}

	transaction.Status = status

	if err := r.db.Save(&transaction).Error; err != nil {
		return nil, err
	}

	return &transaction, nil
}

func (r *transactionRepository) DeleteTransaction(transactionID uint) error {
	var transaction models.Transaction
	if err := r.db.Delete(&transaction, transactionID).Error; err != nil {
		return err
	}

	return nil
}