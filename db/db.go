// The db package uses our configuration to instantiate a database connection
package db

import (
	"fmt"

	"github.com/logan-bobo/url_shortener/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func connectDB(connectionString string) {
	con, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db = con
}

func Init() {
	config.Init()

	dbConfig := config.GetDBConfig()

	// Build connection string
	connectionString := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s",
		dbConfig.Host,
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Name,
		dbConfig.Port,
	)

	connectDB(connectionString)
}

func GetDB() *gorm.DB {
	return db
}
