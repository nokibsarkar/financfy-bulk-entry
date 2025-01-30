package cashflow

import (
	"financify/bulk-entry/database"
	"financify/bulk-entry/models"
	"fmt"
)

type CashFlowService struct {
	CashbookID string
}

func (c *CashFlowService) ListCashFlowByCashbookID() []models.CashflowSingle {
	// List all cashflows by cashbook id
	db, close := database.GetDatabaseConnection()
	defer close()
	var cashflows []database.CashFlow
	res := db.Where(database.CashFlow{CashbookID: c.CashbookID}).Find(&cashflows).Order("id desc")
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
