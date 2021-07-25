package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

func logout(res http.ResponseWriter, req *http.Request) {
	c := &http.Cookie{
		Name:    "session",
		Value:   "",
		Path:    "/",
		Expires: time.Unix(0, 0),

		HttpOnly: true,
	}
	http.SetCookie(res, c)
	res.Header().Set("Location", "/")
	res.WriteHeader(http.StatusSeeOther)
}

func remove_pic(res http.ResponseWriter, req *http.Request) {
	pic_name := req.FormValue("remove_pic")
	//path :="C:\Users\david\photo_blog\public\pics\"
	delete_pic(pic_name)
	wd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	path := filepath.Join(wd, "public", "pics", pic_name)
	err = os.Remove(path)
	if err != nil {
		panic(err)
	}
	res.Header().Set("Location", "/enter_gallery")
	res.WriteHeader(http.StatusSeeOther)

}
