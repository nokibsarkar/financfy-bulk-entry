package cashbook

import (
	"financify/bulk-entry/database"
	"financify/bulk-entry/models"
	"financify/bulk-entry/services"

	"gorm.io/gorm"
)

type CashbookRepository struct{}

func (c *CashbookRepository) CreateCashbook(db *gorm.DB, cashbook *models.CreateCashBookInput) (*database.Cashbook, error) {
	result := db.Create(&database.Cashbook{
		Name:          cashbook.Name,
		Description:   "",
		TotalIncoming: 0,
		TotalOutgoing: 0,
		TotalBalance:  0,
		Currency:      "BDT",
		ID:            services.GenerateSnowFlake(),
	})
	if result.Error != nil {
		return nil, result.Error
	}
	return nil, nil
}
func (c *CashbookRepository) LisCashBooks(db *gorm.DB) ([]database.Cashbook, error) {
	var cashbooks []database.Cashbook
	result := db.Find(&cashbooks)
	if result.Error != nil {
		return nil, result.Error
	}
	return cashbooks, nil
}
