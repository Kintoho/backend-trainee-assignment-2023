package service

import (
	"github.com/Kintoho/backend-trainee-assignment-2023/internal/database"
	"github.com/Kintoho/backend-trainee-assignment-2023/structure"
)

type AuthService struct {
	repo database.Authorization
}

func NewAuthService(repo database.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user structure.User) (int, error) {
	return s.repo.CreateUser(user)
}

func (s *AuthService) UserExists(user_id int) (bool, error) {
	return s.repo.UserExists(user_id)
}
