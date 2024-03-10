package repository

import (
	"golang_cheatsheet/mock_gorm/internal/entity"

	"gorm.io/gorm"
)

type UserRepository interface {
	GetAll() ([]entity.User, error)
	GetUser(id int) (entity.User, error)
	CreateUser(user entity.User) error
	UpdateUser(id int, user entity.User) error
	DeleteUser(id int) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

func (r *userRepository) GetAll() ([]entity.User, error) {
	err := r.db.Find(&[]entity.User{}).Error
	return []entity.User{}, err
}

func (r *userRepository) GetUser(id int) (entity.User, error) {
	err := r.db.First(&entity.User{}, id).Error
	return entity.User{}, err
}

func (r *userRepository) CreateUser(user entity.User) error {
	return r.db.Create(&user).Error
}

func (r *userRepository) UpdateUser(id int, user entity.User) error {
	return r.db.Model(&entity.User{}).Where("id = ?", id).Updates(user).Error
}

func (r *userRepository) DeleteUser(id int) error {
	return r.db.Delete(&entity.User{}, id).Error
}
