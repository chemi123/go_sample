package main

/*
とりあえずpostリクエストを受けてbodyをいい感じで処理したいのでその勉強がてらのサンプルプログラム
*/

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

// bodyの読み方その1
func postHandler1(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		fmt.Fprintf(w, "this is not post request\n")
		return
	}

	// bufをbytes.Bufferでnewする
	// newとは特定の型を0埋めして、その型のポインタを返す
	buf := new(bytes.Buffer)

	// Bodyをbufに読み込む
	if _, err := buf.ReadFrom(r.Body); err != nil {
		fmt.Fprintf(w, "failed to read body\n")
		return
	}

	req := buf.Bytes()
	fmt.Fprintf(w, "content: %v\n", string(req))
}

// bodyの読み方その2
func postHandler2(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		fmt.Fprintf(w, "this is not post request\n")
		return
	}

	// bodyはbyte型のスライス
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		// 良くない
		panic(err)
	}

	fmt.Fprintf(w, "content: %v\n", string(body))
}

func main() {
	http.HandleFunc("/post1", postHandler1)
	http.HandleFunc("/post2", postHandler2)
	http.ListenAndServe(":8080", nil)
}
