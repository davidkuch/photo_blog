package db

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

func IsUserCreds(name string, userpassword string) (stt bool) {
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

func IsUser(name string) bool {
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

func GetNames() []string {
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

func GetUsersGalleries(username string) []string {
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

func Set_new_gallery(gallery_name string, username string) {
	connect()
	date_created := time.Now()
	sqlstt := `insert into galleries values ($1,$2,$3,$4)`
	_, err := db.Exec(sqlstt, username, gallery_name, date_created, date_created)
	if err != nil {
		panic(err)
	}
}

func Set_pic_annotate(username string, gallery_name string, pic_name string, annotate string) {
	connect()
	date_uploaded := time.Now()
	sqlstt := `insert into pics values($1,$2,$3,$4,$5);`
	_, err := db.Exec(sqlstt, pic_name, annotate, date_uploaded, username, gallery_name)
	if err != nil {
		panic(err)
	}
}

func Get_pics_annotations(username string, gallery_name string) map[string]string {
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

func Delete_pics(pics []string) {
	connect()
	for _, pic := range pics {
		sqlstt := `delete from pics where pic_name=$1`
		_, err := db.Exec(sqlstt, pic)
		if err != nil {
			panic(err)
		}
	}
}

func Delete_gallery(gallery string) {
	connect()
	sqlstt := `delete from galleries where gallery_name=$1`
	_, err := db.Exec(sqlstt, gallery)
	if err != nil {
		panic(err)
	}
}

//gives map with names of all published galleries and their owners
func Get_published_galleries() map[string]string {
	connect()
	sqlstt := `select gallery_name,username from galleries where stat='published' `
	rows, err := db.Query(sqlstt)
	if err != nil {
		// handle this error better than this
		panic(err)
	}
	result := make(map[string]string)
	for rows.Next() {
		var gallery_name string
		var username string
		err = rows.Scan(&gallery_name, &username)
		if err != nil {
			// handle this error
			panic(err)
		}
		result[gallery_name] = username
	}
	// get any error encountered during iteration
	err = rows.Err()
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	return result
}

//to make a gallery public
//given gallery name and username
// change stat to 'public'
func Publish(username, gallery_name string) {
	connect()
	sqlstt := `update galleries set stat= 'published' where username=$1 and gallery_name=$2`
	_, err := db.Exec(sqlstt, username, gallery_name)
	if err != nil {
		panic(err)
	}

}

func Share_gallery(owner, other, gallery, level string) {
	connect()
	date := time.Now()
	sqlstt := `insert into shared values($1,$2,$3,$4,$5)`
	_, err := db.Exec(sqlstt, owner, other, gallery, date, level)
	if err != nil {
		panic(err)
	}
}

func Get_shared(name string) map[string]string {
	connect()
	sqlstt := `select owner,gallery from shared where other=$1`
	rows, err := db.Query(sqlstt, name)
	if err != nil {
		// handle this error better than this
		panic(err)
	}
	result := make(map[string]string)
	for rows.Next() {
		var gallery_name string
		var username string
		err = rows.Scan(&gallery_name, &username)
		if err != nil {
			// handle this error
			panic(err)
		}
		result[gallery_name] = username
	}
	// get any error encountered during iteration
	err = rows.Err()
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	return result
}
