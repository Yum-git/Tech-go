package Routing

import (
	"TechDojo_http/Token"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"net/http"
)

// routing：/user/update
// requests：トークン（jwt)，名前　responses：なし
func UpdateUser(db *sql.DB) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){
		if r.Method != "PUT"{
			return
		}

		tokenstring := r.Header.Get("x-token")

		token, err := Token.ConfirToken(tokenstring)
		if err != nil{
			fmt.Println(err)
			return
		}

		decodetoken := token.Claims.(jwt.MapClaims)

		defer r.Body.Close()
		// bodyを読み取る際の受取先構造体
		var response UpdateRequest
		// POSTで受け取った情報のBodyを引っ張ってくる
		dec := json.NewDecoder(r.Body)
		if err := dec.Decode(&response); err != nil{
			fmt.Println(err)
		}

		apt, err := db.Prepare("UPDATE userdata set name = ? WHERE id = ?")
		if err != nil{
			fmt.Println(err)
			return
		}
		apt.Exec(response.Name, decodetoken["id"])
	}
}

type UpdateRequest struct {
	Name string`json:"name"`
}
