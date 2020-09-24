package Routing

import (
	"TechDojo_http/Response"
	"TechDojo_http/Token"
	"database/sql"
	"encoding/json"
	"net/http"
)

// routing：/user/create
// requests：名前　responses：トークン（jwt）
func CreateUser(db *sql.DB) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){
		// POSTかどうか
		if r.Method != "POST"{
			Response.ErrorBack(w, http.StatusBadRequest, "POST is failed")
			return
		}

		defer r.Body.Close()
		// bodyを読み取る際の受取先構造体
		var response CreateRequest
		// POSTで受け取った情報のBodyを引っ張ってくる
		dec := json.NewDecoder(r.Body)
		if err := dec.Decode(&response); err != nil{
			Response.ErrorBack(w, http.StatusBadRequest, err.Error())
			return
		}

		// sqlに挿入する文
		ins, err := db.Prepare("INSERT INTO userdata(name) VALUES (?)")
		if err != nil{
			Response.ErrorBack(w, http.StatusBadRequest, err.Error())
			return
		}

		// nameデータを保存する（idは自動生成）
		ins.Exec(response.Name)

		// id取得
		var ID string
		if err := db.QueryRow("SELECT id FROM userdata WHERE name = ? LIMIT 1", response.Name).Scan(&ID); err != nil{
			Response.ErrorBack(w, http.StatusBadRequest, err.Error())
			return
		}

		token, err := Token.CreateToken(ID)
		if err != nil{
			Response.ErrorBack(w, http.StatusBadRequest, err.Error())
			return
		}

		Tokenbase := &TokenResponse{
			Token: token,
		}

		Response.SuccessBack(w, Tokenbase)
	}
}

type CreateRequest struct {
	Name string`json:"name"`
}

type TokenResponse struct {
	Token string`json:"token"`
}


