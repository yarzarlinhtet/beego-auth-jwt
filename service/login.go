package service

import (
	"auth-api/models"
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

type LoginService interface {
	Login(username string, password string) bool
}

type loginService struct {
	user models.User
}

func NewLoginService(user models.User) LoginService {
	return &loginService{
		user: user,
	}
}

func (service *loginService) Login(username string, password string) bool {
	fmt.Println(service.user)

	if err := bcrypt.CompareHashAndPassword([]byte(service.user.Password), []byte(password)); err != nil {
		return false
	}

	return true
}
