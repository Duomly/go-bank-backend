package helpers

import (
	"crypto/md5"
	"encoding/hex"

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