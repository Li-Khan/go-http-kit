package go_http_kit

import "net/http"

// Middleware adds middleware to be used by all routes.
func (hk *HttpKit) Middleware(handlers ...http.Handler) *HttpKit {
	hk.middlewares = append(hk.middlewares, handlers...)
	return hk
}

// Middleware adding middleware functionality for HttpKit and Route
func (r *Route) Middleware(handlers ...http.Handler) *Route {
	r.middlewares = append(r.middlewares, handlers...)
	return r
}
