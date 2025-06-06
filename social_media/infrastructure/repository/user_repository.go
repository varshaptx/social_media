package repository

import (
	"errors"
	"social_media/domain/entity"
)

type userRepository struct {
	users map[string]*entity.User
}

func NewUserRepository() *userRepository {
	return &userRepository{
		users: make(map[string]*entity.User),
	}
}

func (r *userRepository) CreateUser(username, privacy string) (*entity.User, error) {
	if _, exists := r.users[username]; exists {
		return nil, errors.New("user already exists")
	}

	user := entity.NewUser(username, privacy)
	r.users[username] = user
	return user, nil
}
