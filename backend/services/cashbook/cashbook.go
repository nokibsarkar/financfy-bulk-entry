package cashbook

import (
	"financify/bulk-entry/models"
)

type CashBookService struct{}

func (c *CashBookService) GetCashBook() models.CashbookSingle {
	// implementatio
	sampleCashbook := models.CashbookSingle{
		ID:            1,
		Name:          "Sample Cashbook",
		Description:   "This is a sample cashbook",
		TotalIncoming: 1000.00,
		TotalOutgoing: 500.00,
		TotalBalance:  500.00,
	}
	return sampleCashbook
}
