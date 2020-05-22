package users

import (
	"time"

	"duomly.com/go-bank-backend/helpers"
	"duomly.com/go-bank-backend/interfaces"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

func Login(username string, pass string) map[string]interface{} {
	// Connect DB
	db := helpers.ConnectDB()
	user := &interfaces.User{}
	if db.Where("username = ? ", username).First(&user).RecordNotFound() {
		return map[string]interface{}{"message": "User not found"}
	}
	// Verify password
	passErr := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(pass))

	if passErr == bcrypt.ErrMismatchedHashAndPassword && passErr != nil {
		return map[string]interface{}{"message": "Wrong password"}
	}
	// Find accounts for the user
	accounts := []interfaces.ResponseAccount{}
	db.Table("accounts").Select("id, name, balance").Where("user_id = ? ", user.ID).Scan(&accounts)

	// Setup response
	responseUser := &interfaces.ResponseUser{
		ID: user.ID,
		Username: user.Username,
		Email: user.Email,
		Accounts: accounts,
	}

	defer db.Close()

	// Sign token
	tokenContent := jwt.MapClaims{
		"user_id": user.ID,
		"expiry": time.Now().Add(time.Minute * 60).Unix(),
	}
	jwtToken := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tokenContent)
	token, err := jwtToken.SignedString([]byte("TokenPassword"))
	helpers.HandleErr(err)

	// Prepare response
	var response = map[string]interface{}{"message": "all is fine"}
	response["jwt"] = token
	response["data"] = responseUser

	return response
}