package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "photo_blog_db"
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
}

func InsertUser(name string, userpassword string) {
	connect()
	sqlStatement := `INSERT INTO users (username,password)
VALUES ($1, $2)`
	_, err := db.Exec(sqlStatement, name, userpassword)
	if err != nil {
		panic(err)
	}
}

func isUserCreds(name string, userpassword string) (stt bool) {
	connect()
	sqlstt := "select username from users where username=$1 and password=$2;"
	var tmpname string
	row := db.QueryRow(sqlstt, name, userpassword)
	switch err := row.Scan(&tmpname); err {
	case sql.ErrNoRows:
		return false
	case nil:
		return true
	default:
		panic(err)
	}

}

func isUser(name string) bool {
	connect()
	sqlstt := "select username from users where username=$1;"
	var tmpname string
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
	rows, err := db.Query(sqlstt, username)
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

func set_new_gallery(gallery_name string, username string) {
	connect()
	date_created := time.Now()
	sqlstt := `insert into galleries values ($1,$2,$3,$4)`
	_, err := db.Exec(sqlstt, username, gallery_name, date_created, date_created)
	if err != nil {
		panic(err)
	}
}
package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "photo_blog_db"
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
}

func InsertUser(name string, userpassword string) {
	connect()
	sqlStatement := `INSERT INTO users (username,password)
VALUES ($1, $2)`
	_, err := db.Exec(sqlStatement, name, userpassword)
	if err != nil {
		panic(err)
	}
}

func isUserCreds(name string, userpassword string) (stt bool) {
	connect()
	sqlstt := "select username from users where username=$1 and password=$2;"
	var tmpname string
	row := db.QueryRow(sqlstt, name, userpassword)
	switch err := row.Scan(&tmpname); err {
	case sql.ErrNoRows:
		return false
	case nil:
		return true
	default:
		panic(err)
	}

}

func isUser(name string) bool {
	connect()
	sqlstt := "select username from users where username=$1;"
	var tmpname string
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
	rows, err := db.Query(sqlstt, username)
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

func set_new_gallery(gallery_name string, username string) {
	connect()
	date_created := time.Now()
	sqlstt := `insert into galleries values ($1,$2,$3,$4)`
	_, err := db.Exec(sqlstt, username, gallery_name, date_created, date_created)
	if err != nil {
		panic(err)
	}
}
