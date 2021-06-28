package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("./*.html"))

}

func main() {
	http.Handle("/", http.HandlerFunc(front))
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.ListenAndServe(":8080", nil)
}

func front(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/html; charset=utf-8")
	err := tpl.ExecuteTemplate(res, "front.html", nil)
	if err != nil {
		log.Fatal(err)
	}
}
