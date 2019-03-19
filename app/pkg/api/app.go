package api

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

type App struct {
	Router *mux.Router
}

func (app *App) Initialize() {
	app.Router = mux.NewRouter()
	app.SetRoutes()
}

var VMList = []string{"a", "b", "c", "d"}

func (app *App) Run() {
	log.Println("Server started and listening for requests")
	corsOpts := cors.New(cors.Options{
		AllowedMethods: []string{
			http.MethodGet, //http methods for your app
			http.MethodPost,
			http.MethodPut,
			http.MethodPatch,
			http.MethodDelete,
			http.MethodOptions,
			http.MethodHead,
		},
		AllowedHeaders: []string{
			"*",
		},
		AllowedOrigins: []string{
			"*",
		},
	})
	log.Fatal(http.ListenAndServe(":8080", corsOpts.Handler(app.Router)))
}
