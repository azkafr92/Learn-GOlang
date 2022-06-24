package interfaces

import (
	"day-13/app/domain"

	"gorm.io/gorm"
)

type BookRepository struct {
	DB *gorm.DB
}

func (ur *BookRepository) FindAll() ([]domain.Book, error) {
	var books []domain.Book
	if err := ur.DB.Find(&books).Error; err != nil {
		return books, err
	}
	return books, nil
}

func (ur *BookRepository) Save(book *domain.Book) (*domain.Book, error) {
	if err := ur.DB.Save(book).Error; err != nil {
		return book, err
	}
	return book, nil
}

func (ur *BookRepository) FindById(id uint) (domain.Book, error) {
	var book domain.Book
	if err := ur.DB.Where("id = ?", id).First(&book).Error; err != nil {
		return book, err
	}
	return book, nil
}

func (ur *BookRepository) UpdateById(id uint, data domain.Book) (*domain.Book, error) {
	book, err := ur.FindById(id)
	if err != nil {
		return &book, err
	}

	if data.Title != "" {
		book.Title = data.Title
	}

	if err := ur.DB.Save(book).Error; err != nil {
		return &book, err
	}
	
	return &book, nil
}

func (ur *BookRepository) DeleteById(id uint) (*domain.Book, error) {
	book, err := ur.FindById(id)
	if err != nil {
		return &book, err
	}
	
	if err := ur.DB.Delete(&book).Error; err != nil {
		return &book, err
	}

	return &book, nil
}
