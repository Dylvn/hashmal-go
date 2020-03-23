package offices

import (
	"time"

	"github.com/Dylvn/hashmal-go/config"
	"github.com/Dylvn/hashmal-go/users"
)

type Office struct {
	ID          int
	Name        string
	Descritpion string
	Price       int
	Active      bool
	Sold        bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
	User        users.User
}

func store(o Office) error {
	stmt, err := config.DB.Prepare("INSERT INTO offices (name, description, price, active, sold, created_at, user_id) VALUES ($1, $2, $3, $4, $5, $6, $7)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(o.Name, o.Descritpion, o.Price, o.Active, false, o.CreatedAt, o.User.ID)
	if err != nil {
		return err
	}

	return nil
}
