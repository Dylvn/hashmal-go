package users

import (
	"errors"
	"net/http"

	"github.com/Dylvn/hashmal-go/config"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

var ErrUserNotConnected = errors.New("User not connected")

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

// GetUser return an error if user not connected of type ErrUserNotConnected
// and an error if the DB can't scan the query
func GetUser(w http.ResponseWriter, r *http.Request) (*User, error) {
	if !isConnected(w, r) {
		return nil, ErrUserNotConnected
	}

	// Not need to get err since isConnected() already tested it
	c, _ := r.Cookie("session-id")
	id, _ := config.Session.Get(c.Value)
	u, err := getUserByID(id.(int))
	if err != nil {
		return nil, err
	}

	return u, nil
}
