package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type Server struct {
	ServerName string
	ServerIP   string
}

type ServerSlice struct {
	Servers []Server
}

func main() {
	var s ServerSlice
	str := `{"servers":[{"serverName":"Shanghai_VPN","serverIP":"127.0.0.1"},{"serverName":"Beijing_VPN","serverIP":"127.0.0.2"}]}`
	json.Unmarshal([]byte(str), &s)
	//fmt.Println(string([]byte(str))) // str全体(json)
	//fmt.Println(s)                   // strのjson解析部分

	var f interface{}
	b := []byte(`{"Name":"Wednesday","Age":6,"Parents":["Gomez","Morticia"]}`)
	fmt.Println("b", reflect.TypeOf(b))
	err := json.Unmarshal(b, &f)
	if err != nil {
		fmt.Println(err)
		return
	}

	m := f.(map[string]interface{})
	fmt.Println(m["Name"], ":", reflect.TypeOf(m["Name"]))
	fmt.Println(m["Age"], ":", reflect.TypeOf(m["Age"]))
	fmt.Println(m["Parents"], ":", reflect.TypeOf(m["Parents"]))
	fmt.Println()

	for k, v := range m {
		switch vv := v.(type) {
		case string:
			fmt.Println(k, "is string", vv)
		case int:
			fmt.Println(k, "is int", vv)
		case float64:
			fmt.Println(k, "is float64", vv)
		case []interface{}:
			fmt.Println(k, "is an array", vv)
			for i, u := range vv {
				fmt.Println(i, u)
			}
		default:
			fmt.Println(k, "is of a type I don't know how to handle")
		}
	}
}
