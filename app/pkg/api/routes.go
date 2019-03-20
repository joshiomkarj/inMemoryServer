package api

import (
	"net/http"
)

func (app *App) SetRoutes() {
	app.Get("/servers", GetServers)

	app.Post("/servers", CreateServer)
	app.Delete("/servers/{id}", DeleteServer)
	app.Get("/servers/{id}", GetServer)
	app.Get("/servers/{id}/status", GetServerStatus)
	app.Get("/check/{name}", CheckName)

	// Using PATCH instead of PUT because PUT requires an entire body to be sent
	// effectively making PUT a subset of PATCH
	app.Patch("/servers/{id}", PatchServer)
}

func (app *App) Post(path string, f func(w http.ResponseWriter, r *http.Request)) {
	app.Router.HandleFunc(path, f).Methods("POST")
}

func (app *App) Get(path string, f func(w http.ResponseWriter, r *http.Request)) {
	app.Router.HandleFunc(path, f).Methods("Get")
}

func (app *App) Delete(path string, f func(w http.ResponseWriter, r *http.Request)) {
	app.Router.HandleFunc(path, f).Methods("Delete")
}

func (app *App) Patch(path string, f func(w http.ResponseWriter, r *http.Request)) {
	app.Router.HandleFunc(path, f).Methods("Patch")
}
