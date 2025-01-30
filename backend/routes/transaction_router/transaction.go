package transaction_router

import (
	"financify/bulk-entry/models"

	"github.com/gin-gonic/gin"
)

func ListAllTransactions(c *gin.Context) {
	// This function is used to list all the Transactions of that specific user
	// return
	sampleTransaction := models.TransactionSingle{
		ID:         1,
		VoucherNo:  "1234",
		Date:       "2021-01-01",
		Amount:     1000.00,
		Contact:    "John Doe",
		CashbookID: 1,
		Reference:  "Ref-1234",
		Remarks:    "Sample Transaction",
	}
	response := models.ResponseMultiple[models.TransactionSingle]{
		Data:  []models.TransactionSingle{sampleTransaction},
		Error: "",
	}
	c.JSON(200, response)
}
func CreateTransaction(*gin.Context) {
	// This function is used to create a new Transaction
	// return
}
func CreateBulkTransactions(*gin.Context) {
	// This function is used to create a new Transaction in bulk
}
func UpdateSingleTransaction(*gin.Context) {
	// This function is used to update a Transaction
	// return
}

// func DeleteSingleTransaction(*gin.Context) {
// 	// This function is used to delete a Transaction
// 	// return
// }

func GetSingleTransaction(c *gin.Context) {
	// This function is used to get a single Transaction
	// return
	sampleTransaction := models.TransactionSingle{
		ID:         1,
		VoucherNo:  "1234",
		Date:       "2021-01-01",
		Amount:     1000.00,
		Contact:    "John Doe",
		CashbookID: 1,
		Reference:  "Ref-1234",
		Remarks:    "Sample Transaction",
	}
	response := models.ResponseSingle[models.TransactionSingle]{Data: sampleTransaction, Error: ""}
	c.JSON(200, response)
}

/**
 * Desc: This function is used to inject the routes for the Transaction endpoint
 */
func Orchestrate(root *gin.RouterGroup) *gin.RouterGroup {

	Transaction_endpoint := root.Group("/Transaction")
	{
		Transaction_endpoint.GET("/", ListAllTransactions)
		Transaction_endpoint.POST("/", CreateTransaction)
		Transaction_endpoint.POST("/bulk", CreateBulkTransactions)
		Transaction_endpoint.PUT("/:id", UpdateSingleTransaction)
		// Transaction_endpoint.DELETE("/:id", DeleteSingleTransaction)
		Transaction_endpoint.GET("/:id", GetSingleTransaction)
	}
	return Transaction_endpoint
}
