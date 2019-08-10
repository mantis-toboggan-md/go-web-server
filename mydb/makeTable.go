package mydb

import (
	"database/sql"
	"log"
)

// MakeTable creates the users table in the provided db instance
func MakeTable(db *sql.DB) error {
	smt, err := db.Prepare("CREATE TABLE users (id int NOT NULL AUTO_INCREMENT, name varchar(50), password varchar(50), PRIMARY KEY(id));")
	if err != nil {
		return err
	}
	_, err = smt.Exec()
	log.Print("users table created")
	return err
}
