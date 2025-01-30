package models

type CashflowSingle struct {
	ID            uint64  `json:"id"`
	Date          string  `json:"date"`
	TotalIncoming float64 `json:"totalIncoming"`
	TotalOutgoing float64 `json:"totalOutgoing"`
	TotalBalance  float64 `json:"totalBalance"`
	CashbookID    uint64  `json:"cashbookId"`
}
