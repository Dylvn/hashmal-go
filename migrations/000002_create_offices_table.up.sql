CREATE TABLE IF NOT EXISTS offices (
    id serial PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description VARCHAR(255),
    price integer NOT NULL,
    active bool NOT NULL,
    sold bool NOT NULL,
    created_at timestamp NOT NULL,
    updated_at timestamp,
    user_id integer REFERENCES users (id)
)