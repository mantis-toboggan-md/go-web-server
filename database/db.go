package database

import (
	"database/sql"
	"fmt"

	// Postgres driver imported anonymously so functions are not exposed here; using database/sql isntead
	_ "github.com/lib/pq"
)

func PingDb() {
	pgConnection := "user=postgres password=ginger dbname=go_test"
	db, err := sql.Open("postgres", pgConnection)
	if err != nil {
		fmt.Println(err)
	} else {
		err := db.Ping()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("connected")
	}
	defer db.Close()
}

func SelectOne() error {
	pgConnection := "user=postgres password=ginger dbname=go_test"
	db, err := sql.Open("postgres", pgConnection)
	if err != nil {
		fmt.Println(err)
	}
	var (
		id   int
		name string
	)
	rows, err := db.Query("SELECT id, username, password FROM users WHERE id = ?", 1)
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&id, &name)
		if err != nil {
			return err
		}
		fmt.Println(id, name)
	}

	defer db.Close()
	return nil
}
