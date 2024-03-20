package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
	"ohlc-price-data/api/db"
	"ohlc-price-data/api/http/controller"
	"ohlc-price-data/api/repository"
	"os"
)

func main() {
	dbConnection, err := db.MakeDBConnectionFrom(fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s",
		os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_HOST"),
		os.Getenv("MYSQL_PORT"),
		os.Getenv("MYSQL_DATABASE_NAME"),
	))

	defer dbConnection.Close()

	if err != nil {
		panic(fmt.Sprintf("There was an issue connecting to the DB: %s", err))
	}

	// --- DI
	tradesRepository, _ := repository.MakeTradesRepositoryFrom(dbConnection)
	tradesController, _ := controller.MakeTradesController(tradesRepository)

	// --- Handlers
	http.HandleFunc("/upload/csv", tradesController.Upload)

	// --- Server
	err = http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("APP_HTTP_SERVING_PORT")), nil)

	if err != nil {
		panic(fmt.Sprintf("There was an issue connecting to the Server: %s", err))
	}
}
