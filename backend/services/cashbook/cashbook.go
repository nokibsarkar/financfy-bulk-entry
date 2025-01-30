package cashbook

import (
	"financify/bulk-entry/models"

	"github.com/godruoyi/go-snowflake"
)

type CashBookService struct{}

func (c *CashBookService) ListCashbooks() []models.CashbookSingle {
	// implementatio
	sampleCashbook := models.CashbookSingle{
		ID:            1,
		Name:          "Sample Cashbook",
		Description:   "This is a sample cashbook",
		TotalIncoming: 1000.00,
		TotalOutgoing: 500.00,
		TotalBalance:  500.00,
	}
	sampleCashbook2 := models.CashbookSingle{
		ID:            2,
		Name:          "Sample Cashbook 2",
		Description:   "This is a sample cashbook 2",
		TotalIncoming: 2000.00,
		TotalOutgoing: 1000.00,
		TotalBalance:  1000.00,
	}
	return []models.CashbookSingle{sampleCashbook, sampleCashbook2}
}

func (c *CashBookService) CreateCashbook(models.CreateCashBookInput) models.CashbookSingle {
	// Create a new cashbook

	newCashBookId := snowflake.ID()
	newCashbook := models.CashbookSingle{
		ID:            newCashBookId,
		Name:          "New Cashbook",
		Description:   "This is a new cashbook",
		TotalIncoming: 0.00,
		TotalOutgoing: 0.00,
		TotalBalance:  0.00,
	}
	return newCashbook
}
func (c *CashBookService) GetSingleCashBook() models.CashbookSingle {
	// implementation
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
