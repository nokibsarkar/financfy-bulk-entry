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
	sampleCashbook := service.ListCashbooks()
	response := models.ResponseMultiple[models.CashbookSingle]{
		Data:  sampleCashbook,
		Error: "",
	}
	c.JSON(200, response)
}
func CreateCashBook(c *gin.Context) {
	// This function is used to create a new cashbook
	// return
	newCashbookInp := models.CreateCashBookInput{}
	if err := c.BindJSON(&newCashbookInp); err != nil {
		resp := models.ResponseSingle[models.CashbookSingle]{Data: nil, Error: err.Error()}
		c.JSON(400, resp)
		return
	}
	service := cashbook_service.CashBookService{}
	newCashbook, err := service.CreateCashbook(newCashbookInp)
	if err != nil {
		response := models.ResponseSingle[models.CashbookSingle]{Data: nil, Error: err.Error()}
		c.JSON(400, response)
		return
	}
	response := models.ResponseSingle[models.CashbookSingle]{Data: newCashbook, Error: ""}
	c.JSON(200, response)
}
func UpdateSingleCashBook(*gin.Context) {
	// This function is used to update a cashbook
	// return
}
func DeleteSingleCashBook(*gin.Context) {
	// This function is used to delete a cashbook
	// return
}

func GetSingleCashBook(c *gin.Context) {
	// This function is used to get a single cashbook
	cashbookId := c.Param("id")
	// Convert the string to string
	// If the conversion fails, return an error

	// Get the cashbook from the service
	service := cashbook_service.CashBookService{}
	sampleCashbook := service.GetSingleCashBook(cashbookId)
	if sampleCashbook == nil {
		response := models.ResponseSingle[models.CashbookSingle]{Data: nil, Error: "Cashbook not found"}
		c.JSON(404, response)
		return
	}
	response := models.ResponseSingle[models.CashbookSingle]{Data: sampleCashbook, Error: ""}
	c.JSON(200, response)
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
