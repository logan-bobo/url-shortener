package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// ConnectDB opens a connection to the database
func ConnectDB(connectionString string) *gorm.DB {
	con, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	return con
}
