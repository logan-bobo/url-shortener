// The config package holds run time configurations such as DB connection parameters
package config

import (
	"os"
)

type DBconfig struct {
	Host     string
	User     string
	Password string
	Name     string
	Port     string
}

var dbConfig *DBconfig

func Init() {
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")

	dbConfig = &DBconfig{
		Host:     dbHost,
		User:     dbUser,
		Password: dbPassword,
		Name:     dbName,
		Port:     dbPort,
	}
}

func GetDBConfig() *DBconfig {
	return dbConfig
}
