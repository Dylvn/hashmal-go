package main

import (
	"net/http"

	"github.com/Dylvn/hashmal-go/users"
)

func main() {
	http.Handle("/assets/", http.StripPrefix("/assets", http.FileServer(http.Dir("assets"))))
	http.HandleFunc("/register", users.Create)
	http.HandleFunc("/users/process", users.Store)
	http.HandleFunc("/login", users.Login)
	http.HandleFunc("/login/process", users.LoginProcess)
	http.ListenAndServe(":8080", nil)
}
