package dtos

import "time"

type CreateTransactionRequest struct {
	UserID uint    `json:"user_id"`
	Amount float64 `json:"amount" validate:"required"`
	Status string  `json:"status"`
}

type TransactionRequest struct {
	UserID uint    `json:"user_id"`
	Amount float64 `json:"amount"`
	Status string  `json:"status"`
}

type TransactionResponse struct {
	ID        uint    `json:"id"`
	UserID    uint    `json:"user_id"`
	Amount    float64 `json:"amount"`
	Status    string  `json:"status"`
	CreatedAt time.Time  `json:"created_at"`
}