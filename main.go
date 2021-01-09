package main

import (
	"duomly.com/go-bank-backend/api"
	"duomly.com/go-bank-backend/database"
)

func main() {
	// Init database
	database.InitDatabase()
	// Do migration
	// migrations.Migrate()
	api.StartApi()
}
