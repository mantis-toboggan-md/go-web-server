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
	POST "/register" - make acct
private routes:
	GET "/other" - placeholder page
*/
func GorillaRouter() *mux.Router {
	r := mux.NewRouter()
	//serve create-react-app static file dir
	r.PathPrefix("/static").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("ui/build/static"))))
	r.HandleFunc("/", handlers.PublicServe)
	r.HandleFunc("/login", handlers.LogIn).Methods("POST")
	r.HandleFunc("/register", handlers.CreateAccount).Methods("POST")
	// middleware is ugly but if more is needed, alic package may be used
	r.Handle("/other", handlers.NeedsToken(http.HandlerFunc(handlers.OtherServe)))
	// check authorization header for token and return ok if valid token
	r.Handle("/auth", handlers.NeedsToken(http.HandlerFunc(handlers.Validated)))
	return r
}
