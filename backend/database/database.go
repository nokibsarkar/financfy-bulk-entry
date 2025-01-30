// This package would contain the database connection and the database schema.
// The database connection would be created using the configuration object created in the configuration package.
// The database schema would be created using the database connection.
// The database schema would be used to create the tables in the database.
package database

import (
	"financify/bulk-entry/consts"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var conf = consts.LoadConfig(".")
var pgst = postgres.Open(conf.DatabaseURL)

func GetDatabaseConnection() (*gorm.DB, func()) {
	db, err := gorm.Open(pgst, &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	raw_db, err := db.DB()
	if err != nil {
		panic("failed to get database connection")
	}
	return db, func() {
		raw_db.Close()
	}
}
func init() {
	//
	log.Println("Creating database schema")
	db, close := GetDatabaseConnection()
	defer close()
	db.AutoMigrate(&Cashbook{})
	db.AutoMigrate(&Transaction{})
	db.AutoMigrate(&CashFlow{})
	log.Println("Database schema created")

}
