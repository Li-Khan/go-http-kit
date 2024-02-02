package main

import (
	goHttpKit "github.com/Li-Khan/go-http-kit"
	"log"
	"net/http"
)

func main() {
	ghk := goHttpKit.New()

	// Middleware functions
	middleware1 := func(next http.HandlerFunc) http.HandlerFunc {
		return func(rw http.ResponseWriter, r *http.Request) {
			log.Println("Executing middleware 1")
			next.ServeHTTP(rw, r)
		}
	}

	middleware2 := func(next http.HandlerFunc) http.HandlerFunc {
		return func(rw http.ResponseWriter, r *http.Request) {
			log.Println("Executing middleware 2")
			next.ServeHTTP(rw, r)
		}
	}

	// Apply middleware globally
	ghk.Middleware(middleware1, middleware2)

	ghk.GET("/", func(rw http.ResponseWriter, r *http.Request) {
		_, _ = rw.Write([]byte("Hello, Middleware!"))
	})

	err := http.ListenAndServe(":8080", ghk.Mux())
	if err != nil {
		log.Fatal(err)
	}
}
