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

	// Creating the main route group using GroupFunc
	ghk.GroupFunc("/main", func(mainGroup *goHttpKit.Group) {
		// Middleware specific to the main route group
		mainGroup.Middleware(mainGroupMiddleware)

		// Creating a nested route group inside the main group using GroupFunc
		mainGroup.GroupFunc("/sub", func(subGroup *goHttpKit.Group) {
			// Middleware specific to the nested route group
			subGroup.Middleware(subGroupMiddleware)
			// Routes within the nested group
			subGroup.GET("/route1", handler1)
			subGroup.GET("/route2", handler2)
		})

		// Another route inside the main group
		mainGroup.GET("/route3", handler3)
	})

	// Another route outside the groups
	ghk.GET("/hello", hello)

	err := http.ListenAndServe(":8080", ghk.Mux())
	if err != nil {
		log.Fatal(err)
	}
}

func globalMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		log.Println("Executing global middleware")
		next.ServeHTTP(rw, r)
	}
}

func mainGroupMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		log.Println("Executing middleware for the main group")
		next.ServeHTTP(rw, r)
	}
}

func subGroupMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		log.Println("Executing middleware for the sub group")
		next.ServeHTTP(rw, r)
	}
}

func handler1(rw http.ResponseWriter, r *http.Request) {
	_, _ = rw.Write([]byte("Route 1"))
}

func handler2(rw http.ResponseWriter, r *http.Request) {
	_, _ = rw.Write([]byte("Route 2"))
}

func handler3(rw http.ResponseWriter, r *http.Request) {
	_, _ = rw.Write([]byte("Route 3"))
}
