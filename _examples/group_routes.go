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

	// Creating a route group
	apiGroup := ghk.Group("/api")

	// Middleware specific to the route group
	groupMiddleware := func(next http.HandlerFunc) http.HandlerFunc {
		return func(rw http.ResponseWriter, r *http.Request) {
			log.Println("Executing middleware for the group")
			next.ServeHTTP(rw, r)
		}
	}
	apiGroup.Middleware(groupMiddleware)

	// Routes within the group
	apiGroup.GET("/users", func(rw http.ResponseWriter, r *http.Request) {
		_, _ = rw.Write([]byte("API Group - List of Users"))
	})

	apiGroup.POST("/users", func(rw http.ResponseWriter, r *http.Request) {
		_, _ = rw.Write([]byte("API Group - Create User"))
	})

	// Another route outside the group
	ghk.GET("/home", func(rw http.ResponseWriter, r *http.Request) {
		_, _ = rw.Write([]byte("Home Page"))
	})

	err := http.ListenAndServe(":8080", ghk.Mux())
	if err != nil {
		log.Fatal(err)
	}
}
