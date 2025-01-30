package models

type CashflowSingle struct {
	ID            int     `json:"id"`
	Date          string  `json:"date"`
	TotalIncoming float64 `json:"total_incoming"`
	TotalOutgoing float64 `json:"total_outgoing"`
	TotalBalance  float64 `json:"total_balance"`
	CashbookID    int     `json:"cashbook_id"`
}
