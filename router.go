package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mantis_toboggan_md/go_test/handlers"
)

/* GorrilaRouter creates mux router to manage "handlers"
public routes:
	GET "/" - index
	POST "/login" - login, get JWT
private routes:
	GET "/other" - placeholder page
*/
func GorillaRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", handlers.PublicServe)
	r.HandleFunc("/login", handlers.LogIn).Methods("POST")
	// middleware is ugly but if more is needed, alic package may be used
	r.Handle("/other", handlers.NeedsToken(http.HandlerFunc(handlers.OtherServe)))
	return r
}
