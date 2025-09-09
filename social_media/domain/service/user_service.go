package service

import (
	"errors"
	"social_media/domain/entity"
	"social_media/domain/repository"
)

type UserService interface {
	CreateUser(username, privacy string) (*entity.User, error)
	GetUser(username string) (*entity.User, error)
	GetAllUsers() map[string]*entity.User
	UpdateUser(user *entity.User) error
	DeleteUser(username string) error
	ValidateUser(user *entity.User) error
}

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{
		userRepo: userRepo,
	}
}

func (s *userService) CreateUser(username, privacy string) (*entity.User, error) {
	if err := s.validatePrivacy(privacy); err != nil {
		return nil, err
	}

	user := entity.NewUser(username, privacy)
	if err := s.ValidateUser(user); err != nil {
		return nil, err
	}

	users := s.GetAllUsers()
	if _, exists := users[username]; exists {
		return nil, errors.New("user already exists")
	}

	return s.userRepo.CreateUser(username, privacy)
}

func (s *userService) GetUser(username string) (*entity.User, error) {
	return s.userRepo.GetUser(username)
}

func (s *userService) GetAllUsers() map[string]*entity.User {
	return s.userRepo.GetAllUsers()
}

func (s *userService) UpdateUser(user *entity.User) error {
	if err := s.ValidateUser(user); err != nil {
		return err
	}
	return s.userRepo.UpdateUser(user)
}

func (s *userService) DeleteUser(username string) error {
	return s.userRepo.DeleteUser(username)
}

func (s *userService) ValidateUser(user *entity.User) error {
	if user == nil {
		return errors.New("user cannot be nil")
	}
	if user.Username == "" {
		return errors.New("username cannot be empty")
	}
	return s.validatePrivacy(user.Privacy)
}

func (s *userService) validatePrivacy(privacy string) error {
	if privacy != "public" && privacy != "private" {
		return errors.New("invalid privacy setting")
	}
	return nil
}
