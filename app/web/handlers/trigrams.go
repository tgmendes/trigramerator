package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/tgmendes/trigramerator/business/trigram"
	"github.com/tgmendes/trigramerator/pkg/web"
)

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

	go func() {
		err := trigram.Learn(t.db, string(newText))
		if err != nil {
			log.Printf("an error occurred learning text: %s", err.Error())
			return
		}
		log.Printf("new text successfully learned!")
	}()

	web.RespondText(w, "", http.StatusAccepted)
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
