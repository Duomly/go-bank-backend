package useraccounts

import (
	"duomly.com/go-bank-backend/helpers"
	"duomly.com/go-bank-backend/interfaces"
)

// UpdateAccount Create function update account
func UpdateAccount(id uint, amount int) {
	db := helpers.ConnectDB()
	db.Model(&interfaces.Account{}).Where("id = ?", id).Update("balance", amount)
	defer db.Close()
}
