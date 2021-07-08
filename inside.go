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
	tmp := make([]string, 5)
	tmp = append(tmp, "myfirst")
	username := redisGetSession(uuid.String())
	println(username)
	//users_galleries := getUsersGalleries(username)
	data := user_data{Name: username, Galleries: tmp}
	tpl.ExecuteTemplate(res, "user_front.html", data)
}
