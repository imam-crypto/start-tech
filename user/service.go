package user

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	RegisterUser(input RegisterUserInput) (User, error)
	Login(LoginInput) (User, error)
	IsEmailAvailable(input CheckEmaiInput) (bool, error)
	SaveAvatar(ID int, fileLocation string) (User, error)
	FindById(ID int) (User, error)
}
type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) RegisterUser(input RegisterUserInput) (User, error) {
	user := User{}
	user.First_name = input.First_name
	user.Email = input.Email

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)

	if err != nil {
		return user, nil
	}
	user.Password = string(passwordHash)

	// file,err := user.Form.

	newUser, er := s.repository.Save(user)

	if er != nil {
		return newUser, er
	}
	return newUser, nil
}

func (s *service) Login(input LoginInput) (User, error) {

	email := input.Email
	password := input.Password

	user, err := s.repository.FindByEmail(email)

	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, errors.New("User Not Found")
	}

	er := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if er != nil {
		return user, er
	}
	return user, nil
}

func (s *service) IsEmailAvailable(input CheckEmaiInput) (bool, error) {
	email := input.Email

	user, err := s.repository.FindByEmail(email)

	if err != nil {
		return false, err
	}
	if user.ID == 0 {
		return true, nil
	}
	return false, nil

}
func (s *service) SaveAvatar(ID int, fileLocation string) (User, error) {

	user, err := s.repository.FindByID(ID)

	if err != nil {
		return user, err
	}
	user.AvatarFileName = fileLocation
	UpdatedUser, er := s.repository.Update(user)
	if er != nil {
		return UpdatedUser, err
	}
	return UpdatedUser, nil
}
func (s *service) FindById(ID int) (User, error) {
	user, err := s.repository.FindByID(ID)

	if err != nil {
		return user, err
	}

	return user, nil

}
