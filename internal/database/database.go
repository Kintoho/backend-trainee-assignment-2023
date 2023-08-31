package database

import (
	"errors"

	"github.com/Kintoho/backend-trainee-assignment-2023/internal/database/postgres"
	"github.com/Kintoho/backend-trainee-assignment-2023/structure"
	"github.com/jmoiron/sqlx"
)

var (
	ErrURLNotFound = errors.New("url not found")
	ErrURLExists   = errors.New("url exists")
)

type Authorization interface {
	CreateUser(user structure.User) (int, error)
	UserExists(user_id int) (bool, error)
}

type Segment interface {
	Create(segment structure.Segment) (int, error)
	Exists(slug string) (bool, error)
	Delete(slug string) error
}

type User interface {
	GetActiveSegment(user_id int) ([]structure.UserSegment, error)
	AddToSegment(user_id int, slug string) (int, error)
	SegmentRelationExists(user_id int, slug string) (bool, error)
	DeleteSegmentRelation(user_id int, slug string) error
}

type Database struct {
	Authorization
	Segment
	User
}

func NewDatabase(db *sqlx.DB) *Database {
	return &Database{
		Authorization: postgres.NewAuthPostgres(db),
		Segment:       postgres.NewSegmentPostgres(db),
		User:          postgres.NewUserPostgres(db),
	}
}
