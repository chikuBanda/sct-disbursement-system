package models

import "time"

type Beneficiary struct {
	ID            uint       `json:"id"`
	Name          string     `json:"name" form:"name" binding:"required"`
	AccountNumber string     `json:"account_number" form:"account_number" binding:"required"`
	Transfers     []Transfer `json:"transfers"` // One-To-Many relationship with 'transfers' table
	CreatedAt     time.Time  `json:"created_at"`
	UpdatedAt     time.Time  `json:"updated_at"`
}

type Transfer struct {
	ID            uint      `json:"id"`
	BeneficiaryId uint      `json:"beneficiary_id" form:"beneficiary_id" binding:"required"`
	Amount        float64   `json:"amount" form:"amount" binding:"required"`
	Status        string    `json:"status" form:"status" binding:"required"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type TransactionHistory struct {
	ID         uint      `json:"id"`
	TransferID uint      `json:"transfer_id" form:"beneficiary_id" binding:"required"`
	Timestamp  time.Time `json:"timestamp"`
	Status     string    `json:"status" form:"status" binding:"required"`
	Transfer   Transfer  `json:"transfer"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
