package transaction

import (
	"financify/bulk-entry/database"
	"financify/bulk-entry/models"
)

type TransactionRepository struct{}

func (t *TransactionRepository) CreateTransaction(transaction *models.TransactionSingleInput) (*models.TransactionSingle, error) {
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
	}
	db := database.GetDatabaseConnection()
	result := db.Create(newTransaction)
	if result.Error != nil {
		return nil, result.Error
	}
	return nil, nil
}
func (t *TransactionRepository) CreateBulkTransactions(transactions []models.TransactionSingle) (*models.TransactionSingle, error) {
	// Create multiple transactions at once
	return nil, nil
}

func (t *TransactionRepository) GetSingleTransaction(id uint64) *models.TransactionSingle {
	// search for transaction by id
	return nil
}
func (t *TransactionRepository) ListAllTransactions() []models.TransactionSingle {
	return nil
}
func (t *TransactionRepository) UpdateSingleTransaction()          {}
func (t *TransactionRepository) DeleteSingleTransaction()          {}
func (t *TransactionRepository) GetSingleTransactionByCashbookID() {}
func (t *TransactionRepository) ListAllTransactionsByCashbookID()  {}
func (t *TransactionRepository) ListAllTransactionsByCategory()    {}
func (t *TransactionRepository) ListAllTransactionsByContact()     {}
func (t *TransactionRepository) ListAllTransactionsByDate()        {}
func (t *TransactionRepository) ListAllTransactionsByReference()   {}
