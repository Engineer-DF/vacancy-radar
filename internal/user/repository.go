package user

import "database/sql"

type Repository struct {
	db *sql.DB
}
