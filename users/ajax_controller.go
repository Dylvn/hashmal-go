package users

import (
	"net/http"

	"github.com/Dylvn/hashmal-go/config"
	"golang.org/x/crypto/bcrypt"
)

func AjaxChangePassword(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	u, err := GetUser(w, r)
	if err != nil {
		if err == ErrUserNotConnected {
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	pass := r.FormValue("password")
	u.hashedPassword, err = bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	err = updatePassword(u, string(u.hashedPassword))
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	config.Tpl.ExecuteTemplate(w, "ajax_change_password.gohtml", nil)
}
