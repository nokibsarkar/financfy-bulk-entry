package cashflow

import (
	"financify/bulk-entry/database"
	"financify/bulk-entry/models"
	"fmt"

	"gorm.io/gorm"
)

type CashFlowRepository struct{}

func (c *CashFlowRepository) CreateOrUpdateCashFlowByDate(db *gorm.DB, cashflow *models.CashflowSingle) (*models.CashflowSingle, error) {
	// This function is used to create or update the CashFlow
	cashfl := &database.CashFlow{}
	result := db.Where(database.CashFlow{Date: cashflow.Date, CashbookID: cashflow.CashbookID}).First(cashfl)
	fmt.Println(result)
	if result.Error != nil {
		// Create a new CashFlow
		newCashFlow := &database.CashFlow{
			CashbookID:    cashflow.CashbookID,
			TotalIncoming: cashflow.TotalIncoming,
			TotalOutgoing: cashflow.TotalOutgoing,
			TotalBalance:  cashflow.TotalBalance,
			Date:          cashflow.Date,
		}
		result := db.Create(newCashFlow)
		if result.Error != nil {
			return nil, result.Error
		}
		return cashflow, nil
	}
	// Update the existing CashFlow
	cashfl.TotalIncoming += cashflow.TotalIncoming
	cashfl.TotalOutgoing += cashflow.TotalOutgoing
	cashfl.TotalBalance += cashflow.TotalBalance
	result = db.Save(cashfl)
	if result.Error != nil {
		return nil, result.Error
	}
	return cashflow, nil
}
