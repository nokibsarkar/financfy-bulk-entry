package cashflow

import (
	"financify/bulk-entry/database"
	"financify/bulk-entry/models"
	"financify/bulk-entry/services"

	"gorm.io/gorm"
)

type CashFlowRepository struct{}

// This function is used to create or update the CashFlow
func (c *CashFlowRepository) CreateOrUpdateCashFlow(db *gorm.DB, cashflow *models.CashflowSingle) (*database.CashFlow, error) {
	cashflow_in_db := &database.CashFlow{Date: cashflow.Date, CashbookID: cashflow.CashbookID}
	result := db.Where(cashflow_in_db).First(cashflow_in_db)
	if result.Error != nil {
		if result.Error.Error() == "record not found" {
			cashflow_in_db.ID = services.GenerateSnowFlake()
			cashflow_in_db.TotalIncoming = cashflow.TotalIncoming
			cashflow_in_db.TotalOutgoing = cashflow.TotalOutgoing
			cashflow_in_db.TotalBalance = cashflow.TotalBalance
			cashflow_in_db.CashbookID = cashflow.CashbookID
			cashflow_in_db.Date = cashflow.Date
			result := db.Create(cashflow_in_db)
			if result.Error != nil {
				return nil, result.Error
			}
			return cashflow_in_db, nil
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
