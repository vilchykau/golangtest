package api

import (
	"github.com/gorilla/mux"
)

func InitRoutes(mux *mux.Router) {
	mux.HandleFunc("/subs", subsHandler)
}
