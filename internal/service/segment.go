package service

import (
	"github.com/Kintoho/backend-trainee-assignment-2023/internal/database"
	"github.com/Kintoho/backend-trainee-assignment-2023/structure"
)

type SegmentService struct {
	repo database.Segment
}

func NewSegmentService(repo database.Segment) *SegmentService {
	return &SegmentService{repo: repo}
}

func (s *SegmentService) Create(segment structure.Segment) (int, error) {
	return s.repo.Create(segment)
}

func (s *SegmentService) Exists(slug string) (bool, error) {
	return s.repo.Exists(slug)
}

func (s *SegmentService) Delete(slug string) error {
	return s.repo.Delete(slug)
}
