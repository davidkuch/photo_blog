package main

import (
	//	"crypto/sha1"
	"crypto/sha256"
	"fmt"
	"hash"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	uuid "github.com/satori/go.uuid"
)

var tpl *template.Template

var hashes = make(map[string]string, 100)

func init() {
	tpl = template.Must(template.ParseGlob("./*.html"))

}

func main() {
	http.Handle("/", http.HandlerFunc(index))
	fs := http.FileServer(http.Dir("./public"))
	http.Handle("/public/", http.StripPrefix("/public", fs))
	// i would try to make it static served
	//http.Handle("/registry",http.HandlerFunc(register_place))
	http.Handle("/register", http.HandlerFunc(register))
	// the above
	//http.Handle("/loginery",http.HandlerFunc(loginery))
	http.Handle("/login", http.HandlerFunc(login))
	//http.Handle("/user_front", http.HandlerFunc(user_front))
	http.Handle("/display", http.HandlerFunc(display))
	http.Handle("/display_all_names", http.HandlerFunc(display_all_names))
	http.ListenAndServe(":8080", nil)
}

func index(res http.ResponseWriter, req *http.Request) {
	// process form submission
	if req.Method == http.MethodPost {
		handle_upload(res, req)
		return
	}
	tpl.ExecuteTemplate(res, "front.html", nil)

}

func handle_upload(res http.ResponseWriter, req *http.Request) {
	mf, fh, err := req.FormFile("nf")
	if err != nil {
		fmt.Println(err)
	}
	h := sha256.New()
	if _, err := io.Copy(h, mf); err != nil {
		log.Fatal(err)
	}
	defer mf.Close()
	split := strings.Split(fh.Filename, ".")
	name, ext := split[0], split[1]
	if !isNew(h) {
		tpl.ExecuteTemplate(res, "error.html", "pic already uploaded!")
		return
	}
	hashes[string(h.Sum(nil))] = name
	//temp username: to be taken from db by sesion uuid
	username := "temp_username"
	fname := fmt.Sprintf(username + name + "." + ext)
	// create new fileS
	wd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	path := filepath.Join(wd, "public", "pics", fname)
	nf, err := os.Create(path)
	if err != nil {
		fmt.Println(err)
	}
	defer nf.Close()
	// copy
	mf.Seek(0, 0)
	io.Copy(nf, mf)

	tpl.ExecuteTemplate(res, "front.html", nil)

}

func register(res http.ResponseWriter, req *http.Request) {
	name := req.FormValue("name")
	password := req.FormValue("password")
	//if isUser(name) {
	//	err := tpl.ExecuteTemplate(res, "registery.html", name+" already exists")
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//	return
	//}
	InsertUser(name, password)
	err := tpl.ExecuteTemplate(res, "front.html", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func login(res http.ResponseWriter, req *http.Request) {
	//res.Header().Set("Content-Type", "text/html; charset=utf-8")
	name := req.FormValue("name")
	password := req.FormValue("password")
	if isUserCreds(name, password) {
		id := uuid.NewV4()
		cookie := &http.Cookie{
			Name:     "session",
			Value:    id.String(),
			HttpOnly: true,
			MaxAge:   600 * 5,
			Path:     "/",
		}
		redisSetSession(name, id.String())
		http.SetCookie(res, cookie)
		user_front(res, req)
		return
	}
	tpl.ExecuteTemplate(res, "front.html", nil)
}

func display(res http.ResponseWriter, req *http.Request) {
	name := req.FormValue("name")
	name_fixed := "/public/pics/" + name + ".jpg"
	tpl.ExecuteTemplate(res, "display.html", name_fixed)
}

func isNew(h hash.Hash) bool {
	if _, ok := hashes[string(h.Sum(nil))]; ok {
		return false
	}
	return true

}

func get_all_names() []string {
	names := make([]string, 0)
	//val in hashes is pics name
	for _, name := range hashes {
		names = append(names, name)
	}
	return names
}

func display_all_names(res http.ResponseWriter, req *http.Request) {
	names := get_all_names()
	tpl.ExecuteTemplate(res, "display_all_names.html", names)

}
