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

func (r *userRepository) GetUser(username string) (*entity.User, error) {
	user, exists := r.users[username]
	if !exists {
		return nil, errors.New("user not found")
	}
	return user, nil
}

func (r *userRepository) GetAllUsers() map[string]*entity.User {
	return r.users
}

func (r *userRepository) UpdateUser(user *entity.User) error {
	if _, exists := r.users[user.Username]; !exists {
		return errors.New("user not found")
	}
	r.users[user.Username] = user
	return nil
}

func (r *userRepository) DeleteUser(username string) error {
	if _, exists := r.users[username]; exists {
		return errors.New("user not found")
	}
	delete(r.users, username)
	return nil
}
