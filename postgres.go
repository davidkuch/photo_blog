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

func set_pic_annotate(username string, gallery_name string, pic_name string, annotate string) {
	connect()
	date_uploaded := time.Now()
	sqlstt := `insert into pics values($1,$2,$3,$4,$5);`
	_, err := db.Exec(sqlstt, pic_name, annotate, date_uploaded, username, gallery_name)
	if err != nil {
		panic(err)
	}
}

func get_pics_annotations(username string, gallery_name string) map[string]string {
	connect()
	pics := make(map[string]string, 5)

	sqlstt := `select pic_name,annotate from pics where username=$1 and gallery_name=$2`
	rows, err := db.Query(sqlstt, username, gallery_name)
	if err != nil {
		// handle this error better than this
		panic(err)
	}

	for rows.Next() {
		var pic_name string
		var annotate string
		err = rows.Scan(&pic_name, &annotate)
		if err != nil {
			// handle this error
			panic(err)
		}
		pics[pic_name] = annotate
	}
	// get any error encountered during iteration
	err = rows.Err()
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	return pics
}
