package usecases

import "day-15/app/domain"

type BookRepository interface {
	FindAll() ([]domain.Book, error)
	Save(*domain.Book) (*domain.Book, error)
	FindById(id uint) (domain.Book, error)
	UpdateById(id uint, data domain.Book) (*domain.Book, error)
	DeleteById(id uint) (*domain.Book, error)
}
