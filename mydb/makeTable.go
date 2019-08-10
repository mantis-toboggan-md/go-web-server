package mydb

import (
	"database/sql"
	"log"
)

// MakeTable creates the users table in the provided db instance
func MakeTable(db *sql.DB) error {
	_, err := db.Exec("CREATE TABLE users (id serial PRIMARY KEY, name varchar(255), password varchar(255), isAdmin bool);")
	if err != nil {
		return err
	}
	defer db.Close()
	log.Print("users table created")
	return err
}
