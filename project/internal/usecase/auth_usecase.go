package usecase

import (
	"errors"
	"github.com/l0ng7h0r/internal/domain"
	"github.com/l0ng7h0r/internal/repository"
	"github.com/l0ng7h0r/pkg/security"
)

type AuthUsecase struct {
	UserRepo *repository.UserRepository
}

func NewAuthUsecase(repo *repository.UserRepository) *AuthUsecase {
	return  &AuthUsecase{UserRepo: repo}
}

func (u *AuthUsecase) Register(email string, password string, roles []string) error {
    hashed, err := security.HashPassword(password)
    if err != nil {
        return err
    }

    user := &domain.User{
        Email:    email,
        Password: hashed,
        Roles:    roles,
    }
    return u.UserRepo.CreateUser(user)
}
func (u *AuthUsecase) Login(email, password string) (string, error) {
	user, err := u.UserRepo.FindUserByEmail(email)
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	if !security.CheckPasswordHash(password, user.Password) {
		return "", errors.New("invalid credentials")
	}

	token, err := security.CreateToken(int64(user.ID), user.Roles)
	if err != nil {
		return "", err
	}
	return token, nil


}