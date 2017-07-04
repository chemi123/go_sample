package main

/*
とりあえずpostリクエストを受けてbodyをいい感じで処理したいのでその勉強がてらのサンプルプログラム
*/

import (
	"bytes"
	"fmt"
	"net/http"
)

func postHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		fmt.Fprintf(w, "this is not post request\n")
		return
	}
	fmt.Fprintf(w, "hello\n")
	fmt.Fprintf(w, "content length: %v\n", r.ContentLength)

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

func main() {
	http.HandleFunc("/post", postHandler)
	http.ListenAndServe(":8080", nil)
}
