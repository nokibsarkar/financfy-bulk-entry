package transaction

import (
	"financify/bulk-entry/consts"
	"financify/bulk-entry/database"
	"financify/bulk-entry/models"
	cashflow_repo "financify/bulk-entry/repositories/cashflow"
	transaction_repo "financify/bulk-entry/repositories/transaction"
	"financify/bulk-entry/services"
	"sort"
	"time"

	"fmt"
	"log"
)

type TransactionService struct{}

func (t *TransactionService) CreateTransaction(transactionInput *models.TransactionSingleInput) *models.TransactionSingle {
	newTransactionId := services.GenerateSnowFlake()
	fmt.Println("Creating new transaction cashbook", transactionInput.CashbookID)
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
	if newTransaction.Type == consts.CashIn {
		cashflowUpdate.TotalIncoming = newTransaction.Amount
		cashflowUpdate.TotalBalance = newTransaction.Amount
	} else {
		cashflowUpdate.TotalOutgoing = newTransaction.Amount
		cashflowUpdate.TotalBalance = -newTransaction.Amount
	}
	conn, close := database.GetDatabaseConnection()
	defer close()
	db := conn.Begin()
	_, err := repo.CreateTransaction(db, newTransaction)
	if err != nil {
		return nil
	}
	cashflow, err := cashflow_repo.CreateOrUpdateCashFlow(db, cashflowUpdate)
	if err != nil {
		return nil
	}
	cashflow.Propagate(db)
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
	// track which cashflow needs to be updated
	cashflowChanges := map[string]*models.CashflowSingle{}
	for _, transaction := range transactions {
		dateV, err := transaction.Date.Value()
		if err != nil {
			resp.FailedCount++
			continue
		}
		dateString := dateV.(time.Time).UTC().GoString()
		// check if the cashflow for that date already exists
		existingCashflow, ok := cashflowChanges[dateString]
		if !ok {
			// if not create a new cashflow
			existingCashflow = &models.CashflowSingle{
				CashbookID:    transaction.CashbookID,
				TotalIncoming: 0.00,
				TotalOutgoing: 0.00,
				TotalBalance:  0.00,
				Date:          transaction.Date,
			}
		}
		if transaction.Type == consts.CashIn {
			// update the cashflow
			existingCashflow.TotalIncoming += transaction.Amount
			existingCashflow.TotalBalance += transaction.Amount
		} else {
			// update the cashflow
			existingCashflow.TotalOutgoing += transaction.Amount
			existingCashflow.TotalBalance -= transaction.Amount
		}
		cashflowChanges[dateString] = existingCashflow
		fmt.Println(*existingCashflow)
		resp.SuccessCount++
	}
	repo := transaction_repo.TransactionRepository{}
	cashflow_repo := cashflow_repo.CashFlowRepository{}
	conn, close := database.GetDatabaseConnection()
	defer close()
	db := conn.Begin()
	_a, err := repo.CreateBulkTransactions(db, transactions)
	_ = _a
	if err != nil {
		resp.FailedCount++
		log.Fatalln(err)
	}
	sortedByDate := []*models.CashflowSingle{}
	for _, cashflow := range cashflowChanges {
		sortedByDate = append(sortedByDate, cashflow)

	}
	sort.Slice(sortedByDate, func(i, j int) bool {
		date1, _ := sortedByDate[i].Date.Value()
		date2, _ := sortedByDate[j].Date.Value()
		return date1.(time.Time).Before(date2.(time.Time))
	})
	for _, cashflow := range sortedByDate {
		_, err := cashflow_repo.CreateOrUpdateCashFlow(db, cashflow)
		if err != nil {
			resp.FailedCount++
			log.Fatalln(err)
		}
	}
	db.Commit()
	return resp

}
func (t *TransactionService) UpdateSingleTransaction() {}
func (t *TransactionService) GetSingleTransaction(id string) *models.TransactionSingle {
	// search for transaction by id
	return nil
}
func (t *TransactionService) ListAllTransactions(filter *models.TransactionFilter) ([]models.TransactionSingle, error) {
	db, close := database.GetDatabaseConnection()
	defer close()
	repo := transaction_repo.TransactionRepository{}
	transactions, err := repo.ListAllTransactions(db, filter)
	return transactions, err
}
