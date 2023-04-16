package repository

import (
	"errors"

	"gorm.io/gorm"

	"mygram-api/helpers"
	"mygram-api/models/domain"
)

type UserRepository interface {
	Register(user *domain.User) (err error)
	Login(user *domain.User) (err error)
}

type UserRepositoryDB struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &UserRepositoryDB{DB: db}
}

func (userRepository *UserRepositoryDB) Register(user *domain.User) (err error) {
	
	if err = userRepository.DB.Create(&user).Error; err != nil {
		return err
	}

	return
}

func (userRepository *UserRepositoryDB) Login(user *domain.User) (err error) {

	password := user.Password

	err = userRepository.DB.Where("email = ?", user.Email).Take(&user).Error
	isValid := helpers.Compare([]byte(user.Password), []byte(password))

	if err != nil || !isValid {
		return errors.New("invalid email or password")
	}

	return
}