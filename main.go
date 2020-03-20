package main

import (
	"net/http"

	"github.com/Dylvn/hashmal-go/users"
)

func main() {
	http.Handle("/assets/", http.StripPrefix("/assets", http.FileServer(http.Dir("assets"))))
	http.HandleFunc("/users/create", users.Create)
	http.HandleFunc("/users/process", users.Store)
	http.ListenAndServe(":8080", nil)
}
