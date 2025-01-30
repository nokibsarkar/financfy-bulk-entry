// This package would contain the database connection and the database schema.
// The database connection would be created using the configuration object created in the configuration package.
// The database schema would be created using the database connection.
// The database schema would be used to create the tables in the database.
package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var connString = "host=localhost user=postgres password=KothinKichu dbname=financify port=5432 sslmode=disable"

func GetDatabaseConnection() *gorm.DB {
	pgst := postgres.Open(connString)
	db, err := gorm.Open(pgst, &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return db
}
func init() {
	//
	log.Println("Creating database schema")
	db := GetDatabaseConnection()
	db.AutoMigrate(&Cashbook{})
	db.AutoMigrate(&Transaction{})
	db.AutoMigrate(&CashFlow{})
	log.Println("Database schema created")
}
