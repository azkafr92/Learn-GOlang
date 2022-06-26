package usecases

import "day-14/app/domain"

type UserInteractor struct {
	UserRepository UserRepository
}

func (ui *UserInteractor) Index() ([]domain.User, error) {
	return ui.UserRepository.FindAll()
}

func (ui *UserInteractor) Store(user *domain.User) (*domain.User, error) {
	return ui.UserRepository.Save(user)
}

func (ui *UserInteractor) Show(userID uint) (domain.User, error) {
	return ui.UserRepository.FindById(userID)
}

func (ui *UserInteractor) Edit(userID uint, data domain.User) (*domain.User, error) {
	return ui.UserRepository.UpdateById(userID, data)
}

func (ui *UserInteractor) Destroy(userID uint) (*domain.User, error) {
	return ui.UserRepository.DeleteById(userID)
}

func (ui *UserInteractor) ShowByEmailAndPassword(email, password string) (domain.User, error) {
	return ui.UserRepository.FindByEmailAndPassword(email, password)
}
