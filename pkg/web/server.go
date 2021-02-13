package web

import (
	"context"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"os"
	"syscall"
)

// A Handler is a type that handles an http request within our own http package.
type Handler func(w http.ResponseWriter, r *http.Request) error

// Server wraps an http router to handle requests. This makes it easy to switch between different router providers.
type Server struct {
	router *httprouter.Router

	// shutdown enables graceful shutdown mechanism
	shutdown chan os.Signal
}

// NewServer creates a new server with the default router.
func NewServer(shutdown chan os.Signal) *Server {
	// We could use a the standard Go mux, but httprouter is more flexible.
	// mux := http.NewServeMux()
	//
	// mux.HandleFunc("/hello", HandleHello)

	r := httprouter.New()

	return &Server{
		router: r,
		shutdown: shutdown,
	}
}

// SignalShutdown will gracefully shutdown the app.
func (s *Server) SignalShutdown()  {
	s.shutdown <- syscall.SIGTERM
}

// ServeHTTP implements the http.HandlerInterface. It defers the serving to the provided router.
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request)  {
	s.router.ServeHTTP(w, r)
}

// Handle defines a new endpoint to be handled by the server.
func (s *Server) Handle(method string, path string, handler Handler) {
	// To make our handler generic, we ignore the specific httprouter params. We can always access them through
	// context.
	h := func(w http.ResponseWriter, r *http.Request) {
		if err := handler(w, r); err != nil {
			s.SignalShutdown()
			return
		}
	}
	s.router.HandlerFunc(method, path, h)
}

// Get is a utility for GET requests.
func (s *Server) Get(path string, handler Handler) {
	s.Handle(http.MethodGet, path, handler)
}

// Post is a utility for POST requests.
func (s *Server) Post(path string, handler Handler) {
	s.Handle(http.MethodPost, path, handler)
}

// Put is a utility for PUT requests.
func (s *Server) Put(path string, handler Handler) {
	s.Handle(http.MethodPut, path, handler)
}

// Patch is a utility for PATCH requests.
func (s *Server) Patch(path string, handler Handler) {
	s.Handle(http.MethodPatch, path, handler)
}

// Delete is a utility for DELETE requests.
func (s *Server) Delete(path string, handler Handler) {
	s.Handle(http.MethodDelete, path, handler)
}

// GetParam is used to retrieve a parameter from the request (such as a path parameter).
func GetParam(ctx context.Context, key string) string {
	params := httprouter.ParamsFromContext(ctx)

	return params.ByName(key)
}
