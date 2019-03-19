package api

import (
	"net/http"
)

func (app *App) SetRoutes() {
	app.Get("/servers", GetServers)

	/*app.Post("/servers", CreateServer)
	app.Delete("/servers", DeleteServer)
	app.Post("/servers/{id:[0-9]+}", GetServers)*/
}

func (app *App) Post(path string, f func(w http.ResponseWriter, r *http.Request)) {
	app.Router.HandleFunc(path, f).Methods("POST")
}

func (app *App) Get(path string, f func(w http.ResponseWriter, r *http.Request)) {
	app.Router.HandleFunc(path, f).Methods("Get")
}
