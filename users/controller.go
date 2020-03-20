package users

import (
	"log"
	"net/http"

	"github.com/Dylvn/hashmal-go/config"
	"golang.org/x/crypto/bcrypt"
)

func Create(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Redirect(w, r, "/", http.StatusMethodNotAllowed)
		return
	}

	config.Tpl.ExecuteTemplate(w, "register.gohtml", nil)
}

func Store(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/", http.StatusMethodNotAllowed)
		return
	}
	var err error

	u := User{}
	u.Username = r.FormValue("username")
	pass := r.FormValue("password")
	u.Email = r.FormValue("email")
	u.hashedPassword, err = bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	u.Password = string(u.hashedPassword)

	err = store(u)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	log.Println("User created with success")
}
