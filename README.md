# hashmal-go
A simple website created in Go. That's my first app in go for try the language.

# Installation

Download <a href="https://github.com/golang-migrate/migrate">migrate</a>

```bash
migrate -database "postgres://user:password@localhost:5432/database-name?sslmode=disable" -path migrations up
```

## Create new migration

```cli
migrate create -ext sql -dir migrations -seq migration_name
```

