package handlers

import (
	"crypto/sha1"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"photo_blog/db"
	"strings"
)

type user_data struct {
	Name      string
	Galleries []string
}

func Get_redis_cookie(req *http.Request, cookie_name string) string {
	uuid, err := req.Cookie("session")
	if err != nil {
		panic(err)
	}
	username := db.RedisGetSession(uuid.Value)
	return username
}

func User_front(res http.ResponseWriter, req *http.Request) {
	username := Get_redis_cookie(req, "session")
	users_galleries := db.GetUsersGalleries(username)
	data := user_data{Name: username, Galleries: users_galleries}
	tpl.ExecuteTemplate(res, "user_front.html", data)
}

func Create_new_gallery(res http.ResponseWriter, req *http.Request) {
	username := Get_redis_cookie(req, "session")
	gallery_name := req.FormValue("gallery_name")
	db.Set_new_gallery(gallery_name, username)
	res.Header().Set("Location", "/user_front")
	res.WriteHeader(http.StatusSeeOther)
}

// rout to gallery.html
func Enter_gallery(res http.ResponseWriter, req *http.Request) {
	username := Get_redis_cookie(req, "session")
	var gallery_name string
	if req.Method == http.MethodGet {
		gallery_name = req.FormValue("enter_gallery_name")
		cookie := &http.Cookie{
			Name:     "gallery",
			Value:    gallery_name,
			HttpOnly: true,
			MaxAge:   600000 * 5,
			Path:     "/",
		}
		//do we use that?
		db.RedisSetSession(username, gallery_name)
		http.SetCookie(res, cookie)
	}
	if req.Method == http.MethodPost {
		Handle_pic_upload(res, req)
	}
	pics := db.Get_pics_annotations(username, gallery_name)
	tpl.ExecuteTemplate(res, "gallery.html", pics)
}

func Handle_pic_upload(res http.ResponseWriter, req *http.Request) {
	mf, fh, err := req.FormFile("new_pic")
	if err != nil {
		fmt.Println(err)
	}
	h := sha1.New()
	if _, err := io.Copy(h, mf); err != nil {
		log.Fatal(err)
	}
	defer mf.Close()
	split := strings.Split(fh.Filename, ".")
	ext := split[1]
	username := Get_redis_cookie(req, "session")
	gallery_name_cookie, err := req.Cookie("gallery")
	if err != nil {
		panic(err)
	}
	gallery_name := gallery_name_cookie.Value
	fname := fmt.Sprintf("%x", h.Sum(nil)) + "." + ext
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
	annotate := req.FormValue("annotate")
	db.Set_pic_annotate(username, gallery_name, fname, annotate)
	res.Header().Set("Location", "/enter_gallery?enter_gallery_name="+gallery_name)
	res.WriteHeader(http.StatusSeeOther)
}

func Publish_gallery(res http.ResponseWriter, req *http.Request) {
	//must have:
	username := Get_redis_cookie(req, "session")
	gallery_name_cookie, err := req.Cookie("gallery")
	if err != nil {
		panic(err)
	}
	gallery_name := gallery_name_cookie.Value
	// action:
	db.Publish(username, gallery_name)

	//return to user
	res.Header().Set("Location", "/enter_gallery?enter_gallery_name="+gallery_name)
	res.WriteHeader(http.StatusSeeOther)

}
