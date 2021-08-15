package handlers

import (
	"html/template"
	"log"
	"net/http"
	"photo_blog/db"

	uuid "github.com/satori/go.uuid"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("./templates/*.html"))

}

func Index(res http.ResponseWriter, req *http.Request) {

	tpl.ExecuteTemplate(res, "front.html", nil)

}

func Register(res http.ResponseWriter, req *http.Request) {
	name := req.FormValue("name")
	password := req.FormValue("password")
	db.InsertUser(name, password)
	err := tpl.ExecuteTemplate(res, "front.html", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func Login(res http.ResponseWriter, req *http.Request) {
	name := req.FormValue("name")
	password := req.FormValue("password")
	if db.IsUserCreds(name, password) {
		id := uuid.NewV4()
		cookie := &http.Cookie{
			Name:     "session",
			Value:    id.String(),
			HttpOnly: true,
			MaxAge:   600000 * 5,
			Path:     "/",
		}
		db.RedisSetSession(name, id.String())
		http.SetCookie(res, cookie)
		res.Header().Set("Location", "/user_front")
		res.WriteHeader(http.StatusSeeOther)
		return
	}
	tpl.ExecuteTemplate(res, "front.html", nil)
}

func Display(res http.ResponseWriter, req *http.Request) {
	name := req.FormValue("name")
	name_fixed := "/public/pics/" + name + ".jpg"
	tpl.ExecuteTemplate(res, "display.html", name_fixed)
}
