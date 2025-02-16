package dtos

import "time"

type CreateTransactionRequest struct {
	UserID uint    `json:"user_id"`
	Amount float64 `json:"amount" validate:"required"`
	Status string  `json:"status"`
}

type UpdateTranscationRequest struct {
	Status string `json:"status" validate:"required"`
}

type TransactionResponse struct {
	ID        uint    `json:"id"`
	UserID    uint    `json:"user_id"`
	Amount    float64 `json:"amount"`
	Status    string  `json:"status"`
	CreatedAt time.Time  `json:"created_at"`
}

type DashboardSummaryResponse struct {
	TotalSuccessTransactions float64            `json:"total_success_transactions"`
	AverageTransactionPerUser float64           `json:"average_transaction_per_user"`
	LatestTransactions        []TransactionResponse `json:"latest_transactions"`
}
