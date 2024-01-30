package go_http_kit

import "net/http"

// Route represents an HTTP route with its pattern, method, handler, and middlewares.
type Route struct {
	pattern     string
	method      string
	handler     func(http.ResponseWriter, *http.Request)
	middlewares []http.Handler
}

// GET method for creating a new GET route with the specified pattern and handler.
func (hk *HttpKit) GET(pattern string, handler func(http.ResponseWriter, *http.Request)) *Route {
	return hk.addRoute(pattern, http.MethodGet, handler)
}

// HEAD method for creating a new HEAD route with the specified pattern and handler.
func (hk *HttpKit) HEAD(pattern string, handler func(http.ResponseWriter, *http.Request)) *Route {
	return hk.addRoute(pattern, http.MethodHead, handler)
}

// POST method for creating a new POST route with the specified pattern and handler.
func (hk *HttpKit) POST(pattern string, handler func(http.ResponseWriter, *http.Request)) *Route {
	return hk.addRoute(pattern, http.MethodPost, handler)
}

// PUT method for creating a new PUT route with the specified pattern and handler.
func (hk *HttpKit) PUT(pattern string, handler func(http.ResponseWriter, *http.Request)) *Route {
	return hk.addRoute(pattern, http.MethodPut, handler)
}

// PATCH method for creating a new PATCH route with the specified pattern and handler.
func (hk *HttpKit) PATCH(pattern string, handler func(http.ResponseWriter, *http.Request)) *Route {
	return hk.addRoute(pattern, http.MethodPatch, handler)
}

// DELETE method for creating a new DELETE route with the specified pattern and handler.
func (hk *HttpKit) DELETE(pattern string, handler func(http.ResponseWriter, *http.Request)) *Route {
	return hk.addRoute(pattern, http.MethodDelete, handler)
}

// CONNECT method for creating a new CONNECT route with the specified pattern and handler.
func (hk *HttpKit) CONNECT(pattern string, handler func(http.ResponseWriter, *http.Request)) *Route {
	return hk.addRoute(pattern, http.MethodConnect, handler)
}

// OPTIONS method for creating a new OPTIONS route with the specified pattern and handler.
func (hk *HttpKit) OPTIONS(pattern string, handler func(http.ResponseWriter, *http.Request)) *Route {
	return hk.addRoute(pattern, http.MethodOptions, handler)
}

// TRACE method for creating a new TRACE route with the specified pattern and handler.
func (hk *HttpKit) TRACE(pattern string, handler func(http.ResponseWriter, *http.Request)) *Route {
	return hk.addRoute(pattern, http.MethodTrace, handler)
}

// addRoute is an internal method to create a new route and add it to the HttpKit's routes list.
func (hk *HttpKit) addRoute(pattern string, method string, handler func(http.ResponseWriter, *http.Request)) *Route {
	hk.routes = append(hk.routes, &Route{
		pattern: pattern,
		method:  method,
		handler: handler,
	})
	return hk.routes[len(hk.routes)-1]
}
