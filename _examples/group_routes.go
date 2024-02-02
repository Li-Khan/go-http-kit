package main

import (
	goHttpKit "github.com/Li-Khan/go-http-kit"
	"log"
	"net/http"
)

func main() {
	ghk := goHttpKit.New()

	// Global middleware for all routes
	ghk.Middleware(globalMiddleware)

	// Creating a route group
	apiGroup := ghk.Group("/api")

	// Middleware specific to the route group
	apiGroup.Middleware(groupMiddleware)

	// Routes within the group
	apiGroup.GET("/route1", handler1)

	apiGroup.POST("/route2", handler2)

	// Another route outside the group
	ghk.GET("/hello", hello)

	err := http.ListenAndServe(":8080", ghk.Mux())
	if err != nil {
		log.Fatal(err)
	}
}

func groupMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		log.Println("Executing middleware for the group")
		next.ServeHTTP(rw, r)
	}
}
