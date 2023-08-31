package postgres

import (
	"fmt"

	"github.com/Kintoho/backend-trainee-assignment-2023/structure"
	"github.com/jmoiron/sqlx"
)

type UserPostgres struct {
	db *sqlx.DB
}

func NewUserPostgres(db *sqlx.DB) *UserPostgres {
	return &UserPostgres{db: db}
}

func (r *UserPostgres) GetActiveSegment(user_id int) ([]structure.UserSegment, error) {
	var activeSegments []structure.UserSegment

	query := fmt.Sprintf(`SELECT user_id, slug
						FROM %s INNER JOIN %s
						ON segments.id = segment_id AND user_id=$1`, usersSegmentsRelationTable, segmentsTable)
	if err := r.db.Select(&activeSegments, query, user_id); err != nil {
		return nil, err
	}

	return activeSegments, nil
}

func (r *UserPostgres) AddToSegment(user_id int, slug string) (int, error) {
	var id int

	query := fmt.Sprintf(`INSERT INTO %s (user_id, segment_id)
						SELECT $1, id AS segment_id
						FROM %s
						WHERE slug=$2 RETURNING id`, usersSegmentsRelationTable, segmentsTable)
	row := r.db.QueryRow(query, user_id, slug)

	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *UserPostgres) SegmentRelationExists(user_id int, slug string) (bool, error) {
	var exists bool

	query := fmt.Sprintf(`SELECT EXISTS (
		SELECT * FROM %s
		INNER JOIN %s 
		ON user_id=$1 AND slug=$2 AND segments.id = segment_id)`, usersSegmentsRelationTable, segmentsTable)

	row := r.db.QueryRow(query, user_id, slug)

	if err := row.Scan(&exists); err != nil {
		return false, err
	}

	return exists, nil
}

func (r *UserPostgres) DeleteSegmentRelation(user_id int, slug string) error {
	query := fmt.Sprintf(`DELETE FROM %s
						USING %s
						WHERE user_id=$1 AND slug=$2 AND segment_id=segments.id`, usersSegmentsRelationTable, segmentsTable)
	_, err := r.db.Exec(query, user_id, slug)
	return err
}
