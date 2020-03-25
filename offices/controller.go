package offices

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/Dylvn/hashmal-go/config"
	"github.com/Dylvn/hashmal-go/users"
)

func List(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	var err error
	var oByPage int = 1
	var offset int
	var page int = 1
	var nextUrl, lastUrl, previousUrl string
	if r.FormValue("page") != "" {
		page, err = strconv.Atoi(r.FormValue("page"))
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		offset = (page * oByPage) - oByPage
	}

	nbOffices, err := count()
	if err != nil {
		log.Fatalln(err)
	}
	nbPages := nbOffices / oByPage
	if page < nbPages {
		nextUrl = fmt.Sprintf("/offices?page=%v", page+1)
	}
	if page > 1 {
		previousUrl = fmt.Sprintf("/offices?page=%v", page-1)
	}
	lastUrl = fmt.Sprintf("/offices?page=%v", nbPages)

	offices, err := getAllPaginate(oByPage, offset)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	u, err := users.GetUser(w, r)
	if err != nil && err != users.ErrUserNotConnected {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	config.Tpl.ExecuteTemplate(w, "offices_list.gohtml", struct {
		Title       string
		User        *users.User
		Offices     []Office
		NextUrl     string
		LastUrl     string
		PreviousUrl string
	}{
		"Listing offices",
		u,
		offices,
		nextUrl,
		lastUrl,
		previousUrl,
	})
}

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
	o.Description = r.FormValue("desc")
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
