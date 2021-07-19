package main

import (
	"net/http"
)

type user_data struct {
	Name      string
	Galleries []string
}

func user_front(res http.ResponseWriter, req *http.Request) {
	uuid, err := req.Cookie("session")
	if err != nil {
		panic(err)
	}
	username := redisGetSession(uuid.Value)
	println("username by redis:" + username)
	users_galleries := getUsersGalleries(username)
	data := user_data{Name: username, Galleries: users_galleries}
	tpl.ExecuteTemplate(res, "user_front.html", data)
}

func create_new_gallery(res http.ResponseWriter, req *http.Request) {
	gallery_name := req.FormValue("gallery_name")
	uuid, err := req.Cookie("session")
	if err != nil {
		panic(err)
	}
	username := redisGetSession(uuid.Value)
	set_new_gallery(gallery_name, username)
	res.Header().Set("Location", "/user_front")
	res.WriteHeader(http.StatusSeeOther)
}
