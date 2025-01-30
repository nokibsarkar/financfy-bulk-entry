package models

type TransactionSingle struct {
	ID         int     `json:"id"`
	VoucherNo  string  `json:"voucher_no"`
	Date       string  `json:"date"`
	Amount     float64 `json:"amount"`
	Contact    string  `json:"contact"`
	CashbookID int     `json:"cashbook_id"`
	Reference  string  `json:"reference"`
	Remarks    string  `json:"remarks"`
}
