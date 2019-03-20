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

var vm1 = Server{VMName: "vm1", VMID: "3a2f9159-08f6-40df-9cd2-b37148dbbee", CPU: "80"}

var VMList = []Server{vm1}

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
