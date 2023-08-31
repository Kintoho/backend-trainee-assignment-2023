package postgres

import (
	"fmt"

	"github.com/Kintoho/backend-trainee-assignment-2023/structure"

	"github.com/jmoiron/sqlx"
)

type SegmentPostgres struct {
	db *sqlx.DB
}

func NewSegmentPostgres(db *sqlx.DB) *SegmentPostgres {
	return &SegmentPostgres{db: db}
}

func (r *SegmentPostgres) Create(segment structure.Segment) (int, error) {
	var id int

	query := fmt.Sprintf("INSERT INTO %s (slug) VALUES ($1) RETURNING id", segmentsTable)
	row := r.db.QueryRow(query, segment.Slug)

	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *SegmentPostgres) Exists(slug string) (bool, error) {
	var exists bool

	query := fmt.Sprintf("SELECT EXISTS (SELECT * FROM %s WHERE slug=$1)", segmentsTable)
	row := r.db.QueryRow(query, slug)

	if err := row.Scan(&exists); err != nil {
		return false, err
	}

	return exists, nil
}

func (r *SegmentPostgres) Delete(slug string) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE slug=$1", segmentsTable)
	_, err := r.db.Exec(query, slug)
	return err
}
