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
	//uuid "github.com/satori/go.uuid"
)

var tpl *template.Template

var hashes = make(map[string]string, 100)

func init() {
	tpl = template.Must(template.ParseGlob("./*.html"))

}

func main() {
	http.Handle("/", http.HandlerFunc(index))
	fs := http.FileServer(http.Dir("./public/pics"))
	http.Handle("/public/pics/", http.StripPrefix("/public/pics", fs))
	http.Handle("/display", http.HandlerFunc(display))
	http.ListenAndServe(":8080", nil)
}

func index(res http.ResponseWriter, req *http.Request) {
	// process form submission
	if req.Method == http.MethodPost {
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
		if isNew(h) {
			hashes[string(h.Sum(nil))] = name
		} else {
			println("pi exists! \n")
		}
		fname := fmt.Sprintf(name + "." + ext)
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
