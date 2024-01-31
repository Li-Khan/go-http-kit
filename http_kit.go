package go_http_kit

import (
	"net/http"
	"sync"
)

// HttpKit represents a wrapper around net/http package.
type HttpKit struct {
	// mu - A mutex for ensuring synchronization when accessing HttpKit's data.
	mu *sync.Mutex
	// mux - ServeMux to handle HTTP requests
	mux *http.ServeMux
	// routes - List of routes
	routes []*Route
	// groupRoutes - List of groups of routes.
	groupRoutes []*Group
	// MethodNotAllowedHandler - Handler for Method Not Allowed responses.
	MethodNotAllowedHandler http.HandlerFunc
	// middlewares - List of middlewares to apply to all routes
	middlewares []Middleware
	cors        *CORS
}

// New creates a new instance of HttpKit.
func New() *HttpKit {
	return &HttpKit{
		mu:  &sync.Mutex{},
		mux: http.NewServeMux(),
		MethodNotAllowedHandler: func(rw http.ResponseWriter, r *http.Request) {
			http.Error(rw, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		},
	}
}

// Mux configures and returns the http.ServeMux for handling HTTP requests by iterating over defined routes.
func (hk *HttpKit) Mux() *http.ServeMux {
	for _, route := range hk.routes {
		handler := hk.configureRoute(route)
		hk.mux.HandleFunc(route.pattern, handler)
	}
	return hk.mux
}

func (hk *HttpKit) Cors(c *CORS) *HttpKit {
	hk.cors = c
	return hk
}

// configureRoute applies middlewares and configurations to a route's handler function.
// It returns the configured handler for the route.
func (hk *HttpKit) configureRoute(route *Route) func(http.ResponseWriter, *http.Request) {
	handler := route.handler
	for i := len(route.middlewares) - 1; i >= 0; i-- {
		handler = route.middlewares[i](handler)
	}

	for i := len(route.group.middlewares) - 1; i >= 0; i-- {
		handler = route.group.middlewares[i](handler)
	}

	for i := len(hk.middlewares) - 1; i >= 0; i-- {
		handler = hk.middlewares[i](handler)
	}

	handler = hk.methodMiddleware(handler, route.method)
	if hk.cors != nil {
		handler = hk.cors.Middleware(handler)
	}
	return handler
}

// methodMiddleware returns a middleware that checks if the incoming request uses the correct HTTP method.
func (hk *HttpKit) methodMiddleware(next http.HandlerFunc, method string) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		if r.Method != method {
			hk.MethodNotAllowedHandler.ServeHTTP(rw, r)
			return
		}
		next.ServeHTTP(rw, r)
	}
}
