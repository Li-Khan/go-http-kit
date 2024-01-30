package go_http_kit

import "net/http"

// HttpKit represents a wrapper around net/http package.
type HttpKit struct {
	// mux - ServeMux to handle HTTP requests
	mux *http.ServeMux
	// routes - List of routes
	routes []*Route
	// middlewares - List of middlewares to apply to all routes
	middlewares []http.Handler
}

// New creates a new instance of HttpKit.
func New() *HttpKit {
	return &HttpKit{
		mux: http.NewServeMux(),
	}
}
