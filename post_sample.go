package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func main() {
	//POSTする
	res, err := http.Post(
		"http://localhost:8080",                                          //URL
		"application/json",                                               //Header["Content-Type"]
		strings.NewReader(`{"ID":0,"PlayersNum":2,"IsTDRequest":false}`), //Body
	)
	defer res.Body.Close()

	//エラーハンドリング
	if err != nil {
		log.Fatal(err) //プログラムの停止(停止しないのならlog.Print(err))
	}
	//レスポンスのボディの受け取り表示する
	b, _ := ioutil.ReadAll(res.Body) //エラーハンドリングは省略
	fmt.Println(string(b))
}
