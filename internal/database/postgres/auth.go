package postgres

import (
	"fmt"

	"github.com/Kintoho/backend-trainee-assignment-2023/structure"
	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user structure.User) (int, error) {
	var id int

	query := fmt.Sprintf("INSERT INTO %s DEFAULT VALUES RETURNING id", usersTable)
	row := r.db.QueryRow(query)

	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *AuthPostgres) UserExists(user_id int) (bool, error) {
	var exists bool

	query := fmt.Sprintf("SELECT EXISTS (SELECT * FROM %s WHERE id=$1)", usersTable)
	row := r.db.QueryRow(query, user_id)

	if err := row.Scan(&exists); err != nil {
		return false, err
	}

	return exists, nil
}
