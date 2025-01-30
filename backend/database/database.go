// This package would contain the database connection and the database schema.
// The database connection would be created using the configuration object created in the configuration package.
// The database schema would be created using the database connection.
// The database schema would be used to create the tables in the database.
package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func GetDatabaseConnection() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return db
}
func init() {
	//
	db := GetDatabaseConnection()
	// fmt.Println("Database connection created", db)
	// Migrate the schema
	db.AutoMigrate(&Cashbook{})
}
