package mydb

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

type User struct {
	Name     string `json:"name, omitempty"`
	Id       int64  `json:"id,string,omitempty"`
	Password string `json:"password, omitempty"`
	IsAdmin  bool   `json:"idAdmin,string,omitempty"`
}

var PgConnection = "user=postgres password=ginger dbname=go_test"

/******************
NO DATA MODIFICATION (read)
	use sql.Query()
					.Scan()
					.Close()
	swl.Rows maintains connection until Close() is executed: do NOT use this for modifying data
*/

/*
PingDB opens a DB connection using provided DSN and ping: returns whether connection successful
  does NOT close connection itself
*/
func PingDB(dbType string, DSN string) (*sql.DB, error) {
	db, err := sql.Open(dbType, DSN)
	if err != nil {
		return nil, err
	}
	//db.Open() doesn't actually open connection, so ping to make sure connection is good
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}

// query db by name string and return array of user structs
func GetManyUsers(db *sql.DB, username string) (interface{}, error) {
	var users []User
	rows, err := db.Query("select id, name, password from users where name = ?", username)
	if err != nil {
		return users, err
	}
	defer rows.Close()
	// iterate over returned database rows
	for rows.Next() {
		var thisUser User
		// add data from each column into variables
		err := rows.Scan(thisUser)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, thisUser)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return users, nil
}

// query db by name string and return single user struct
func GetOneUser(db *sql.DB, name string) (User, error) {
	var thisUser User
	// shortcut if only one row is expected back
	err := db.QueryRow("select id, name, password from users where name = $1;", name).Scan(&thisUser.Id, &thisUser.Name, &thisUser.Password)
	if err != nil {
		return User{}, err
	}
	return thisUser, nil
}

/******************/

/********
DATA MODIFICATION (write, uodate, delete)
	uses db.Prepare()
					.Exec()
					.LastInsertedId()
					.RowsAffected()
*/

// add one user to db
func InsertOneUser(db *sql.DB, user User) (int64, error) {
	smt, err := db.Prepare("INSERT INTO users(name, password) VALUES($1, $2)")
	if err != nil {
		return 0, err
	}
	res, err := smt.Exec(user.Name, user.Password)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

// update name of user by id
func UpdateOneUser(db *sql.DB, id int64, name string) (int64, error) {
	smt, err := db.Prepare("UPDATE users SET name = ? WHERE id = ?")
	if err != nil {
		return 0, err
	}
	res, err := smt.Exec(name, id)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}
