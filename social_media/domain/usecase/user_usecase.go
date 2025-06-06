package usecase

import (
	"social_media/domain/entity"
	"social_media/domain/service"
)

type UserUseCase interface {
	CreateUser(username, privacy string) (*entity.User, error)
}

type userUseCase struct {
	userService service.UserService
}

func NewUserUseCase(userService service.UserService) UserUseCase {
	return &userUseCase{
		userService: userService,
	}
}

func (u *userUseCase) CreateUser(username, privacy string) (*entity.User, error) {
	return u.userService.CreateUser(username, privacy)
}
