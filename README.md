# go-http-kit

`go-http-kit` is a Go package that provides a flexible and easy-to-use wrapper around the `net/http` package. It simplifies the process of defining HTTP routes, middleware, and handling CORS (Cross-Origin Resource Sharing) in your Go applications.

## Features

- **Route Handling:** Define HTTP routes easily for various methods (GET, POST, PUT, etc.).
- **Middleware Support:** Add middleware functions globally, per route, or per route group.
- **CORS Configuration:** Simplify CORS configuration with a dedicated package.
- **Route Groups:** Organize routes into groups for better structure and management.

## Installation

```bash
go get github.com/Li-Khan/go-http-kit
```

## Usage
Here is a basic example demonstrating the usage of go-http-kit:

```go
package main

import (
	goHttpKit "github.com/Li-Khan/go-http-kit"
	"log"
	"net/http"
)

func main() {
	// Create a new instance of HttpKit.
	ghk := goHttpKit.New()

	// Middleware example
	ghk.Middleware(M1, M2)

	// CORS example
	cors := goHttpKit.NewCORS().
		SetAllowedMethods("example1", "example2").
		SetAllowedOrigins("example1", "example2").
		SetAllowedHeaders("example1", "example2").
		SetExposeHeaders("example1", "example2")

	ghk.Cors(cors)

	// Define routes
	ghk.GET("/welcome/get", welcome)
	ghk.POST("/welcome/post", welcome)

	// Route groups
	v1Group := ghk.Group("/v1")
	v1Group.GET("/welcome/get", welcome)

	// Start the HTTP server
	err := http.ListenAndServe(":8080", ghk.Mux())
	if err != nil {
		log.Fatal(err)
	}
}

func welcome(rw http.ResponseWriter, r *http.Request) {
	_, _ = rw.Write([]byte(r.Method + " welcome, " + r.RemoteAddr))
}
```

## Middleware
Middleware functions can be added globally, per route, or per route group. They provide a way to execute code before or after handling an HTTP request

## CORS Configuration
The package includes a convenient CORS package for easy configuration of Cross-Origin Resource Sharing.

## Route Groups
Organize your routes into groups to improve code organization and maintainability.

## Examples
Check the [examples](https://github.com/Li-Khan/go-http-kit/blob/main/_examples) directory for additional usage examples.

## Contributing
Feel free to contribute by opening issues or submitting pull requests.

## License
This project is licensed under the MIT License - see the [LICENSE](https://github.com/Li-Khan/go-http-kit/blob/main/LICENSE) file for details.

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
