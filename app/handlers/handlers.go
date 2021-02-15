package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/tgmendes/trigramerator/domain/trigram"
	"github.com/tgmendes/trigramerator/pkg/web"
)

// TrigramsAPI starts a server and defines the handlers to be used for the APP.
func TrigramsAPI(shutdown chan os.Signal, DB trigram.Storer) http.Handler {
	server := web.NewServer(shutdown)

	ts := trigramService{DB}
	server.Post("/learn", ts.handleLearn)
	server.Get("/generate", ts.handleGenerate)

	return server
}

type trigramService struct {
	db trigram.Storer
}

func (t trigramService) handleLearn(w http.ResponseWriter, r *http.Request) error {
	newText, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("an error occured reading provided text: %s", err.Error())
		web.RespondError(w, fmt.Sprintf("error loading input text: %s", err.Error()), http.StatusBadRequest)
		return err
	}

	err = trigram.Learn(t.db, string(newText))
	if err != nil {
		log.Printf("an error occurred learning text: %s", err.Error())
		web.RespondError(w, fmt.Sprintf("error occurred learning text: %s", err.Error()), http.StatusInternalServerError)
		return err
	}

	web.RespondText(w, "", http.StatusNoContent)
	return nil
}

func (t trigramService) handleGenerate(w http.ResponseWriter, r *http.Request) error {
	text, err := trigram.Generate(t.db, "", "")
	if err != nil {
		web.RespondError(w, fmt.Sprintf("error occurred generating text: %s", err.Error()), http.StatusInternalServerError)
		return err
	}

	web.RespondText(w, text, http.StatusOK)
	return nil
}
