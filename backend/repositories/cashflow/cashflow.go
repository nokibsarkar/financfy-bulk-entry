package cashflow

import (
	"financify/bulk-entry/database"
	"financify/bulk-entry/models"
	"financify/bulk-entry/services"
	"fmt"

	"gorm.io/gorm"
)

type CashFlowRepository struct{}

// This function is used to create or update the CashFlow
func (c *CashFlowRepository) CreateOrUpdateCashFlow(db *gorm.DB, cashflow *models.CashflowSingle) (*database.CashFlow, error) {
	cashflow_in_db := &database.CashFlow{Date: cashflow.Date, CashbookID: cashflow.CashbookID}

	result := db.Where(cashflow_in_db).First(cashflow_in_db)
	if result.Error != nil {
		if result.Error.Error() == "record not found" {
			fmt.Println("Creating new CashFlow", cashflow.CashbookID)
			new_cashflow := &database.CashFlow{
				ID:            services.GenerateSnowFlake(),
				TotalIncoming: cashflow.TotalIncoming,
				TotalOutgoing: cashflow.TotalOutgoing,
				TotalBalance:  cashflow.TotalBalance,
				CashbookID:    cashflow.CashbookID,
				Date:          cashflow.Date,
			}
			result := db.Create(new_cashflow)
			if result.Error != nil {
				return nil, result.Error
			}
			return new_cashflow, nil
		} else {
			return nil, result.Error
		}
	}
	// Update the existing CashFlow
	cashflow_in_db.TotalIncoming += cashflow.TotalIncoming
	cashflow_in_db.TotalOutgoing += cashflow.TotalOutgoing
	cashflow_in_db.TotalBalance += cashflow.TotalBalance
	cashflow_in_db.Propagate(db)
	result = db.Save(cashflow_in_db)
	if result.Error != nil {
		return nil, result.Error
	}
	return cashflow_in_db, nil
}
func (c *CashFlowRepository) ListCashFlowByCashbookID(db *gorm.DB, CashbookID string) []models.CashflowSingle {
	var cashflows []database.CashFlow
	db.Where(database.CashFlow{CashbookID: CashbookID}).Find(&cashflows).Order("id desc")
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
