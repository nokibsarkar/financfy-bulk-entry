package database

import (
	"fmt"
	"log"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

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
func (c *CashFlow) Propagate(db *gorm.DB) {

	YCB := &CashFlow{CashbookID: c.CashbookID}
	// Fetch the immediate previous cashflow
	res := db.Where(YCB).Where("date < ?", c.Date).Order("date desc").First(YCB)
	fmt.Printf("Yesterday's Cashflow: %v\n", *YCB)
	if res.Error != nil {
		c.OpeningBalance = 0
	} else {
		c.OpeningBalance = YCB.ClosingBalance
	}
	fmt.Println("Opening Balance: ", c.OpeningBalance)
	fmt.Println("Total Incoming: ", c.TotalIncoming)
	fmt.Println("Total Outgoing: ", c.TotalOutgoing)
	fmt.Println("Total Balance: ", c.TotalBalance)
	c.TotalBalance = c.TotalIncoming - c.TotalOutgoing
	c.ClosingBalance = c.OpeningBalance + c.TotalBalance
	db.Save(c)

	allSubsequentCashflows := []CashFlow{}
	// Get all subsequent cashflows
	res = db.Where(CashFlow{CashbookID: c.CashbookID}).Where("date > ?", c.Date).Order("date asc").Find(&allSubsequentCashflows)
	if res.Error != nil {
		log.Println(res.Error)
		return
	}
	// Now propagate the changes to all subsequent cashflows
	var lastClosingBalance float64 = c.ClosingBalance
	for _, cashflow := range allSubsequentCashflows {
		cashflow.OpeningBalance = lastClosingBalance
		cashflow.TotalBalance = cashflow.TotalIncoming - cashflow.TotalOutgoing
		cashflow.ClosingBalance = cashflow.OpeningBalance + cashflow.TotalBalance
		db.Save(cashflow)
		lastClosingBalance = cashflow.ClosingBalance
	}
	// Update the cashbook
	cbook := &Cashbook{ID: c.CashbookID}
	db.First(cbook)
	cbook.TotalIncoming += c.TotalIncoming
	cbook.TotalOutgoing += c.TotalOutgoing
	cbook.TotalBalance += c.TotalBalance
	db.Save(cbook)
	fmt.Println((cbook))
}
