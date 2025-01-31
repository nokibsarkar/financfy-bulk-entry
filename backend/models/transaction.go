package models

import (
	"time"

	"gorm.io/datatypes"
)

type TransactionSingle struct {
	ID         string         `json:"id"`
	VoucherNo  string         `json:"voucherNo"`
	Date       datatypes.Date `json:"date"`
	Amount     float64        `json:"amount"`
	Contact    string         `json:"contact"`
	CashbookID string         `json:"cashbookId"`
	Reference  string         `json:"reference"`
	Remarks    string         `json:"remarks"`
	Category   string         `json:"category"`
}

type TransactionSingleInput struct {
	ID         string         `json:"id"`
	VoucherNo  string         `json:"voucherNo" binding:"required"`
	Date       datatypes.Date `json:"date" binding:"required"`
	Amount     float64        `json:"amount" binding:"required"`
	Contact    string         `json:"contact"`
	CashbookID string         `json:"cashbookId" binding:"required"`
	Reference  string         `json:"reference"`
	Remarks    string         `json:"remarks"`
	Category   string         `json:"category" binding:"required"`
	Type       string         `json:"type" binding:"required"`
}

type BulkTransactionResponse struct {
	SuccessCount int `json:"successCount"`
	FailedCount  int `json:"failedCount"`
}
type TransactionFilter struct {
	StartDate  *time.Time `json:"startDate"`
	EndDate    *time.Time `json:"endDate"`
	CashflowID string     `json:"cashflowId"`
}
