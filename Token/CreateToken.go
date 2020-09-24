package Token

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

// Tokenを生成する
// idを渡せば返ってくる
func CreateToken(id string) (string, error){
	token := jwt.New(jwt.GetSigningMethod("HS256"))

	token.Claims = jwt.MapClaims{
		"id": id,
		"exp": time.Now().Add(time.Hour * 1).Unix(),
	}

	var secretKey = "secret"
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil{
		return "", err
	}
	return tokenString, nil
}
