package cashflow

import (
	"financify/bulk-entry/database"
	"financify/bulk-entry/models"
	"fmt"
)

type CashFlowService struct {
}

func (c *CashFlowService) ListCashFlowByCashbookID(CashbookID string) []models.CashflowSingle {
	// List all cashflows by cashbook id
	db, close := database.GetDatabaseConnection()
	defer close()
	var cashflows []database.CashFlow
	res := db.Where(database.CashFlow{CashbookID: CashbookID}).Find(&cashflows).Order("id desc")
	fmt.Println(res)
	cs := []models.CashflowSingle{}
	for _, cashflow := range cashflows {
		cs = append(cs, models.CashflowSingle{
			ID:             cashflow.ID,
			CashbookID:     cashflow.CashbookID,
			OpeningBalance: cashflow.OpeningBalance,
			TotalIncoming:  cashflow.TotalIncoming,
			TotalOutgoing:  cashflow.TotalOutgoing,
			TotalBalance:   cashflow.TotalBalance,
			ClosingBalance: cashflow.ClosingBalance,
			Date:           cashflow.Date,
		})
	}
	return cs
}
func (c *CashFlowService) GetSingleCashFlowByID(id string) *models.CashflowSingle {
	// Get a single cashflow by id
	db, close := database.GetDatabaseConnection()
	defer close()
	cashflow := &database.CashFlow{ID: id}
	res := db.Where(cashflow).First(cashflow)
	if res.Error != nil {
		return nil
	}
	return &models.CashflowSingle{
		ID:             cashflow.ID,
		CashbookID:     cashflow.CashbookID,
		OpeningBalance: cashflow.OpeningBalance,
		TotalIncoming:  cashflow.TotalIncoming,
		TotalOutgoing:  cashflow.TotalOutgoing,
		TotalBalance:   cashflow.TotalBalance,
		ClosingBalance: cashflow.ClosingBalance,
		Date:           cashflow.Date,
	}
}
