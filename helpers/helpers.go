package helpers

import (
	"crypto/md5"
	"encoding/hex"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)


func HandleErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func HashAndSalt(pass []byte) string {
	hashed, err := bcrypt.GenerateFromPassword(pass, bcrypt.MinCost)
	HandleErr(err)

	return string(hashed)
}

func HashOnlyVulnerable(pass []byte) string {
	hash := md5.New()
  hash.Write(pass)
  return hex.EncodeToString(hash.Sum(nil))
}

func ConnectDB() *gorm.DB {
	db, err := gorm.Open("postgres", "host=127.0.0.1 port=5432 user=postgres dbname=bankapp password=postgres sslmode=disable")
	HandleErr(err)
	return db
}