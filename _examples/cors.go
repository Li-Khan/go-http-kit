package main

import (
	goHttpKit "github.com/Li-Khan/go-http-kit"
	"log"
	"net/http"
)

func main() {
	ghk := goHttpKit.New()

	// CORS Configuration
	cors := goHttpKit.NewCORS().
		SetAllowedOrigins("http://example.com").
		SetAllowedMethods("GET", "POST").
		SetAllowedHeaders("Content-Type", "Authorization").
		SetExposeHeaders("Custom-Header")

	ghk.Cors(cors)

	ghk.GET("/", hello)

	err := http.ListenAndServe(":8080", ghk.Mux())
	if err != nil {
		log.Fatal(err)
	}
}
