package handlers

import (
	"net/http"
	"os"

	"github.com/tgmendes/trigramerator/business/trigram"
	"github.com/tgmendes/trigramerator/pkg/web"
)

// API starts a server and defines the handlers to be used for the APP.
func API(shutdown chan os.Signal, DB trigram.Storer) http.Handler {
	server := web.NewServer(shutdown)

	ts := trigramService{DB}
	server.Post("/learn", ts.handleLearn)
	server.Get("/generate", ts.handleGenerate)

	return server
}
