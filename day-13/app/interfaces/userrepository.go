package interfaces

import (
	"day-13/app/domain"

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

func (ur *UserRepository) FindById(id uint) (domain.User, error) {
	var user domain.User
	if err := ur.DB.Where("id = ?", id).First(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (ur *UserRepository) UpdateById(id uint, data domain.User) (*domain.User, error) {
	user, err := ur.FindById(id)
	if err != nil {
		return &user, err
	}

	if data.Email != "" {
		user.Email = data.Email
	}
	if data.Password != "" {
		user.Password = data.Password
	}

	if err := ur.DB.Save(user).Error; err != nil {
		return &user, err
	}
	
	return &user, nil
}

func (ur *UserRepository) DeleteById(id uint) (*domain.User, error) {
	user, err := ur.FindById(id)
	if err != nil {
		return &user, err
	}
	
	if err := ur.DB.Delete(&user).Error; err != nil {
		return &user, err
	}

	return &user, nil
}

func (ur *UserRepository) FindByEmailAndPassword(email, password string) (domain.User, error) {
	var user domain.User
	if err := ur.DB.Where("email = ?", email).Where("password = ?", password).First(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}
