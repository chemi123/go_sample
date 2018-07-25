package main

/*
* 8081をlistenしてGETかPOSTのみを受け付けるサーバのサンプル
* GETであれば適当に文字列を返し、POSTであればRequest Bodyのechoサーバとなる
*/

import (
        "fmt"
	"io/ioutil"
        "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		responseToGETRequest(w, r)
	} else if r.Method == "POST" {
		responseToPOSTRequest(w, r)
	} else {
		fmt.Fprintf(w, "Method should be GET or POST");
	}

	return
}

func responseToGETRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hoge")
}

func responseToPOSTRequest(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Something is wrong")
	}

	fmt.Fprint(w, "content: %v\n", string(body))
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8081", nil)
}
