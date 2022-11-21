package app

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/vilchykau/golangtest/internal/api"
)

type OnShutdownExec func()

type App struct {
	onShutdownE []OnShutdownExec
}

func (a *App) Start() {
	mux := mux.NewRouter()

	api.InitRoutes(mux)

	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With"})
	originsOk := handlers.AllowedOrigins([]string{"localhost:8080", "127.0.0.1:5432"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	cors := handlers.CORS(originsOk, headersOk, methodsOk)

	fmt.Println("Server is listening...")
	log.Fatal(http.ListenAndServe(":8080"+os.Getenv("PORT"), cors(mux)))
}

func (a *App) AddOnShutdown(shut OnShutdownExec) {
	a.onShutdownE = append(a.onShutdownE, shut)
}

func (a *App) ShutDown() {
	for _, on := range a.onShutdownE {
		on()
	}
}
