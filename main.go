package main

import (
	"net/http"

	"github.com/Dylvn/hashmal-go/config"
	"github.com/Dylvn/hashmal-go/offices"
	"github.com/Dylvn/hashmal-go/users"
)

func main() {
	// Assets
	http.Handle("/assets/", http.StripPrefix("/assets", http.FileServer(http.Dir("assets"))))

	// Users
	http.HandleFunc("/register", users.Register)
	http.HandleFunc("/register/process", users.RegisterProcess)
	http.HandleFunc("/login", users.Login)
	http.HandleFunc("/login/process", users.LoginProcess)
	http.HandleFunc("/profile", users.Profile)
	http.HandleFunc("/ajax/change-password", users.AjaxChangePassword)

	// Offices
	http.HandleFunc("/offices/create", offices.Create)
	http.HandleFunc("/offices/create/process", offices.Store)

	// Index
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	u, err := users.GetUser(w, r)
	if err != nil && err != users.ErrUserNotConnected {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}

	config.Tpl.ExecuteTemplate(w, "index.gohtml", struct {
		Title string
		User  *users.User
	}{
		"Home",
		u,
	})
}
