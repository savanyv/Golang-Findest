package models

type Transaction struct {
	ID        uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	UserID    uint   `json:"user_id"`
	User      User   `json:"user" gorm:"foreignKey:UserID"`
	Amount    int    `json:"amount" gorm:"not null"`
	Status    string `json:"status" gorm:"not null"`
	CreatedAt string `json:"created_at" gorm:"not null"`
}