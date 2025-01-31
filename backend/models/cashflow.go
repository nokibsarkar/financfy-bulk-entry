package models

import "gorm.io/datatypes"

type CashflowSingle struct {
	ID             string         `json:"id"`
	Date           datatypes.Date `json:"date"`
	OpeningBalance float64        `json:"openingBalance"`
	TotalIncoming  float64        `json:"totalIn"`
	TotalOutgoing  float64        `json:"totalOut"`
	TotalBalance   float64        `json:"netBalance"`
	ClosingBalance float64        `json:"closingBalance"`
	CashbookID     string         `json:"cashbookId"`
}
type CashflowFilter struct {
	CashbookID string `json:"cashbookId"`
}
