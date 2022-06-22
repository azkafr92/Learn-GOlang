package interfaces

import (
	"day-12/app/domain"

	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func (ur *UserRepository) FindAll() ([]domain.User, error) {
	var users []domain.User
	if err := ur.DB.Find(&users).Error; err != nil {
		return users, err
	}
	return users, nil
}

func (ur *UserRepository) Save(user *domain.User) (*domain.User, error) {
	if err := ur.DB.Save(user).Error; err != nil {
		return user, err
	}
	return user, nil
}
