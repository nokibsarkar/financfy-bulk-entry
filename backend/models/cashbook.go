package models

type CashbookSingle struct {
	ID            int     `json:"id"`
	Name          string  `json:"name"`
	Description   string  `json:"description"`
	TotalIncoming float64 `json:"total_incoming"`
	TotalOutgoing float64 `json:"total_outgoing"`
	TotalBalance  float64 `json:"total_balance"`
	Currency      string  `json:"currency"`
}

// type Cash
