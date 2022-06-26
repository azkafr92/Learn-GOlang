package usecases

import "day-14/app/domain"

type BookInteractor struct {
	BookRepository BookRepository
}

func (bi *BookInteractor) Index() ([]domain.Book, error) {
	return bi.BookRepository.FindAll()
}

func (bi *BookInteractor) Store(book *domain.Book) (*domain.Book, error) {
	return bi.BookRepository.Save(book)
}

func (bi *BookInteractor) Show(bookID uint) (domain.Book, error) {
	return bi.BookRepository.FindById(bookID)
}

func (bi *BookInteractor) Edit(bookID uint, data domain.Book) (*domain.Book, error) {
	return bi.BookRepository.UpdateById(bookID, data)
}

func (bi *BookInteractor) Destroy(bookID uint) (*domain.Book, error) {
	return bi.BookRepository.DeleteById(bookID)
}
