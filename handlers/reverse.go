package handlers

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"photo_blog/db"
	"time"
)

func Logout(res http.ResponseWriter, req *http.Request) {
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

func Remove_pic(res http.ResponseWriter, req *http.Request) {
	pic_name := req.FormValue("remove_pic")
	//path :="C:\Users\david\photo_blog\public\pics\"
	db.Delete_pics([]string{pic_name})
	wd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	path := filepath.Join(wd, "public", "pics", pic_name)
	err = os.Remove(path)
	if err != nil {
		panic(err)
	}
	gallery_name_cookie, err := req.Cookie("gallery")
	if err != nil {
		panic(err)
	}
	gallery_name := gallery_name_cookie.Value
	res.Header().Set("Location", "/enter_gallery?enter_gallery_name="+gallery_name)
	res.WriteHeader(http.StatusSeeOther)

}

func Remove_gallery(res http.ResponseWriter, req *http.Request) {
	//can't use cookie: assume user not in  gallery
	gallery_name := req.FormValue("gallery_name")
	//actions:
	//	1- delete gallery from galleries table
	//	2- delete pics info from pics table- extract slice of pic names before
	//	3- delete pics files from 'public', using info from pics table before deletion
	user_name := Get_redis_cookie(req, "session")
	pics_annotations := db.Get_pics_annotations(user_name, gallery_name)
	pics := make([]string, len(pics_annotations))
	i := 0
	for pic := range pics_annotations {
		pics[i] = pic
		i++
	}
	db.Delete_pics(pics)

	wd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}

	for _, pic := range pics {
		path := filepath.Join(wd, "public", "pics", pic)
		err := os.Remove(path)
		if err != nil {
			panic(err)
		}

	}
	db.Delete_gallery(gallery_name)

	res.Header().Set("Location", "/user_front")
	res.WriteHeader(http.StatusSeeOther)
}
