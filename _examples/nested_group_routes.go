package main

import (
	goHttpKit "github.com/Li-Khan/go-http-kit"
	"log"
	"net/http"
)

func main() {
	ghk := goHttpKit.New()

	// Global middleware for all routes
	globalMiddleware := func(next http.HandlerFunc) http.HandlerFunc {
		return func(rw http.ResponseWriter, r *http.Request) {
			log.Println("Executing global middleware")
			next.ServeHTTP(rw, r)
		}
	}
	ghk.Middleware(globalMiddleware)

	// Creating the main route group
	mainGroup := ghk.Group("/main")

	// Middleware specific to the main route group
	mainGroupMiddleware := func(next http.HandlerFunc) http.HandlerFunc {
		return func(rw http.ResponseWriter, r *http.Request) {
			log.Println("Executing middleware for the main group")
			next.ServeHTTP(rw, r)
		}
	}
	mainGroup.Middleware(mainGroupMiddleware)

	// Creating a nested route group inside the main group
	subGroup := mainGroup.Group("/sub")

	// Middleware specific to the nested route group
	subGroupMiddleware := func(next http.HandlerFunc) http.HandlerFunc {
		return func(rw http.ResponseWriter, r *http.Request) {
			log.Println("Executing middleware for the sub group")
			next.ServeHTTP(rw, r)
		}
	}
	subGroup.Middleware(subGroupMiddleware)

	// Routes within the nested group
	subGroup.GET("/route1", func(rw http.ResponseWriter, r *http.Request) {
		_, _ = rw.Write([]byte("Sub Group - Route 1"))
	})

	subGroup.GET("/route2", func(rw http.ResponseWriter, r *http.Request) {
		_, _ = rw.Write([]byte("Sub Group - Route 2"))
	})

	// Another route outside both groups
	ghk.GET("/home", func(rw http.ResponseWriter, r *http.Request) {
		_, _ = rw.Write([]byte("Home Page"))
	})

	err := http.ListenAndServe(":8080", ghk.Mux())
	if err != nil {
		log.Fatal(err)
	}
}
