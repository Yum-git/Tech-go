package Response

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func SuccessBack(w http.ResponseWriter, data interface{}){
	jsondata, _ := json.Marshal(data)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsondata)
}

func ErrorBack(w http.ResponseWriter, code int, message string){
	w.Header().Set("Content-Type", "application/json")
	jsondata, err := json.Marshal(&ErrorResponse{
		Code: code,
		Message: message,
	})
	if err != nil{
		fmt.Println(err)
	}

	w.WriteHeader(code)
	w.Write(jsondata)
}

type ErrorResponse struct {
	Code int`json:"code"`
	Message string`json:"message"`
}

