package main

import (
	"fmt"
	"net/http"
)

//displaya list of public galleries with creators name as links
func public_galleries(res http.ResponseWriter, req *http.Request) {
	data := get_published_galleries()
	fmt.Println(data)
	tpl.ExecuteTemplate(res, "public_galleries.html", data)

}
