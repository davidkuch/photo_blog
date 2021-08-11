package main

import (
	"hash"
	"html/template"
	"log"
	"net/http"

	uuid "github.com/satori/go.uuid"
)

var tpl *template.Template

var hashes = make(map[string]string, 100)

func init() {
	tpl = template.Must(template.ParseGlob("./templates/*.html"))

}

func main() {
	http.Handle("/", http.HandlerFunc(index))
	fs := http.FileServer(http.Dir("./public"))
	http.Handle("/public/", http.StripPrefix("/public", fs))
	http.Handle("/register", http.HandlerFunc(register))
	http.Handle("/login", http.HandlerFunc(login))
	http.Handle("/logout", http.HandlerFunc(logout))
	http.Handle("/user_front", http.HandlerFunc(user_front))
	http.Handle("/display_all_names", http.HandlerFunc(display_all_names))
	http.Handle("/create_new_gallery", http.HandlerFunc(create_new_gallery))
	http.Handle("/enter_gallery", http.HandlerFunc(enter_gallery))
	http.Handle("/remove_pic", http.HandlerFunc(remove_pic))
	http.Handle("/remove_gallery", http.HandlerFunc(remove_gallery))
	http.Handle("/publish_gallery", http.HandlerFunc(publish_gallery))
	http.Handle("/public_galleries", http.HandlerFunc(public_galleries))
	http.Handle("/display_published", http.HandlerFunc(display_published))

	http.ListenAndServe(":8080", nil)
}

//type rcv struct{
//	first map[string]string
//	second
//}

func index(res http.ResponseWriter, req *http.Request) {
	//adding random pic.
	// make a call-> store the response json
	// parse the json: take the first array object,
	//	respi, err := http.Get("https://api.thecatapi.com/v1/images/search")
	//	if err != nil {
	//		panic(err)
	//	}
	//	body, err := ioutil.ReadAll(respi.Body)
	//	if err != nil {
	//		panic(err)
	//	}
	//println(string(body))
	//	defer respi.Body.Close()
	//	var temp interface{}
	//	err = json.Unmarshal(body, &temp)
	//	if err != nil {
	///		panic(err)
	//}
	//println(string(body))
	//m := temp.(map[string]interface{})
	//for k, v := range m {
	//	switch vv := v.(type) {
	//	case string:
	//		fmt.Println(k, "is string", vv)
	//	case int:
	//		fmt.Println(k, "is int", vv)
	//	case []interface{}:
	//		fmt.Println(k, "is an array:")
	//		for i, u := range vv {
	//			fmt.Println(i, u)
	//		}
	//	default:
	//		fmt.Println(k, "is of a type I don't know how to handle")
	//	}
	//}
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
			MaxAge:   600000 * 5,
			Path:     "/",
		}
		redisSetSession(name, id.String())
		http.SetCookie(res, cookie)
		res.Header().Set("Location", "/user_front")
		res.WriteHeader(http.StatusSeeOther)
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
