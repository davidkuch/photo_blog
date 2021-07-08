package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "notesdb"
)

var db *sql.DB

func connect() {
	var psqlInfo = fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	var err error
	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	//defer db.Close()
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	println("got here!\n")
}

func InsertUser(name string, userpassword string) {
	connect()
	sqlStatement := `
INSERT INTO users (username,password)
VALUES ($1, $2)`
	_, err := db.Exec(sqlStatement, name, userpassword)
	if err != nil {
		panic(err)
	}
}

func isUserCreds(name string, userpassword string) (stt bool) {
	connect()
	sqlstt := "select * from users where username=$1 and password=$2;"
	var tmpname, tmppassword string
	row := db.QueryRow(sqlstt, name, userpassword)
	switch err := row.Scan(&tmpname, &tmppassword); err {
	case sql.ErrNoRows:
		return false
	case nil:
		return true
	default:
		panic(err)
	}

}

func isUser(name string) bool {
	println("flag 1 !! \n")
	connect()
	sqlstt := "select username from users where username=$1;"
	var tmpname string
	println("flag 3 !! \n")
	row := db.QueryRow(sqlstt, name)
	switch err := row.Scan(&tmpname); err {
	case sql.ErrNoRows:
		return false
	case nil:
		return true
	default:
		panic(err)
	}

}

func getNames() []string {
	connect()
	var names = make([]string, 0)
	sqlstt := `select username from users`
	rows, err := db.Query(sqlstt)
	if err != nil {
		// handle this error better than this
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		var name string
		err = rows.Scan(&name)
		if err != nil {
			// handle this error
			panic(err)
		}
		names = append(names, name)
	}
	// get any error encountered during iteration
	err = rows.Err()
	if err != nil {
		panic(err)
	}
	return names
}

func getUsersGalleries(username string) []string {
	connect()
	var names = make([]string, 0)
	sqlstt := `select gallery_name from galleries where username=$1;`
	rows, err := db.Query(sqlstt)
	if err != nil {
		// handle this error better than this
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		var name string
		err = rows.Scan(&name)
		if err != nil {
			// handle this error
			panic(err)
		}
		names = append(names, name)
	}
	// get any error encountered during iteration
	err = rows.Err()
	if err != nil {
		panic(err)
	}
	return names
}
