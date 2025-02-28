package cashbook

import (
	"financify/bulk-entry/database"
	"financify/bulk-entry/models"
	cashbook_repository "financify/bulk-entry/repositories/cashbook"
	"financify/bulk-entry/services"
)

var tempCahbooks = make(map[string]*models.CashbookSingle, 0)

type CashBookService struct{}

func (c *CashBookService) ListCashbooks() []models.CashbookSingle {
	// implementatio
	// iterate over the cashbooks and return them
	repo := cashbook_repository.CashbookRepository{}
	db, close := database.GetDatabaseConnection()
	defer close()
	cashbooks, _ := repo.LisCashBooks(db)
	var cashs = make([]models.CashbookSingle, 0)
	for _, cashbook := range cashbooks {
		cash := models.CashbookSingle{
			ID:            cashbook.ID,
			Name:          cashbook.Name,
			Description:   cashbook.Description,
			TotalIncoming: cashbook.TotalIncoming,
			TotalOutgoing: cashbook.TotalOutgoing,
			TotalBalance:  cashbook.TotalBalance,
		}
		cashs = append(cashs, cash)
	}
	return cashs
}

func (c *CashBookService) CreateCashbook(inp models.CreateCashBookInput) (*models.CashbookSingle, error) {
	// Create a new cashbook

	// every validation is done in the controller
	newCashBookId := services.GenerateSnowFlake()
	newCashbook := &models.CashbookSingle{
		ID:            newCashBookId,
		Name:          inp.Name,
		Description:   inp.Name,
		Currency:      "BDT",
		TotalIncoming: 0.00,
		TotalOutgoing: 0.00,
		TotalBalance:  0.00,
	}
	repo := cashbook_repository.CashbookRepository{}
	db, close := database.GetDatabaseConnection()
	defer close()
	_, err := repo.CreateCashbook(db, &inp)
	if err != nil {
		return nil, err
	}
	return newCashbook, nil
}
func (c *CashBookService) GetSingleCashBook(id string) *models.CashbookSingle {
	// implementation
	sampleCashbook, err := tempCahbooks[id]
	if !err {
		return nil
	}
	return sampleCashbook
}
