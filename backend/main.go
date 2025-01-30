package main

import (
	"financify/bulk-entry/consts"
	"financify/bulk-entry/routes"
	"log"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	log.Println("Starting the application...")
	log.Println("Setting up the configuration...")
	currentDir, _ := os.Getwd()
	consts.LoadConfig(currentDir + "/config.yaml")
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
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
