package models

import "time"

type Transaction struct {
	ID        uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	UserID    uint   `json:"user_id"`
	User      User   `json:"user" gorm:"foreignKey:UserID"`
	Amount    float64    `json:"amount" gorm:"not null"`
	Status    string `json:"status" gorm:"not null; default:pending"`
	CreatedAt time.Time `json:"created_at" gorm:"not null"`
}