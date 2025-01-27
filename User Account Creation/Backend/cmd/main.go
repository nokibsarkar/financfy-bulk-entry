package main

import (
	"log"

	"github.com/adibaruet/financfy-backend/core"
	"github.com/adibaruet/financfy-backend/handler"
)

func main() {
	// Connection to Database Pool
	pool := core.ConnectDB()
	defer func() {
		pool.Close()
		log.Println("Database connection pool closed")
	}()

	// Calling Backend Server
	handler.Server(pool)
}
