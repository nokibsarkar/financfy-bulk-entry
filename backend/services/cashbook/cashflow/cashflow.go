package cashflow

import (
	"financify/bulk-entry/database"
	"financify/bulk-entry/models"
	cashflow_repo "financify/bulk-entry/repositories/cashflow"
)

type CashFlowService struct{}

func (c *CashFlowService) ListCashFlowByCashbookID(CashbookID string) []models.CashflowSingle {
	// List all cashflows by cashbook id
	db, close := database.GetDatabaseConnection()
	defer close()
	service := cashflow_repo.CashFlowRepository{}
	return service.ListCashFlowByCashbookID(db, CashbookID)
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
