package cashflow

import (
	"financify/bulk-entry/database"
	"financify/bulk-entry/models"
	"fmt"
)

type CashFlowService struct {
	CashbookID uint64
}

func (c *CashFlowService) ListCashFlowByCashbookID() []models.CashflowSingle {
	// List all cashflows by cashbook id
	db := database.GetDatabaseConnection()
	var cashflows []database.CashFlow
	res := db.Where(database.CashFlow{CashbookID: c.CashbookID}).Find(&cashflows)
	fmt.Println(res)
	cs := []models.CashflowSingle{}
	for _, cashflow := range cashflows {
		cs = append(cs, models.CashflowSingle{
			CashbookID:    cashflow.CashbookID,
			TotalIncoming: cashflow.TotalIncoming,
			TotalOutgoing: cashflow.TotalOutgoing,
			TotalBalance:  cashflow.TotalBalance,
			Date:          cashflow.Date,
		})
	}
	return cs
}
