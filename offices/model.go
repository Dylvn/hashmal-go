package offices

import (
	"time"

	"github.com/Dylvn/hashmal-go/config"
	"github.com/Dylvn/hashmal-go/users"
)

type Office struct {
	ID          int
	Name        string
	Description string
	Price       int
	Active      bool
	Sold        bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
	User        users.User
}

func store(o Office) error {
	stmt, err := config.DB.Prepare("INSERT INTO offices (name, description, price, active, sold, created_at, updated_at, user_id) VALUES ($1, $2, $3, $4, $5, $6, $7)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(o.Name, o.Description, o.Price, o.Active, false, o.CreatedAt, o.UpdatedAt, o.User.ID)
	if err != nil {
		return err
	}

	return nil
}

func count() (int, error) {
	offices, err := getAll()
	if err != nil {
		return 0, err
	}

	return len(offices), nil
}

func getAll() ([]Office, error) {
	var offices []Office
	var o Office
	var userId int
	rows, err := config.DB.Query("SELECT * FROM offices INNER JOIN users u on offices.user_id = u.id")
	if err != nil {
		return offices, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(
			&o.ID,            // offices.id
			&o.Name,          // offices.name
			&o.Description,   // offices.description
			&o.Price,         // offices.price
			&o.Active,        // offices.active
			&o.Sold,          // offices.sold
			&o.CreatedAt,     // offices.created_at
			&o.UpdatedAt,     // offices.updated_at
			&userId,          // offices.user_id
			&o.User.ID,       // users.id
			&o.User.Username, // users.username
			&o.User.Password, // users.password
			&o.User.Email,    // users.email
			&o.User.Admin,    // users.admin
		)
		if err != nil {
			return offices, nil
		}
		offices = append(offices, o)
	}

	return offices, nil
}

// getAllPaginate returns a number of Offices
// limit is the number of offices the function returns
// offset is the number of starting query
func getAllPaginate(limit, offset int) ([]Office, error) {
	var offices []Office
	var o Office
	var userId int
	rows, err := config.DB.Query("SELECT * FROM offices INNER JOIN users u on offices.user_id = u.id LIMIT $1 OFFSET $2", limit, offset)
	if err != nil {
		return offices, nil
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(
			&o.ID,            // offices.id
			&o.Name,          // offices.name
			&o.Description,   // offices.description
			&o.Price,         // offices.price
			&o.Active,        // offices.active
			&o.Sold,          // offices.sold
			&o.CreatedAt,     // offices.created_at
			&o.UpdatedAt,     // offices.updated_at
			&userId,          // offices.user_id
			&o.User.ID,       // users.id
			&o.User.Username, // users.username
			&o.User.Password, // users.password
			&o.User.Email,    // users.email
			&o.User.Admin,    // users.admin
		)
		if err != nil {
			return offices, nil
		}
		offices = append(offices, o)
	}

	return offices, nil
}
