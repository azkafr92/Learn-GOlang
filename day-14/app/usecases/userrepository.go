package usecases

import "day-14/app/domain"

type UserRepository interface {
	FindAll() ([]domain.User, error)
	Save(*domain.User) (*domain.User, error)
	FindById(id uint) (domain.User, error)
	UpdateById(id uint, data domain.User) (*domain.User, error)
	DeleteById(id uint) (*domain.User, error)
	FindByEmailAndPassword(email, password string) (domain.User, error)
}
