package cashflow_router

import (
	"financify/bulk-entry/models"
	cashflow_service "financify/bulk-entry/services/cashbook/cashflow"

	"github.com/gin-gonic/gin"
)

func ListAllCashFlows(c *gin.Context) {
	// This function is used to list all the CashFlows of that specific user
	// return
	service := cashflow_service.CashFlowService{CashbookID: 2147006226616025000}
	sampleCashFlow := service.ListCashFlowByCashbookID()
	response := models.ResponseMultiple[models.CashflowSingle]{
		Data:  sampleCashFlow,
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
