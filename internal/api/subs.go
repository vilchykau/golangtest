package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/vilchykau/golangtest/internal/controller"
)

type AddSubcriptionBody struct {
	Email string `json:"email" validate:"required" example:"Subcriber@email.com"`
	Url   string `json:"url" validate:"required" example:"https:https://kufar.by"`
}

func subsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		postSubcription(w, r)
	default:
		route404(w, r)
	}
}

func postSubcription(w http.ResponseWriter, r *http.Request) {
	var s AddSubcriptionBody
	json.NewDecoder(r.Body).Decode(&s)

	code, msg := controller.AddSubcription(s.Email, s.Url)

	w.WriteHeader(code)
	fmt.Fprint(w, msg)
}
