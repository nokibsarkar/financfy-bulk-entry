// Desc: This file is used to define all the routes for the application
package routes

import (
	"financify/bulk-entry/routes/cashbook_router"
	"financify/bulk-entry/routes/cashflow_router"
	"financify/bulk-entry/routes/transaction_router"

	"github.com/gin-gonic/gin"
)

func ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func Orchestrate(root gin.RouterGroup) *gin.RouterGroup {
	api_endpoint := root.Group("/api/v1")
	{
		api_endpoint.GET("/ping", ping)
		// Inject the cashbook routes
		cashbook_router.Orchestrate(api_endpoint)
		transaction_router.Orchestrate(api_endpoint)
		cashflow_router.Orchestrate(api_endpoint)
	}
	return api_endpoint
}
