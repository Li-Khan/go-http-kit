package main

import (
	goHttpKit "github.com/Li-Khan/go-http-kit"
	"log"
	"net/http"
)

func main() {
	ghk := goHttpKit.New()

	ghk.GET("/", hello)

	err := http.ListenAndServe(":8080", ghk.Mux())
	if err != nil {
		log.Fatal(err)
	}
}

func hello(rw http.ResponseWriter, r *http.Request) {
	_, _ = rw.Write([]byte("Hello, go-http-kit!"))
}
