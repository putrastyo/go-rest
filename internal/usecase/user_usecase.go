package usecase

import (
	"errors"
	"praktik-todo/internal/entity"
	"praktik-todo/internal/repository"
	"praktik-todo/pkg/utils"
)

type UserUsecase interface {
	CreateUser(user *entity.User) error
	Login(email, password string) (string, error)
	GetAllUsers() ([]entity.User, error)
	GetUserByID(id uint) (*entity.User, error)
	GetUserByEmail(email string) (*entity.User, error)
	UpdateUser(user *entity.User) error
	DeleteUser(id uint) error
}

type userUsecase struct {
	userRepo repository.UserRepository
}

func NewUserUsecase(userRepo repository.UserRepository) UserUsecase {
	return &userUsecase{userRepo: userRepo}
}

func (u *userUsecase) CreateUser(user *entity.User) error {
	hashedPwd, err := utils.Hash(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashedPwd

	if err := u.userRepo.Create(user); err != nil {
		return err
	}
	return nil
}

func (u *userUsecase) Login(email string, password string) (string, error) {
	// search user by email
	user, err := u.userRepo.FindByEmail(email)
	if err != nil {
		return "", errors.New("invalid email or password")
	}

	// compare password
	if !utils.CheckHash(user.Password, password) {
		return "", errors.New("invalid email or password")
	}

	// return token
	return "JWT NGENTOT", nil
}

func (u *userUsecase) GetAllUsers() ([]entity.User, error) {
	return u.userRepo.FindAll()
}

func (u *userUsecase) GetUserByID(id uint) (*entity.User, error) {
	return u.userRepo.FindByID(id)
}

func (u *userUsecase) GetUserByEmail(email string) (*entity.User, error) {
	return u.userRepo.FindByEmail(email)
}

func (u *userUsecase) UpdateUser(user *entity.User) error {
	return u.userRepo.Update(user)
}

func (u *userUsecase) DeleteUser(id uint) error {
	return u.userRepo.Delete(id)
}
