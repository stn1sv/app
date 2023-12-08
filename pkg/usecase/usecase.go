package usecase

import (
	"app/models"
	"app/pkg/repository"
)

type Authorization interface {
	CreateUser(user models.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}
type Usecase struct {
	Authorization
}

func NewUsecase(repos *repository.Repository) *Usecase {
	return &Usecase{NewAuthService(repos)}
}
