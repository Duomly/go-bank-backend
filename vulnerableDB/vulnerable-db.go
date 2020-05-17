package vulnerableDB

import (
	"database/sql"
	"fmt"

	"duomly.com/go-bank-backend/helpers"
	_ "github.com/lib/pq"
)

type User struct {
	ID int
	Username string
	Email string
	Accounts []Account
}

type Account struct {
	ID int
	Name string
	Balance int
}

func connectDB() *sql.DB {
	db, err := sql.Open("postgres", "host=127.0.0.1 port=5432 user=user dbname=dbname password=password sslmode=disable")
	helpers.HandleErr(err)
	return db
}

func dbCall(query string) *sql.Rows {
	db := connectDB()

	call, err := db.Query(query)

	helpers.HandleErr(err)
	return call
}

// func correctAuthUser(accPass string, typedPass string) bool {
// 	accPassByte := []byte(accPass)
// 	typedPassByte := []byte(typedPass)
// 	err := bcrypt.CompareHashAndPassword(accPassByte, typedPassByte)
// 	if err != nil {
// 			return false
// 	}
//   return true
// }

func VulnerableLogin(username string, pass string) []User {
	// This is a very vulnerable way of doing it, you should always hash and salt, you should compare them with CompareHashAndPassword as well
	password := helpers.HashOnlyVulnerable([]byte(pass))
	results := dbCall("SELECT id, username, email FROM users x WHERE username='" + username + "' AND password='" + password + "'")
	var users []User

	for results.Next() {
		var user User
		err := results.Scan(&user.ID, &user.Username, &user.Email)
		helpers.HandleErr(err)
		accounts := dbCall("SELECT id, name, balance FROM accounts x WHERE user_id=" + fmt.Sprint(user.ID) + "")
		var userAccounts []Account

		for accounts.Next() {
			var account Account
			err := accounts.Scan(&account.ID, &account.Name, &account.Balance)
			helpers.HandleErr(err)
			userAccounts = append(userAccounts, account)
		}

		user.Accounts = userAccounts
		users = append(users, user)
	}
	defer results.Close()
	return users
}