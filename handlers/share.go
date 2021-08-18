package handlers

import (
	"fmt"
	"net/http"
	"photo_blog/db"
	"strings"
)

type Published_gallery struct {
	Owner        string
	Gallery_name string
	Pics         map[string]string
}

//display a list of public galleries with creators name as links
func Public_galleries(res http.ResponseWriter, req *http.Request) {
	var data map[string]string
	if req.FormValue("name") != "" {
		name := req.FormValue("name")
		data = db.Get_shared(name)

	} else {
		data = db.Get_published_galleries()
	}
	fmt.Println(data)
	tpl.ExecuteTemplate(res, "public_galleries.html", data)

}

//displays a published gallery- protected mode
func Display_published(res http.ResponseWriter, req *http.Request) {
	creds := req.FormValue("creds")
	//temp, unifnished
	split := strings.Split(creds, "!")
	username := split[1]
	gallery_name := split[0]
	pics := db.Get_pics_annotations(username, gallery_name)
	data := Published_gallery{username, gallery_name, pics}
	tpl.ExecuteTemplate(res, "public_gallery.html", data)

}

//shares a gallery with another user or group
// takes: (gallery id?)gallery name, owner name, others name,level
// action: insert to "shared" table in db
func Share_gallery(res http.ResponseWriter, req *http.Request) {
	owner := Get_redis_cookie(req, "session")
	gallery_name_cookie, err := req.Cookie("gallery")
	if err != nil {
		panic(err)
	}
	gallery := gallery_name_cookie.Value
	other := req.FormValue("other")
	level := "temp"
	db.Share_gallery(owner, other, gallery, level)
	//return to user
	res.Header().Set("Location", "/enter_gallery?enter_gallery_name="+gallery)
	res.WriteHeader(http.StatusSeeOther)
}
