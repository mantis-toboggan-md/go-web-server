package server

import (
	"net/http"
)

// StartListening starts http server on provided port :0000
func StartListening(port string) error {
	http.Handle("/", http.FileServer(http.Dir("ui/build/")))
	if err := (http.ListenAndServe(port, nil)); err != nil {
		return err
	}
	return nil
}
