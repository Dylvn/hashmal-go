package users

import (
	"net/http"

	"github.com/Dylvn/hashmal-go/config"
	"golang.org/x/crypto/bcrypt"
)

func Register(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	_, err := GetUser(w, r)
	if err == nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	} else if err != ErrUserNotConnected {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	config.Tpl.ExecuteTemplate(w, "register.gohtml", struct {
		Title string
		User  *User
	}{
		"Register",
		nil,
	})
}

func RegisterProcess(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	var err error

	_, err = GetUser(w, r)
	if err == nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	} else if err != ErrUserNotConnected {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

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

	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	_, err := GetUser(w, r)
	if err == nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	} else if err != ErrUserNotConnected {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	config.Tpl.ExecuteTemplate(w, "login.gohtml", struct {
		Title string
		User  *User
	}{
		"Login",
		nil,
	})
}

func LoginProcess(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	var err error

	_, err = GetUser(w, r)
	if err == nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	} else if err != ErrUserNotConnected {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	u := User{}
	u.Username = r.FormValue("username")
	u.Password = r.FormValue("password")
	u.hashedPassword, err = bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	_, err = authenticate(&u, w, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func Profile(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
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

	config.Tpl.ExecuteTemplate(w, "profile.gohtml", struct {
		Title string
		User  *User
	}{
		Title: "Profile",
		User:  u,
	})
}
