package models

import "gorm.io/datatypes"

type CashflowSingle struct {
	ID            string         `json:"id"`
	Date          datatypes.Date `json:"date"`
	TotalIncoming float64        `json:"totalIn"`
	TotalOutgoing float64        `json:"totalOut"`
	TotalBalance  float64        `json:"netBalance"`
	CashbookID    string         `json:"cashbookId"`
}
