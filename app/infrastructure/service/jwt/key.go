package jwt

import (
	"os"
)

var (
	jwtKey = []byte(os.Getenv("TOKEN_PASSWORD"))
)

func GetJwtKey() []byte {
	return jwtKey
}
