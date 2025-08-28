package services

import (
	"errors"

	"github.com/riyajath-ahamed/microfrontend-golang/internal/models"
)

type UserService interface {
	GetUserByID(id string) (*models.User, error)
	CreateUser(user *models.User) (*models.User, error)
	ListUsers() ([]*models.User, error)
}

type userService struct {
	users map[string]*models.User
}

func NewUserService() UserService {
	return &userService{
		users: make(map[string]*models.User),
	}
}

func (s *userService) GetUserByID(id string) (*models.User, error) {
	if user, ok := s.users[id]; ok {
		return user, nil
	}

	return nil, errors.New("user not found")
}

func (s *userService) CreateUser(user *models.User) (*models.User, error) {
	if _, exists := s.users[user.ID]; exists {

		return nil, errors.New("user already exists")

	}
	s.users[user.ID] = user
	return user, nil
}

func (s *userService) ListUsers() ([]*models.User, error) {
	var list []*models.User
	for _, user := range s.users {
		list = append(list, user)
	}
	return list, nil
}
