package users

import (
	"net/http"

	"github.com/Dylvn/hashmal-go/config"
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
}
