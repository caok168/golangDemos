package main

import (
	"fmt"
	"golangDemos/gormtest/app"
)

func main() {
	dbURL := "host=localhost port=5431 user=test password=test dbname=test sslmode=disable"
	db, err := app.ConnectPostgres(dbURL, false)
	if err != nil {
		fmt.Printf("db init failed: %v\n", err)
	}
	defer db.Close()
}
