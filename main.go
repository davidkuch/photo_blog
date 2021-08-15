package main

import (
	"net/http"
	"photo_blog/handlers"
)

func main() {
	http.Handle("/", http.HandlerFunc(handlers.Index))
	fs := http.FileServer(http.Dir("./public"))
	http.Handle("/public/", http.StripPrefix("/public", fs))
	http.Handle("/register", http.HandlerFunc(handlers.Register))
	http.Handle("/login", http.HandlerFunc(handlers.Login))
	http.Handle("/logout", http.HandlerFunc(handlers.Logout))
	http.Handle("/user_front", http.HandlerFunc(handlers.User_front))
	http.Handle("/create_new_gallery", http.HandlerFunc(handlers.Create_new_gallery))
	http.Handle("/enter_gallery", http.HandlerFunc(handlers.Enter_gallery))
	http.Handle("/remove_pic", http.HandlerFunc(handlers.Remove_pic))
	http.Handle("/remove_gallery", http.HandlerFunc(handlers.Remove_gallery))
	http.Handle("/publish_gallery", http.HandlerFunc(handlers.Publish_gallery))
	http.Handle("/public_galleries", http.HandlerFunc(handlers.Public_galleries))
	http.Handle("/display_published", http.HandlerFunc(handlers.Display_published))

	http.ListenAndServe(":8080", nil)
}
