package transaction

import (
	"financify/bulk-entry/database"
	"financify/bulk-entry/models"
	"financify/bulk-entry/services"
	"log"

	"gorm.io/gorm"
)

type TransactionRepository struct{}

func (t *TransactionRepository) CreateTransaction(db *gorm.DB, transaction *models.TransactionSingleInput) (*models.TransactionSingle, error) {
	// Create a new transaction
	newTransaction := &database.Transaction{
		VoucherNo: transaction.VoucherNo,
		Date:      transaction.Date,
		Amount:    transaction.Amount,
		Contact:   transaction.Contact,
		Cashbook:  transaction.CashbookID,
		Reference: transaction.Reference,
		Remarks:   transaction.Remarks,
		Category:  transaction.Category,
		ID:        services.GenerateSnowFlake(),
	}
	result := db.Create(newTransaction)
	if result.Error != nil {
		return nil, result.Error
	}
	return nil, nil
}
func (t *TransactionRepository) CreateBulkTransactions(db *gorm.DB, transactions []models.TransactionSingleInput) (*models.TransactionSingle, error) {
	// Create multiple transactions at once
	trs := []*database.Transaction{}
	for _, transaction := range transactions {
		newTransaction := &database.Transaction{
			ID:        services.GenerateSnowFlake(),
			VoucherNo: transaction.VoucherNo,
			Date:      transaction.Date,
			Amount:    transaction.Amount,
			Contact:   transaction.Contact,
			Cashbook:  transaction.CashbookID,
			Reference: transaction.Reference,
			Remarks:   transaction.Remarks,
			Category:  transaction.Category,
		}
		trs = append(trs, newTransaction)
	}
	result := db.Create(trs)
	if result.Error != nil {
		return nil, result.Error
	}
	log.Println("Bulk transactions created", result.RowsAffected)
	return nil, nil
}

func (t *TransactionRepository) GetSingleTransaction(id string) *models.TransactionSingle {
	// search for transaction by id
	return nil
}
func (t *TransactionRepository) ListAllTransactions(db *gorm.DB) ([]models.TransactionSingle, error) {
	// List all transactions
	transactions := []database.Transaction{}
	result := db.Find(&transactions)
	if result.Error != nil {
		return nil, result.Error
	}
	f_transactions := []models.TransactionSingle{}
	for _, transaction := range transactions {
		f_transactions = append(f_transactions, models.TransactionSingle{
			ID:         transaction.ID,
			VoucherNo:  transaction.VoucherNo,
			Date:       transaction.Date,
			Amount:     transaction.Amount,
			Contact:    transaction.Contact,
			CashbookID: transaction.Cashbook,
			Reference:  transaction.Reference,
			Remarks:    transaction.Remarks,
			Category:   transaction.Category,
		})
	}
	return f_transactions, nil
}
func (t *TransactionRepository) UpdateSingleTransaction()          {}
func (t *TransactionRepository) DeleteSingleTransaction()          {}
func (t *TransactionRepository) GetSingleTransactionByCashbookID() {}
func (t *TransactionRepository) ListAllTransactionsByCashbookID()  {}
func (t *TransactionRepository) ListAllTransactionsByCategory()    {}
func (t *TransactionRepository) ListAllTransactionsByContact()     {}
func (t *TransactionRepository) ListAllTransactionsByDate()        {}
func (t *TransactionRepository) ListAllTransactionsByReference()   {}
