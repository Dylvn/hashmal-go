package offices

import (
	"net/http"
	"strconv"
	"time"

	"github.com/Dylvn/hashmal-go/config"
	"github.com/Dylvn/hashmal-go/users"
)

func Create(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	u, err := users.GetUserAdmin(w, r)
	if err != nil {
		if err == users.ErrUserNotConnected || err == users.ErrUserNotAdmin {
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	config.Tpl.ExecuteTemplate(w, "offices_create.gohtml", struct {
		Title string
		User  *users.User
	}{
		"Create office",
		u,
	})
}

func Store(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	u, err := users.GetUserAdmin(w, r)
	if err != nil {
		if err == users.ErrUserNotConnected || err == users.ErrUserNotAdmin {
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	var active bool
	if r.FormValue("active") != "" {
		active = true
	}
	price, err := strconv.Atoi(r.FormValue("price"))
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	o := Office{}
	o.Name = r.FormValue("name")
	o.Descritpion = r.FormValue("desc")
	o.Price = price
	o.Active = active
	o.User = *u
	o.CreatedAt = time.Now()

	err = store(o)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	// TODO modify the route
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
