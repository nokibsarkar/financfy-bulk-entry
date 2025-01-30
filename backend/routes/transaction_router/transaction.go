package transaction_router

import (
	"financify/bulk-entry/models"
	transaction_service "financify/bulk-entry/services/cashbook/cashflow/transaction"

	"github.com/gin-gonic/gin"
)

func ListAllTransactions(c *gin.Context) {
	// This function is used to list all the Transactions of that specific user
	// return
	service := transaction_service.TransactionService{}
	transactions, err := service.ListAllTransactions()
	if err != nil {
		response := models.ResponseMultiple[models.TransactionSingle]{Data: nil, Error: err.Error()}
		c.JSON(400, response)
		return
	}
	response := models.ResponseMultiple[models.TransactionSingle]{
		Data:  transactions,
		Error: "",
	}
	c.JSON(200, response)
}
func CreateTransaction(c *gin.Context) {
	// This function is used to create a new Transaction
	// return
	newTransactionInp := &models.TransactionSingleInput{}
	if err := c.BindJSON(newTransactionInp); err != nil {
		resp := models.ResponseSingle[models.TransactionSingle]{Data: nil, Error: err.Error()}
		c.JSON(400, resp)
		return
	}
	service := transaction_service.TransactionService{}
	newTransaction := service.CreateTransaction(newTransactionInp)
	response := models.ResponseSingle[models.TransactionSingle]{Data: newTransaction, Error: ""}
	c.JSON(200, response)
}
func CreateBulkTransactions(c *gin.Context) {
	// This function is used to create a new Transaction in bulk
	newTransactions := []models.TransactionSingleInput{}
	if err := c.BindJSON(&newTransactions); err != nil {
		resp := models.ResponseMultiple[models.TransactionSingle]{Data: nil, Error: err.Error()}
		c.JSON(400, resp)
		return
	}
	service := transaction_service.TransactionService{}
	resp := service.CreateBulkTransactions(newTransactions)
	c.JSON(200, resp)
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
	// Convert it to string
	// If the conversion fails, return an error
	trxID := c.Param("id")
	service := transaction_service.TransactionService{}
	sampleTransaction := service.GetSingleTransaction(trxID)
	response := models.ResponseSingle[models.TransactionSingle]{Data: sampleTransaction, Error: ""}
	c.JSON(200, response)
}

/**
 * Desc: This function is used to inject the routes for the Transaction endpoint
 */
func Orchestrate(root *gin.RouterGroup) *gin.RouterGroup {

	Transaction_endpoint := root.Group("/transaction")
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
