package handlers

import (
	"net/http"

	"github.com/tgmendes/trigramerator/pkg/web"
)

type HelloResponse struct {
	Greet string `json:"greet"`
}

func handleHello(w http.ResponseWriter, r *http.Request) error {
	w.WriteHeader(http.StatusInternalServerError)

	resp := HelloResponse{
		Greet: web.GetParam(r.Context(), "name"),
	}

	err := web.RespondJSON(w, resp, http.StatusOK)
	return err
}
