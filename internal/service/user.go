package service

import (
	"github.com/Kintoho/backend-trainee-assignment-2023/internal/database"
	"github.com/Kintoho/backend-trainee-assignment-2023/structure"
)

type UserService struct {
	db database.User
}

func NewUserService(repo database.User) *UserService {
	return &UserService{db: repo}
}

func (s *UserService) GetActiveSegment(user_id int) ([]structure.UserSegment, error) {
	return s.db.GetActiveSegment(user_id)
}

func (s *UserService) AddToSegment(user_id int, slug string) (int, error) {
	return s.db.AddToSegment(user_id, slug)
}

func (s *UserService) SegmentRelationExists(user_id int, slug string) (bool, error) {
	return s.db.SegmentRelationExists(user_id, slug)
}

func (s *UserService) DeleteSegmentRelation(user_id int, slug string) error {
	return s.db.DeleteSegmentRelation(user_id, slug)
}
