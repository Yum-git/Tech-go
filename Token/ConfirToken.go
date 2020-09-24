package Token

import (
	"github.com/dgrijalva/jwt-go"
)

func ConfirToken(tokenString string) (*jwt.Token, error){
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error){
		return []byte("secret"), nil
	})

	if err != nil{
		return token, err
	}
	return token, nil
}
