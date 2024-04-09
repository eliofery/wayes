package wayes

import (
	"fmt"
	"net/http"
)

// Validater is an interface that defines methods for configuring and performing validation.
type Validater interface {
	// Struct validates the structure of the provided data against predefined rules and
	// returns any validation errors encountered.
	Struct(data any) error
}

// Handler defines a function signature for handling HTTP requests.
type Handler func(ctx Ctx) error

// Wayes is an interface that defines methods for working with HTTP routes.
type Wayes interface {
	// Head registers a handler function for the HEAD method and the specified path.
	Head(path string, handler Handler)

	// Get registers a handler function for the GET method and the specified path.
	Get(path string, handler Handler)

	// Options registers a handler function for the Options method and the specified path.
	Options(path string, handler Handler)

	// Post registers a handler function for the POST method and the specified path.
	Post(path string, handler Handler)

	// Patch registers a handler function for the PATCH method and the specified path.
	Patch(path string, handler Handler)

	// Put registers a handler function for the PUT method and the specified path.
	Put(path string, handler Handler)

	// Delete registers a handler function for the DELETE method and the specified path.
	Delete(path string, handler Handler)

	// Group creates a new route group.
	Group(path string) Wayes

	// Use registers middleware for the wayes.
	Use(handlers ...Handler)

	// Combine combines multiple routers into a single wayes.
	Combine(routers ...*http.ServeMux) *http.ServeMux

	// Mux returns the underlying http.ServeMux.
	Mux() *http.ServeMux
}

// wayes represents a structure that implements the [wayes] interface.
type wayes struct {
	validator   Validater
	mux         *http.ServeMux
	middlewares []Handler
}

// New creates a new instance of [Wayes].
func New(validator ...Validater) Wayes {
	if len(validator) == 0 {
		validator = []Validater{nil}
	}

	return &wayes{
		validator:   validator[0],
		mux:         http.NewServeMux(),
		middlewares: make([]Handler, 0, 10),
	}
}

// handler executes the handler function and encodes the response.
func (rt *wayes) handler(handler Handler, w http.ResponseWriter, r *http.Request) {
	context := NewCtx(rt.validator, w, r)

	for _, middleware := range rt.middlewares {
		if err := middleware(context); err != nil {
			http.Error(context.Response(), err.Error(), http.StatusInternalServerError)
			return
		}
	}

	if err := handler(context); err != nil {
		http.Error(context.Response(), err.Error(), http.StatusInternalServerError)
	}
}

// Head registers a handler function for the HEAD method and the specified path.
func (rt *wayes) Head(path string, handler Handler) {
	rt.mux.HandleFunc(fmt.Sprintf("HEAD %s", path), func(w http.ResponseWriter, r *http.Request) {
		rt.handler(handler, w, r)
	})
}

// Get registers a handler function for the GET method and the specified path.
func (rt *wayes) Get(path string, handler Handler) {
	rt.mux.HandleFunc(fmt.Sprintf("GET %s", path), func(w http.ResponseWriter, r *http.Request) {
		rt.handler(handler, w, r)
	})
}

// Options registers a handler function for the Options method and the specified path.
func (rt *wayes) Options(path string, handler Handler) {
	rt.mux.HandleFunc(fmt.Sprintf("OPTIONS %s", path), func(w http.ResponseWriter, r *http.Request) {
		rt.handler(handler, w, r)
	})
}

// Post registers a handler function for the POST method and the specified path.
func (rt *wayes) Post(path string, handler Handler) {
	rt.mux.HandleFunc(fmt.Sprintf("POST %s", path), func(w http.ResponseWriter, r *http.Request) {
		rt.handler(handler, w, r)
	})
}

// Patch registers a handler function for the PATCH method and the specified path.
func (rt *wayes) Patch(path string, handler Handler) {
	rt.mux.HandleFunc(fmt.Sprintf("PATCH %s", path), func(w http.ResponseWriter, r *http.Request) {
		rt.handler(handler, w, r)
	})
}

// Put registers a handler function for the PUT method and the specified path.
func (rt *wayes) Put(path string, handler Handler) {
	rt.mux.HandleFunc(fmt.Sprintf("PUT %s", path), func(w http.ResponseWriter, r *http.Request) {
		rt.handler(handler, w, r)
	})
}

// Delete registers a handler function for the DELETE method and the specified path.
func (rt *wayes) Delete(path string, handler Handler) {
	rt.mux.HandleFunc(fmt.Sprintf("DELETE %s", path), func(w http.ResponseWriter, r *http.Request) {
		rt.handler(handler, w, r)
	})
}

// Group creates a new route group.
func (rt *wayes) Group(path string) Wayes {
	group := New(rt.validator)
	group.Use(rt.middlewares...)
	rt.mux.Handle(fmt.Sprintf("%s/", path), http.StripPrefix(path, group.Mux()))

	return group
}

// Use registers middleware for the wayes.
func (rt *wayes) Use(handlers ...Handler) {
	for _, handler := range handlers {
		rt.middlewares = append(rt.middlewares, handler)
	}
}

// Combine combines multiple routers into a single wayes.
func (rt *wayes) Combine(routers ...*http.ServeMux) *http.ServeMux {
	for _, router := range routers {
		rt.Mux().Handle("/", router)
	}

	return rt.Mux()
}

// Mux returns the underlying http.ServeMux.
func (rt *wayes) Mux() *http.ServeMux {
	return rt.mux
}
