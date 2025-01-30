package transaction

import (
	"financify/bulk-entry/models"

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
	newTransaction := models.TransactionSingle{
		ID:         newTransactionId,
		VoucherNo:  transactionInput.VoucherNo,
		Date:       transactionInput.Date,
		Amount:     transactionInput.Amount,
		Contact:    transactionInput.Contact,
		CashbookID: transactionInput.CashbookID,
		Reference:  transactionInput.Reference,
		Remarks:    transactionInput.Remarks,
		Category:   transactionInput.Category,
	}
	transactionList = append(transactionList, newTransaction)
	return &newTransaction
}
func (t *TransactionService) CreateBulkTransactions(transactions []models.TransactionSingleInput) models.BulkTransactionResponse {
	// Create multiple transactions at once
	resp := models.BulkTransactionResponse{
		SuccessCount: 0,
		FailedCount:  0,
	}
	for _, transaction := range transactions {
		transactionList = append(transactionList, models.TransactionSingle{
			ID:         snowflake.ID(),
			VoucherNo:  transaction.VoucherNo,
			Date:       transaction.Date,
			Amount:     transaction.Amount,
			Contact:    transaction.Contact,
			CashbookID: transaction.CashbookID,
			Reference:  transaction.Reference,
			Remarks:    transaction.Remarks,
			Category:   transaction.Category,
		})
		resp.SuccessCount++
	}
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
func (t *TransactionService) ListAllTransactions() []models.TransactionSingle {
	return transactionList
}
