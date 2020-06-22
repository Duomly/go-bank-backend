package transactions

import (
	"duomly.com/go-bank-backend/helpers"
	"duomly.com/go-bank-backend/interfaces"
)
// Add function create transaction
func CreateTransaction(From uint, To uint, Amount int) {
	db := helpers.ConnectDB()
	transaction := &interfaces.Transaction{From: From, To: To, Amount: Amount}
	db.Create(&transaction)

	defer db.Close()
}