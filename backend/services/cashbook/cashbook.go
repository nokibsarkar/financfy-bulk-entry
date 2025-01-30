package cashbook

import (
	"financify/bulk-entry/database"
	"financify/bulk-entry/models"
	cashbook_repository "financify/bulk-entry/repositories/cashbook"

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

func (c *CashBookService) CreateCashbook(inp models.CreateCashBookInput) (*models.CashbookSingle, error) {
	// Create a new cashbook

	// every validation is done in the controller
	newCashBookId := snowflake.ID()
	newCashbook := &models.CashbookSingle{
		ID:            newCashBookId,
		Name:          "New Cashbook",
		Description:   "This is a new cashbook",
		TotalIncoming: 0.00,
		TotalOutgoing: 0.00,
		TotalBalance:  0.00,
	}
	repo := cashbook_repository.CashbookRepository{}
	db := database.GetDatabaseConnection()
	_, err := repo.CreateCashbook(db, &inp)
	if err != nil {
		return nil, err
	}
	return newCashbook, nil
}
func (c *CashBookService) GetSingleCashBook(id uint64) *models.CashbookSingle {
	// implementation
	sampleCashbook, err := tempCahbooks[id]
	if !err {
		return nil
	}
	return sampleCashbook
}
