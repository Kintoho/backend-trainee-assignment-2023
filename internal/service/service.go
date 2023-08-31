package service

import (
	"github.com/Kintoho/backend-trainee-assignment-2023/internal/database"
	"github.com/Kintoho/backend-trainee-assignment-2023/structure"
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

type Service struct {
	Authorization
	Segment
	User
}

func NewService(repos *database.Database) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		Segment:       NewSegmentService(repos.Segment),
		User:          NewUserService(repos.User),
	}
}
