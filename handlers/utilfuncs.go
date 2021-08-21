package handlers

import (
	"net/http"
	"photo_blog/db"
)

func Get_redis_cookie(req *http.Request, cookie_name string) string {
	uuid, err := req.Cookie("session")
	if err != nil {
		panic(err)
	}
	username := db.RedisGetSession(uuid.Value)
	return username
}

func Get_galleryname_cookie(req *http.Request) (gallery_name string) {
	gallery_name_cookie, err := req.Cookie("gallery")
	if err != nil {
		panic(err)
	}
	gallery_name = gallery_name_cookie.Value
}
