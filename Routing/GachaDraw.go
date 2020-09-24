package Routing

import (
	"TechDojo_http/Response"
	"TechDojo_http/Token"
	"database/sql"
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"net/http"
)

func GachaDraw(db *sql.DB) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST"{
			return
		}

		tokenstring := r.Header.Get("x-token")

		token, err := Token.ConfirToken(tokenstring)
		if err != nil{
			Response.ErrorBack(w, http.StatusBadRequest, err.Error())
			return
		}

		decodetoken := token.Claims.(jwt.MapClaims)

		defer r.Body.Close()
		var response GachaTimes
		dec := json.NewDecoder(r.Body)
		if err := dec.Decode(&response); err != nil{
			Response.ErrorBack(w, http.StatusBadRequest, err.Error())
			return
		}
		if response.Times != 0{
			Response.ErrorBack(w, http.StatusBadRequest, "times values fails data")
			return
		}

		rows, err := db.Query("SELECT * FROM gachatable")
		if err != nil{
			Response.ErrorBack(w, http.StatusBadRequest, err.Error())
			return
		}
	}
}

type GachaTimes struct {
	Times int`json:"times"`
}

type GachaTables struct {
	Rarity string`json:"rarity"`
	Probability string`json:"probability"`
}