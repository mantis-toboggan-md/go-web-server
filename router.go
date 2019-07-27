package main

import (
	"github.com/gorilla/mux"
	"github.com/mantis_toboggan_md/go_test/handlers"
)

// GorrilaRouter creates mux router to manage "handlers"
func GorillaRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", handlers.PublicServe)
	r.HandleFunc("/other", handlers.OtherServe)
	return r
}
