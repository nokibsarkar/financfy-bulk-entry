package database

import "gorm.io/datatypes"

type CashFlow struct {
	ID             string         `json:"id"`
	Date           datatypes.Date `json:"date"`
	Amount         float64        `json:"amount"`
	OpeningBalance float64        `json:"openingBalance"`
	TotalIncoming  float64        `json:"totalIncoming"`
	TotalOutgoing  float64        `json:"totalOutgoing"`
	TotalBalance   float64        `json:"totalBalance"`
	ClosingBalance float64        `json:"closingBalance"`
	CashbookID     string         `json:"cashbookId"`
	Cashbook       Cashbook       `json:"-" gorm:"foreignKey:CashbookID"`
}

func (c *CashFlow) TableName() string {
	return "cashflow"
}
