package Routing

import (
	"TechDojo_http/Response"
	"TechDojo_http/Token"
	"database/sql"
	"github.com/dgrijalva/jwt-go"
	"net/http"
)

// routing：/user/get
// requests：トークン（jwt)　responses：名前
func GetUser(db *sql.DB) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){
		if r.Method != "GET"{
			return
		}

		tokenstring := r.Header.Get("x-token")

		token, err := Token.ConfirToken(tokenstring)
		if err != nil{
			Response.ErrorBack(w, http.StatusBadRequest, err.Error())
			return
		}

		decodetoken := token.Claims.(jwt.MapClaims)
		var Name string

		if err := db.QueryRow("SELECT name FROM userdata WHERE id = ? LIMIT 1", decodetoken["id"]).Scan(&Name); err != nil{
			Response.ErrorBack(w, http.StatusBadRequest, err.Error())
			return
		}

		Namebase := &UserResponse{
			Name: Name,
		}

		Response.SuccessBack(w, Namebase)
	}
}

type TokenRequest struct {
	Token string`json:"token"`
}

type UserResponse struct {
	Name string`json:"name"`
}