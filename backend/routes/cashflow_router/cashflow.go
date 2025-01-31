package cashflow_router

import (
	"financify/bulk-entry/models"
	cashflow_service "financify/bulk-entry/services/cashbook/cashflow"
	"fmt"

	"github.com/gin-gonic/gin"
)

func ListAllCashFlows(c *gin.Context) {
	// This function is used to list all the CashFlows of that specific user
	// return
	service := cashflow_service.CashFlowService{}
	CashBookID := c.Query("cashbook_id")
	if CashBookID == "" {
		CashBookID = "1rj4m1je40000"
		fmt.Println(CashBookID)
	}
	cashflows := service.ListCashFlowByCashbookID(CashBookID)
	response := models.ResponseMultiple[models.CashflowSingle]{
		Data:  cashflows,
		Error: "",
	}
	c.JSON(200, response)
}

func GetSingleCashFlow(c *gin.Context) {
	// This function is used to get a single CashFlow
	// return
	CashflowID, ok := c.Params.Get("id")
	fmt.Println(CashflowID)
	if !ok {
		response := models.ResponseSingle[models.CashflowSingle]{Data: nil, Error: "Cashflow not found"}
		c.JSON(404, response)
		return
	}
	service := cashflow_service.CashFlowService{}
	singleCashflow := service.GetSingleCashFlowByID(CashflowID)
	if singleCashflow == nil {
		response := models.ResponseSingle[models.CashflowSingle]{Data: nil, Error: "Cashflow not found"}
		c.JSON(404, response)
		return
	}
	response := models.ResponseSingle[models.CashflowSingle]{
		Data:  singleCashflow,
		Error: "",
	}
	c.JSON(200, response)

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
