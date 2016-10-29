package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method)
	// リクエストを取得するメソッド
	if r.Method == "GET" {
		t, _ := template.ParseFiles("login.gtpl")
		t.Execute(w, nil)
	} else {
		r.ParseForm()
		// ログインデータがリクエストされ、ログインのロジック判断が実行
		fmt.Println("username:", r.Form["username"])
		fmt.Println("password:", r.Form["password"])
	}
}

func main() {
	http.HandleFunc("/login", login)
	// アクセスのルーティングを設定

	err := http.ListenAndServeTLS(":50124", "ssl/development/myself.crt", "ssl/development/myself.key", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
