package models

type CashflowSingle struct {
	ID            uint64  `json:"id"`
	Date          string  `json:"date"`
	TotalIncoming float64 `json:"totalIn"`
	TotalOutgoing float64 `json:"totalOut"`
	TotalBalance  float64 `json:"netBalance"`
	CashbookID    uint64  `json:"cashbookId"`
}
