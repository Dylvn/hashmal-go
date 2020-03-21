package main

import (
	"net/http"

	"github.com/Dylvn/hashmal-go/config"
	"github.com/Dylvn/hashmal-go/users"
)

func main() {
	// Assets
	http.Handle("/assets/", http.StripPrefix("/assets", http.FileServer(http.Dir("assets"))))

	// Users
	http.HandleFunc("/register", users.Create)
	http.HandleFunc("/users/process", users.Store)
	http.HandleFunc("/login", users.Login)
	http.HandleFunc("/login/process", users.LoginProcess)

	// Offices

	// Index
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	config.Tpl.ExecuteTemplate(w, "index.gohtml", struct {
		Title string
	}{
		"Home",
	})
}
