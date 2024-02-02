package go_http_kit

import (
	"log"
	"net/http"
)

// Group represents a group of routes with a common URL pattern and middlewares.
type Group struct {
	// hk - Reference to the parent HttpKit instance.
	hk *HttpKit
	// pattern - Common URL pattern for all routes in the group.
	pattern string
	// middlewares - Middlewares to be applied to all routes in the group.
	middlewares []Middleware
	// routes - List of routes belonging to this group.
	routes []*Route
}

// Group creates a new group with the specified URL pattern and middlewares.
func (hk *HttpKit) Group(pattern string, middlewares ...Middleware) *Group {
	g := &Group{
		hk:          hk,
		pattern:     pattern,
		middlewares: middlewares,
	}
	hk.mu.Lock()
	hk.groupRoutes = append(hk.groupRoutes, g)
	hk.mu.Unlock()
	return g
}

// GroupFunc creates a new group with the specified URL pattern, executes the provided function, and applies middlewares.
func (hk *HttpKit) GroupFunc(pattern string, f func(g *Group), middlewares ...Middleware) *Group {
	g := &Group{
		hk:          hk,
		pattern:     pattern,
		middlewares: middlewares,
	}
	hk.mu.Lock()
	hk.groupRoutes = append(hk.groupRoutes, g)
	hk.mu.Unlock()
	f(g)
	return g
}

func (g *Group) Group(pattern string, middlewares ...Middleware) *Group {
	group := &Group{
		hk:          g.hk,
		pattern:     g.pattern + pattern,
		middlewares: append(g.middlewares, middlewares...),
	}
	g.hk.mu.Lock()
	g.hk.groupRoutes = append(g.hk.groupRoutes, group)
	g.hk.mu.Unlock()
	return g
}

func (g *Group) GroupFunc(pattern string, f func(g *Group), middlewares ...Middleware) *Group {
	group := &Group{
		hk:      g.hk,
		pattern: g.pattern + pattern,
	}
	group.middlewares = append(group.middlewares, g.middlewares...)
	group.middlewares = append(group.middlewares, middlewares...)
	g.hk.mu.Lock()
	g.hk.groupRoutes = append(g.hk.groupRoutes, group)
	g.hk.mu.Unlock()
	f(group)
	return g
}

// GET method for creating a new GET route with the specified pattern and handler.
func (g *Group) GET(pattern string, handler func(http.ResponseWriter, *http.Request)) *Route {
	log.Println(g.pattern + pattern)
	return g.hk.addRoute(g.pattern+pattern, http.MethodGet, handler, g)
}

// HEAD method for creating a new HEAD route with the specified pattern and handler.
func (g *Group) HEAD(pattern string, handler func(http.ResponseWriter, *http.Request)) *Route {
	return g.hk.addRoute(g.pattern+pattern, http.MethodHead, handler, g)
}

// POST method for creating a new POST route with the specified pattern and handler.
func (g *Group) POST(pattern string, handler func(http.ResponseWriter, *http.Request)) *Route {
	return g.hk.addRoute(g.pattern+pattern, http.MethodPost, handler, g)
}

// PUT method for creating a new PUT route with the specified pattern and handler.
func (g *Group) PUT(pattern string, handler func(http.ResponseWriter, *http.Request)) *Route {
	return g.hk.addRoute(g.pattern+pattern, http.MethodPut, handler, g)
}

// PATCH method for creating a new PATCH route with the specified pattern and handler.
func (g *Group) PATCH(pattern string, handler func(http.ResponseWriter, *http.Request)) *Route {
	return g.hk.addRoute(g.pattern+pattern, http.MethodPatch, handler, g)
}

// DELETE method for creating a new DELETE route with the specified pattern and handler.
func (g *Group) DELETE(pattern string, handler func(http.ResponseWriter, *http.Request)) *Route {
	return g.hk.addRoute(g.pattern+pattern, http.MethodDelete, handler, g)
}

// CONNECT method for creating a new CONNECT route with the specified pattern and handler.
func (g *Group) CONNECT(pattern string, handler func(http.ResponseWriter, *http.Request)) *Route {
	return g.hk.addRoute(g.pattern+pattern, http.MethodConnect, handler, g)
}

// OPTIONS method for creating a new OPTIONS route with the specified pattern and handler.
func (g *Group) OPTIONS(pattern string, handler func(http.ResponseWriter, *http.Request)) *Route {
	return g.hk.addRoute(g.pattern+pattern, http.MethodOptions, handler, g)
}

// TRACE method for creating a new TRACE route with the specified pattern and handler.
func (g *Group) TRACE(pattern string, handler func(http.ResponseWriter, *http.Request)) *Route {
	return g.hk.addRoute(g.pattern+pattern, http.MethodTrace, handler, g)
}
