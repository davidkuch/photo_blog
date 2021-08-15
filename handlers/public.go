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
	data := db.Get_published_galleries()
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
