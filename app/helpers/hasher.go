package helpers

import (
	"crypto/sha512"
	"encoding/base64"
)

const (
	passwordSalt = "!@#$%^&*()"
)

// HashPassword scrambles a given password with a given email
// and a passwordSalt using sha512 hashing algorithm
func HashPassword(email, password string) string {

	hash := sha512.New()
	hash.Write([]byte(passwordSalt))
	hash.Write([]byte(email))
	hash.Write([]byte(password))

	return base64.URLEncoding.EncodeToString(hash.Sum(nil))
}
