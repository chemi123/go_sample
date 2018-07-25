package main

/*
* 9000ポートをlistenして8081ポートにリクエウストをforwardするサンプル
*/

import (
	"log"
	"net/http"
	"net/http/httputil"
)

func main() {
	director := func(request *http.Request) {
		request.URL.Scheme = "http"
		request.URL.Host = "localhost:8081"
		// request.URL.Host = ":8081" <- 上と同じ
	}
	rp := &httputil.ReverseProxy{Director: director}
	server := http.Server{
		Addr:    ":9000",
		Handler: rp,
	}
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err.Error())
	}
}
