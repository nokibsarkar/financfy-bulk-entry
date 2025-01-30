package cashbook_router

import (
	"financify/bulk-entry/models"
	cashbook_service "financify/bulk-entry/services/cashbook"

	"github.com/gin-gonic/gin"
)

func ListAllCashBooks(c *gin.Context) {
	// This function is used to list all the cashbooks of that specific user
	// return
	service := cashbook_service.CashBookService{}
	sampleCashbook := service.GetCashBook()
	response := models.ResponseMultiple[models.CashbookSingle]{
		Data:  []models.CashbookSingle{sampleCashbook},
		Error: "",
	}
	c.JSON(200, response)
}
func CreateCashBook(*gin.Context) {
	// This function is used to create a new cashbook
	// return
}
func UpdateSingleCashBook(*gin.Context) {
	// This function is used to update a cashbook
	// return
}
func DeleteSingleCashBook(*gin.Context) {
	// This function is used to delete a cashbook
	// return
}

func GetSingleCashBook(*gin.Context) {
	// This function is used to get a single cashbook
	// return
}

/**
 * Desc: This function is used to inject the routes for the cashbook endpoint
 */
func Orchestrate(root *gin.RouterGroup) *gin.RouterGroup {

	cashbook_endpoint := root.Group("/cashbook")
	{
		cashbook_endpoint.GET("/", ListAllCashBooks)
		cashbook_endpoint.POST("/", CreateCashBook)
		cashbook_endpoint.PUT("/:id", UpdateSingleCashBook)
		cashbook_endpoint.DELETE("/:id", DeleteSingleCashBook)
		cashbook_endpoint.GET("/:id", GetSingleCashBook)
	}
	return cashbook_endpoint
}
