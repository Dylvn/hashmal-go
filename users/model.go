package users

import (
	"github.com/Dylvn/hashmal-go/config"
)

type User struct {
	ID             int
	Username       string
	Password       string
	Email          string
	Admin          bool
	hashedPassword []byte
}

func store(u User) error {
	stmt, err := config.DB.Prepare("INSERT INTO users (username, password, email, admin) VALUES ($1, $2, $3, $4)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(u.Username, u.Password, u.Email, u.Admin)
	if err != nil {
		return err
	}

	return nil
}

func getByUsername(username string) (*User, error) {
	var u User
	row := config.DB.QueryRow("SELECT * FROM users WHERE username = $1", username)
	err := row.Scan(&u.ID, &u.Username, &u.Password, &u.Email, &u.Admin)
	if err != nil {
		return nil, err
	}

	return &u, nil
}
