package go_http_kit

import (
	"net/http"
	"sync"
)

// Route represents an HTTP route with its pattern, method, handler, and middlewares.
type Route struct {
	// mu - A mutex for ensuring synchronization when accessing route-specific data.
	mu *sync.Mutex
	// pattern - URL pattern for the route.
	pattern string
	// method - HTTP method for the route (e.g., GET, POST).
	method string
	// handler - Handler function for processing the HTTP request.
	handler func(http.ResponseWriter, *http.Request)
	// middlewares - List of middlewares to be applied to this specific route.
	middlewares []Middleware
	// group - Reference to the group to which this route belongs.
	group *Group
}

// addRoute is an internal method to create a new route and add it to the HttpKit's routes list.
func (hk *HttpKit) addRoute(pattern string, method string, handler func(http.ResponseWriter, *http.Request), group *Group) *Route {
	hk.mu.Lock()
	defer hk.mu.Unlock()
	hk.routes = append(hk.routes, &Route{
		mu:      &sync.Mutex{},
		pattern: pattern,
		method:  method,
		handler: handler,
		group:   group,
	})
	return hk.routes[len(hk.routes)-1]
}

// GET method for creating a new GET route with the specified pattern and handler.
func (hk *HttpKit) GET(pattern string, handler func(http.ResponseWriter, *http.Request)) *Route {
	return hk.addRoute(pattern, http.MethodGet, handler, nil)
}

// HEAD method for creating a new HEAD route with the specified pattern and handler.
func (hk *HttpKit) HEAD(pattern string, handler func(http.ResponseWriter, *http.Request)) *Route {
	return hk.addRoute(pattern, http.MethodHead, handler, nil)
}

// POST method for creating a new POST route with the specified pattern and handler.
func (hk *HttpKit) POST(pattern string, handler func(http.ResponseWriter, *http.Request)) *Route {
	return hk.addRoute(pattern, http.MethodPost, handler, nil)
}

// PUT method for creating a new PUT route with the specified pattern and handler.
func (hk *HttpKit) PUT(pattern string, handler func(http.ResponseWriter, *http.Request)) *Route {
	return hk.addRoute(pattern, http.MethodPut, handler, nil)
}

// PATCH method for creating a new PATCH route with the specified pattern and handler.
func (hk *HttpKit) PATCH(pattern string, handler func(http.ResponseWriter, *http.Request)) *Route {
	return hk.addRoute(pattern, http.MethodPatch, handler, nil)
}

// DELETE method for creating a new DELETE route with the specified pattern and handler.
func (hk *HttpKit) DELETE(pattern string, handler func(http.ResponseWriter, *http.Request)) *Route {
	return hk.addRoute(pattern, http.MethodDelete, handler, nil)
}

// CONNECT method for creating a new CONNECT route with the specified pattern and handler.
func (hk *HttpKit) CONNECT(pattern string, handler func(http.ResponseWriter, *http.Request)) *Route {
	return hk.addRoute(pattern, http.MethodConnect, handler, nil)
}

// OPTIONS method for creating a new OPTIONS route with the specified pattern and handler.
func (hk *HttpKit) OPTIONS(pattern string, handler func(http.ResponseWriter, *http.Request)) *Route {
	return hk.addRoute(pattern, http.MethodOptions, handler, nil)
}

// TRACE method for creating a new TRACE route with the specified pattern and handler.
func (hk *HttpKit) TRACE(pattern string, handler func(http.ResponseWriter, *http.Request)) *Route {
	return hk.addRoute(pattern, http.MethodTrace, handler, nil)
}
