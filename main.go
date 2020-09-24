package main

import (
	"TechDojo_http/Routing"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
)

var db *sql.DB

func init(){
	db, _ = sql.Open("mysql", "")
}

func main(){
	// ユーザー系
	http.HandleFunc("/user/create", Routing.CreateUser(db))
	http.HandleFunc("/user/get", Routing.GetUser(db))
	http.HandleFunc("/user/update", Routing.UpdateUser(db))

	// ガチャ系
	http.HandleFunc("/gacha/draw", Routing.GachaDraw(db))

	fmt.Println("server starting....")

	http.ListenAndServe(":8080", nil)
}