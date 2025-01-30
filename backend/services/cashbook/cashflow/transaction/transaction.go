package transaction

import (
	"financify/bulk-entry/database"
	"financify/bulk-entry/models"
	cashflow_repo "financify/bulk-entry/repositories/cashflow"
	transaction_repo "financify/bulk-entry/repositories/transaction"
	"log"

	"github.com/godruoyi/go-snowflake"
)

var transactionList = []models.TransactionSingle{
	{
		ID:         1,
		VoucherNo:  "1234",
		Date:       "2021-01-01",
		Amount:     1000.00,
		Contact:    "John Doe",
		CashbookID: 1,
		Reference:  "Ref-1234",
		Remarks:    "Sample Transaction",
	},
}

type TransactionService struct{}

func (t *TransactionService) CreateTransaction(transactionInput *models.TransactionSingleInput) *models.TransactionSingle {
	newTransactionId := snowflake.ID()
	newTransaction := &models.TransactionSingleInput{
		ID:         newTransactionId,
		VoucherNo:  transactionInput.VoucherNo,
		Date:       transactionInput.Date,
		Amount:     transactionInput.Amount,
		Contact:    transactionInput.Contact,
		CashbookID: transactionInput.CashbookID,
		Reference:  transactionInput.Reference,
		Remarks:    transactionInput.Remarks,
		Category:   transactionInput.Category,
		Type:       transactionInput.Type,
	}
	repo := transaction_repo.TransactionRepository{}
	cashflow_repo := cashflow_repo.CashFlowRepository{}
	cashflowUpdate := &models.CashflowSingle{
		CashbookID:    newTransaction.CashbookID,
		TotalIncoming: 0.00,
		TotalOutgoing: 0.00,
		TotalBalance:  0.00,
		Date:          newTransaction.Date,
	}
	if newTransaction.Type == "cashin" {
		cashflowUpdate.TotalIncoming = newTransaction.Amount
		cashflowUpdate.TotalBalance = newTransaction.Amount
	} else {
		cashflowUpdate.TotalOutgoing = newTransaction.Amount
		cashflowUpdate.TotalBalance = -newTransaction.Amount
	}
	db := database.GetDatabaseConnection().Begin()
	_, err := repo.CreateTransaction(db, newTransaction)
	if err != nil {
		return nil
	}
	_, err = cashflow_repo.CreateOrUpdateCashFlowByDate(db, cashflowUpdate)
	if err != nil {
		return nil
	}
	db.Commit()
	newTransactionSingle := &models.TransactionSingle{
		ID:         newTransactionId,
		VoucherNo:  newTransaction.VoucherNo,
		Date:       newTransaction.Date,
		Amount:     newTransaction.Amount,
		Contact:    newTransaction.Contact,
		CashbookID: newTransaction.CashbookID,
		Reference:  newTransaction.Reference,
		Remarks:    newTransaction.Remarks,
		Category:   newTransaction.Category,
	}
	return newTransactionSingle
}
func (t *TransactionService) CreateBulkTransactions(transactions []models.TransactionSingleInput) models.BulkTransactionResponse {
	// Create multiple transactions at once
	resp := models.BulkTransactionResponse{
		SuccessCount: 0,
		FailedCount:  0,
	}

	cashflowChanges := map[string]*models.CashflowSingle{}
	for _, transaction := range transactions {
		existingCashflow, ok := cashflowChanges[transaction.Date]
		if !ok {
			existingCashflow = &models.CashflowSingle{
				CashbookID:    transaction.CashbookID,
				TotalIncoming: 0.00,
				TotalOutgoing: 0.00,
				TotalBalance:  0.00,
				Date:          transaction.Date,
			}
		}
		if transaction.Type == "cashin" {
			existingCashflow.TotalIncoming += transaction.Amount
			existingCashflow.TotalBalance += transaction.Amount
		} else {
			existingCashflow.TotalOutgoing += transaction.Amount
			existingCashflow.TotalBalance -= transaction.Amount
		}
		cashflowChanges[transaction.Date] = existingCashflow
		resp.SuccessCount++
	}
	repo := transaction_repo.TransactionRepository{}
	cashflow_repo := cashflow_repo.CashFlowRepository{}
	db := database.GetDatabaseConnection().Begin()
	_a, err := repo.CreateBulkTransactions(db, transactions)
	_ = _a
	if err != nil {
		resp.FailedCount++
		log.Fatalln(err)
	}
	for _, cashflow := range cashflowChanges {
		_, err := cashflow_repo.CreateOrUpdateCashFlowByDate(db, cashflow)
		if err != nil {
			resp.FailedCount++
		}
	}
	db.Commit()
	return resp

}
func (t *TransactionService) UpdateSingleTransaction() {}
func (t *TransactionService) GetSingleTransaction(id uint64) *models.TransactionSingle {
	// search for transaction by id
	for _, transaction := range transactionList {
		if transaction.ID == id {
			return &transaction
		}
	}
	return nil
}
func (t *TransactionService) ListAllTransactions() ([]models.TransactionSingle, error) {
	db := database.GetDatabaseConnection()
	repo := transaction_repo.TransactionRepository{}
	transactions, err := repo.ListAllTransactions(db)
	return transactions, err
}
