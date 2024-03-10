package usecase

import (
	"golang_cheatsheet/mock_gorm/internal/entity"
	"golang_cheatsheet/mock_gorm/internal/repository"
)

type UserUsecase interface {
	GetAll() ([]entity.User, error)
	GetUser(id int) (entity.User, error)
	CreateUser(user entity.User) error
	UpdateUser(id int, user entity.User) error
	DeleteUser(id int) error
}

type userUsecase struct {
	userRepo repository.UserRepository
}

func NewUserUsecase(userRepo repository.UserRepository) UserUsecase {
	return &userUsecase{userRepo}

}

func (u *userUsecase) GetAll() ([]entity.User, error) {
	return u.userRepo.GetAll()
}

func (u *userUsecase) GetUser(id int) (entity.User, error) {
	return u.userRepo.GetUser(id)
}

func (u *userUsecase) CreateUser(user entity.User) error {
	return u.userRepo.CreateUser(user)
}

func (u *userUsecase) UpdateUser(id int, user entity.User) error {
	return u.userRepo.UpdateUser(id, user)
}

func (u *userUsecase) DeleteUser(id int) error {
	return u.userRepo.DeleteUser(id)
}
