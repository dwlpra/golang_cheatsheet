package routine

import (
	"golang_cheatsheet/mock_gorm/internal/usecase"
)

type UserRoutine struct {
	UserUsecase usecase.UserUsecase
}

func NewUserRoutine(userUsecase usecase.UserUsecase) *UserRoutine {
	return &UserRoutine{userUsecase}
}
