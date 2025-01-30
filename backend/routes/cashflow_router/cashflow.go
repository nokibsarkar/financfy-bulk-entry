package cashflow_router

import (
	"financify/bulk-entry/models"
	"fmt"

	"github.com/gin-gonic/gin"
)

func ListAllCashFlows(c *gin.Context) {
	// This function is used to list all the CashFlows of that specific user
	// return
	sampleCashFlow := models.CashflowSingle{
		ID:            1,
		TotalIncoming: 1000.00,
		TotalOutgoing: 500.00,
		TotalBalance:  500.00,
	}
	fmt.Println(sampleCashFlow)
	response := models.ResponseMultiple[models.CashflowSingle]{
		Data:  []models.CashflowSingle{sampleCashFlow},
		Error: "",
	}
	c.JSON(200, response)
}

func GetSingleCashFlow(*gin.Context) {
	// This function is used to get a single CashFlow
	// return
}

/**
 * Desc: This function is used to inject the routes for the CashFlow endpoint
 */
func Orchestrate(root *gin.RouterGroup) *gin.RouterGroup {

	CashFlow_endpoint := root.Group("/cashflow")
	{
		CashFlow_endpoint.GET("/", ListAllCashFlows)
		CashFlow_endpoint.GET("/:id", GetSingleCashFlow)
	}
	return CashFlow_endpoint
}
