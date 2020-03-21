package users

import (
	"net/http"

	"github.com/Dylvn/hashmal-go/config"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func authenticate(u *User, w http.ResponseWriter, r *http.Request) (*User, error) {
	uDB, err := getByUsername(u.Username)
	if err != nil {
		return nil, err
	}
	if err = bcrypt.CompareHashAndPassword([]byte(uDB.Password), []byte(u.Password)); err != nil {
		return nil, err
	}

	uuid := uuid.New()

	c := &http.Cookie{
		Name:     "session-id",
		Value:    uuid.String(),
		HttpOnly: true,
		MaxAge:   0,
		Path:     "/",
	}
	http.SetCookie(w, c)

	config.Session.Set(uuid.String(), uDB.ID)

	return uDB, nil
}

func isConnected(w http.ResponseWriter, r *http.Request) bool {
	c, err := r.Cookie("session-id")
	if err == http.ErrNoCookie {
		return false
	}

	_, err = config.Session.Get(c.Value)
	if err != nil {
		c := &http.Cookie{
			Name:   "session-id",
			MaxAge: -1,
		}
		http.SetCookie(w, c)

		return false
	}

	return true
}
