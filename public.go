package main

import (
	"fmt"
	"net/http"
	"strings"
)

//display a list of public galleries with creators name as links
func public_galleries(res http.ResponseWriter, req *http.Request) {
	data := get_published_galleries()
	fmt.Println(data)
	tpl.ExecuteTemplate(res, "public_galleries.html", data)

}

//displays a published gallery- protected mode
func display_published(res http.ResponseWriter, req *http.Request) {
	//must have:
	creds := req.FormValue("creds")
	//temp, unifnished
	split := strings.FieldsFunc()
	username := req.FormValue("username")
	gallery_name := req.FormValue("gallery_name")
	pics := get_pics_annotations(username, gallery_name)
	tpl.ExecuteTemplate(res, "public_gallery.html", pics)

}
