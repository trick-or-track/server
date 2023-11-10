package utils

import (
	"time"

	"github.com/golang-jwt/jwt"
)

// TODO move go env
var JWTSecret = []byte("!!SECRET!!")

func GenerateJWT(id int) string {
	token := jwt.New((jwt.SigningMethodHS256))
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = id
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	signedToken, _ := token.SignedString(JWTSecret)
	return signedToken
}
