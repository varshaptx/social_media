package repository

import "social_media/domain/entity"

type UserRepository interface {
	CreateUser(username, privacy string) (*entity.User, error)
	GetUser(username string) (*entity.User, error)
	GetAllUsers() map[string]*entity.User
	UpdateUser(user *entity.User) error
	DeleteUser(username string) error
}
