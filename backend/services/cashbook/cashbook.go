package cashbook

import (
	"financify/bulk-entry/models"

	"github.com/godruoyi/go-snowflake"
)

var tempCahbooks = make(map[uint64]*models.CashbookSingle, 0)

type CashBookService struct{}

func (c *CashBookService) ListCashbooks() []models.CashbookSingle {
	// implementatio
	// iterate over the cashbooks and return them
	cashbooks := []models.CashbookSingle{}
	for _, cashbook := range tempCahbooks {
		cashbooks = append(cashbooks, *cashbook)
	}
	return cashbooks
}

func (c *CashBookService) CreateCashbook(models.CreateCashBookInput) *models.CashbookSingle {
	// Create a new cashbook

	newCashBookId := snowflake.ID()
	newCashbook := &models.CashbookSingle{
		ID:            newCashBookId,
		Name:          "New Cashbook",
		Description:   "This is a new cashbook",
		TotalIncoming: 0.00,
		TotalOutgoing: 0.00,
		TotalBalance:  0.00,
	}
	tempCahbooks[newCashBookId] = newCashbook
	return newCashbook
}
func (c *CashBookService) GetSingleCashBook(id uint64) *models.CashbookSingle {
	// implementation
	sampleCashbook, err := tempCahbooks[id]
	if !err {
		return nil
	}
	return sampleCashbook
}
