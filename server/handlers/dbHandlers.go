package server

import (
	"database/sql"
	"net/http"
	"net/url"
	"strconv"

	"github.com/mantis_toboggan_md/sql_server/mydb"
)

// GetUser returns user info by id in route
func GetUser(h http.Handler, db *sql.DB) (http.Handler, error) {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.RawQuery
		queries, err := url.ParseQuery(query)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}
		id, err := strconv.ParseInt(queries["user_id"][0], 10, 64)
		if err != nil {
			// do soemthing if no id given
		}
		mydb.GetOneUser(db, id)
	}), nil
}
