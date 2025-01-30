package main

import (
	"log"

	"financify/bulk-entry/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	routes.Orchestrate(*r.Group("/"))

	defer log.Println("Exiting...")
	defer func() {
		// recover from panic if one occured. Set err to nil otherwise
		if err := recover(); err != nil {
			log.Println(err)
		}
	}()
	log.Println("Starting the application...")
	r.Run(":8080")
}
