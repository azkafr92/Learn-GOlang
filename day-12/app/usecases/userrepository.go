package usecases

import "day-12/app/domain"

type UserRepository interface {
	FindAll() ([]domain.User, error)
	Save(*domain.User) (*domain.User, error)
}
