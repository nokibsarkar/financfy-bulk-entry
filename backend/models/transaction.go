package models

type TransactionSingle struct {
	ID         int     `json:"id"`
	VoucherNo  string  `json:"voucherNo"`
	Date       string  `json:"date"`
	Amount     float64 `json:"amount"`
	Contact    string  `json:"contact"`
	CashbookID int     `json:"cashbookId"`
	Reference  string  `json:"reference"`
	Remarks    string  `json:"remarks"`
	Category   string  `json:"category"`
}
