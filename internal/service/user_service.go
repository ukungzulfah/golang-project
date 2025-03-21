package service

import (
	"analytics_project/internal/repository"
)

type UserService struct {
	Repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{Repo: repo}
}

func (s *UserService) GetAllUsersWithMenu() ([]map[string]interface{}, error) {
	return s.Repo.GetAllUsersWithMenu()
}

func (s *UserService) RunServiceQuery(name string) ([]map[string]interface{}, error) {
	return s.Repo.RunRepoQuery(name)
}
