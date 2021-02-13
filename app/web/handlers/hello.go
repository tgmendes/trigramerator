package handlers

import (
	"github.com/tgmendes/go-service-template/pkg/web"
	"net/http"
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
