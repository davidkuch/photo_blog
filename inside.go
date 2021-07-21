package main

import (
	"crypto/sha256"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

type user_data struct {
	Name      string
	Galleries []string
}

func get_username_session(req *http.Request) string {
	uuid, err := req.Cookie("session")
	if err != nil {
		panic(err)
	}
	username := redisGetSession(uuid.Value)
	return username
}

func user_front(res http.ResponseWriter, req *http.Request) {
	username := get_username_session(req)
	users_galleries := getUsersGalleries(username)
	data := user_data{Name: username, Galleries: users_galleries}
	tpl.ExecuteTemplate(res, "user_front.html", data)
}

func create_new_gallery(res http.ResponseWriter, req *http.Request) {
	username := get_username_session(req)
	gallery_name := req.FormValue("gallery_name")
	set_new_gallery(gallery_name, username)
	res.Header().Set("Location", "/user_front")
	res.WriteHeader(http.StatusSeeOther)
}

// rout to gallery.html
func enter_gallery(res http.ResponseWriter, req *http.Request) {
	println("got here! \n")
	username := get_username_session(req)
	if req.Method == http.MethodPost {
		//error here: no such value given. we are uploading pic,
		// from inside gallery.html:where gallery name comes from to here?
		//gallery_name := req.FormValue("enter_gallery_name")
		annotate := req.FormValue("annotate")
		handle_pic_upload(res, req, username, gallery_name, annotate)
	}
	gallery_name := req.FormValue("enter_gallery_name")
	pics := get_pics_annotations(username, gallery_name)
	data := make(map[string]string)
	for pic_name, annotate := range pics {
		fullname := fmt.Sprintf(username + "#" + gallery_name + "#" + pic_name)
		data[fullname] = annotate
	}
	tpl.ExecuteTemplate(res, "gallery.html", data)
}

func handle_pic_upload(res http.ResponseWriter, req *http.Request, username string, gallery_name string, annotate string) {
	mf, fh, err := req.FormFile("new_pic")
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
	fname := fmt.Sprintf(username + "#" + gallery_name + "#" + name + "." + ext)
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
	set_pic_annotate(username, gallery_name, name, annotate)
	tpl.ExecuteTemplate(res, "gallery.html", nil)

}
