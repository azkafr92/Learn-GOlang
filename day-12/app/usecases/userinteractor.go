package usecases

import "day-12/app/domain"

type UserInteractor struct {
	UserRepository UserRepository
}

func (ui *UserInteractor) Index() ([]domain.User, error) {
	return ui.UserRepository.FindAll()
}

func (ui *UserInteractor) Store(user *domain.User) (*domain.User, error) {
	return ui.UserRepository.Save(user)
}
