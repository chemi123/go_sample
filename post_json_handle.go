package main

// 公式の提供する未知のjsonのパース+post requestの処理
// bitly社のsimplejsonというものもある
// https://github.com/bitly/go-simplejson

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func jsonHandle(w http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		fmt.Fprintf(w, "Only takes post request\n")
		return
	}

	reqBody, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Fprintf(w, "Failed to read body\n")
		return
	}

	// debug
	fmt.Println("reqBody:", string(reqBody))

	var i interface{}
	err = json.Unmarshal(reqBody, &i)
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}

	bodyMap := i.(map[string]interface{})
	for key, value := range bodyMap {
		switch v := value.(type) {
		case string:
			fmt.Println(key, "has string value", v)
		case int:
			fmt.Println(key, "has int value", v)
		case float64:
			fmt.Println(key, "has float64 value", v)
		case []interface{}:
			fmt.Println(key, "has an array value", v)
			for i, u := range v {
				fmt.Println(i, u)
			}
		default:
			fmt.Println(key, "is of a type I don't know")
		}
	}
}

func main() {
	http.HandleFunc("/", jsonHandle)
	http.ListenAndServe(":8080", nil)
}
