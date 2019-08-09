package handlers

import (
	"fmt"
	"net/http"

	"github.com/mantis_toboggan_md/go_test/database"
)

// PublicServe returns handler to serve files from "./public"
// default serve index.html from "/"
func PublicServe(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	http.FileServer(http.Dir("public/")).ServeHTTP(w, r)
}

// OtherServe returns something obviously different to ensure router working correctly
func OtherServe(w http.ResponseWriter, r *http.Request) {
	database.PingDb()
	database.SelectOne()
	fmt.Fprintf(w, "Other page")
}
