package models

type TransactionSingle struct {
	ID         uint64  `json:"id"`
	VoucherNo  string  `json:"voucherNo"`
	Date       string  `json:"date"`
	Amount     float64 `json:"amount"`
	Contact    string  `json:"contact"`
	CashbookID int     `json:"cashbookId"`
	Reference  string  `json:"reference"`
	Remarks    string  `json:"remarks"`
	Category   string  `json:"category"`
}

type TransactionSingleInput struct {
	VoucherNo  string  `json:"voucherNo" binding:"required"`
	Date       string  `json:"date" binding:"required"`
	Amount     float64 `json:"amount" binding:"required"`
	Contact    string  `json:"contact"`
	CashbookID int     `json:"cashbookId" binding:"required"`
	Reference  string  `json:"reference"`
	Remarks    string  `json:"remarks"`
	Category   string  `json:"category" binding:"required"`
	Type       string  `json:"type" binding:"required"`
}
type BulkTransactionResponse struct {
	SuccessCount int `json:"successCount"`
	FailedCount  int `json:"failedCount"`
}
