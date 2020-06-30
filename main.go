package main

import (
	"duomly.com/go-bank-backend/api"
	"duomly.com/go-bank-backend/database"
)



func main() {
	// Do migration
	// migrations.MigrateTransactions()
	
	// Init database
	database.InitDatabase()
	api.StartApi()
}
