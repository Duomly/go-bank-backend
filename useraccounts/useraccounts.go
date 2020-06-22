package useraccounts

import (
	"fmt"

	"duomly.com/go-bank-backend/helpers"
	"duomly.com/go-bank-backend/interfaces"
	"duomly.com/go-bank-backend/transactions"
)

// Update function updateAccount
func updateAccount(id uint, amount int) interfaces.ResponseAccount {
	db := helpers.ConnectDB()
	account := interfaces.Account{}
	responseAcc := interfaces.ResponseAccount{}

	db.Where("id = ? ", id).First(&account)
	account.Balance = uint(amount)
	db.Save(&account)
	
	responseAcc.ID = account.ID
	responseAcc.Name = account.Name
	responseAcc.Balance = int(account.Balance)
	defer db.Close()
	return responseAcc
}

// Create function getAccount
func getAccount(id uint) *interfaces.Account{
	db := helpers.ConnectDB()
	account := &interfaces.Account{}
	if db.Where("id = ? ", id).First(&account).RecordNotFound() {
		return nil
	}
	defer db.Close()
	return account
}

// Create function Transaction
func Transaction(userId uint, from uint, to uint, amount int, jwt string) map[string]interface{} {
	// Conver uint to string
	userIdString := fmt.Sprint(userId)
	// Validate ownership
	isValid := helpers.ValidateToken(userIdString, jwt)
	if isValid {
		// Take sender and receiver
		fromAccount := getAccount(from)
		toAccount := getAccount(to)
		// Handle errors
		if fromAccount == nil || toAccount == nil {
			return map[string]interface{}{"message": "Account not found"}
		} else if fromAccount.UserID != userId {
			return map[string]interface{}{"message": "You are not owner of the account"}
		} else if int(fromAccount.Balance) < amount {
			return map[string]interface{}{"message": "Account balance is too small"}
		}
		// Update account
		updatedAccount := updateAccount(from, int(fromAccount.Balance) - amount)
		updateAccount(to, int(fromAccount.Balance) + amount)

		// Create transaction
		transactions.CreateTransaction(from, to, amount)

		// Return response
		var response = map[string]interface{}{"message": "all is fine"}
		response["data"] = updatedAccount
		return response
	} else {
		return map[string]interface{}{"message": "Not valid token"}
	}
}