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

	// Creating the main route group
	mainGroup := ghk.Group("/main")
	// Middleware specific to the main route group
	mainGroup.Middleware(mainGroupMiddleware)

	// Creating a nested route group inside the main group
	subGroup := mainGroup.Group("/sub")
	// Middleware specific to the nested route group
	subGroup.Middleware(subGroupMiddleware)

	// Routes within the nested group
	subGroup.GET("/route1", handler1)
	subGroup.GET("/route2", handler2)

	// Another route outside both groups
	ghk.GET("/hello", hello)

	err := http.ListenAndServe(":8080", ghk.Mux())
	if err != nil {
		log.Fatal(err)
	}
}
