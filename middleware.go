package go_http_kit

import "net/http"

// Middleware represents a function signature for HTTP middleware.
type Middleware func(next http.HandlerFunc) http.HandlerFunc

// Middleware adds middleware to be used by all routes.
func (hk *HttpKit) Middleware(middlewares ...Middleware) *HttpKit {
	hk.mu.Lock()
	defer hk.mu.Unlock()
	hk.middlewares = append(hk.middlewares, middlewares...)
	return hk
}

// Middleware method for Group adds one or more middleware functions to the group.
// Middleware functions are applied to all routes within the group, affecting the entire group's behavior.
func (g *Group) Middleware(middlewares ...Middleware) *Group {
	g.hk.mu.Lock()
	defer g.hk.mu.Unlock()
	g.middlewares = append(g.middlewares, middlewares...)
	return g
}

// Middleware adding middleware functionality for Route
func (r *Route) Middleware(middlewares ...Middleware) *Route {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.middlewares = append(r.middlewares, middlewares...)
	return r
}
