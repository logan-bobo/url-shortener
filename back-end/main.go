package main

import (
	"fmt"
	"os"

	"github.com/logan-bobo/url_shortener/controllers"
	"github.com/logan-bobo/url_shortener/db"
	"github.com/logan-bobo/url_shortener/models"
	"github.com/logan-bobo/url_shortener/server"
)

type DBconfig struct {
	Host     string
	User     string
	Password string
	Name     string
	Port     string
}

func main() {
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")

	dbConfig := &DBconfig{
		Host:     dbHost,
		User:     dbUser,
		Password: dbPassword,
		Name:     dbName,
		Port:     dbPort,
	}

	connectionString := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s",
		dbConfig.Host,
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Name,
		dbConfig.Port,
	)

	con := db.ConnectDB(connectionString)

	err := con.AutoMigrate(&models.SavedURL{})
	if err != nil {
		panic(err)
	}

	bh := controllers.NewBaseHandler(con)

	server.StartServer(bh)
}
