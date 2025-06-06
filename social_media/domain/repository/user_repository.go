package repository

import "social_media/domain/entity"

type UserRepository interface {
	CreateUser(username, privacy string) (*entity.User, error)
}
