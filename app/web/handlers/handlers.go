package handlers

import (
	"github.com/tgmendes/go-service-template/pkg/web"
	"net/http"
	"os"
)

// API starts a server and defines the handlers to be used for the APP.
func API(shutdown chan os.Signal) http.Handler {
	server := web.NewServer(shutdown)

	server.Get( "/hello/:name", handleHello)

	return server
}


