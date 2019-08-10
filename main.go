package main

import (
	"log"
	"net/http"
)

func main() {
	r := GorillaRouter()
	// db, _ := mydb.PingDB("postgres", mydb.PgConnection)
	// err := mydb.MakeTable(db)
	// log.Printf(err.Error())
	if err := http.ListenAndServe(":5000", r); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}
