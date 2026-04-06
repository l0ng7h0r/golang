package usecase

import (
	"errors"
	"os"
	"strconv"

	"github.com/l0ng7h0r/golang/internal/domain"
	"github.com/l0ng7h0r/golang/internal/repository"
	"github.com/l0ng7h0r/golang/pkg/security"
)

type AuthUsecase struct {
	userRepo *repository.UserRepository
}

func NewAuthUsecase(userRepo *repository.UserRepository) *AuthUsecase {
	return &AuthUsecase{userRepo: userRepo}
}

func (u *AuthUsecase) Register(email, password string, roles []string) error {
	hashed, err := security.HashPassword(password)
	if err != nil {
		return err
	}

	user := &domain.User{
		Email: email,
		Password: hashed,
		Roles: roles,
	}

	err = u.userRepo.CreateUser(user)
	if err != nil {
		return err
	}

	return nil
}	


func (u *AuthUsecase) Login(email, password string) (string, string, error) {
	user, err := u.userRepo.FindByEmail(email)
	if err != nil {
		return "", "", errors.New("invalid credentials")
	}

	if err := security.VerifyPassword(user.Password, password); err != nil {
		return "", "", errors.New("invalid credentials")
	}

	accessToken, err := security.GenerateToken(strconv.FormatInt(user.ID, 10), user.Roles, os.Getenv("ACCESS_TOKEN_EXPIRY"))
	if err != nil {
		return "", "", err
	}

	refreshToken, err := security.GenerateToken(strconv.FormatInt(user.ID, 10), user.Roles, os.Getenv("REFRESH_TOKEN_EXPIRY"))
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

func (u *AuthUsecase) Refresh(token string) (string, string, error) {
	userID, err := u.userRepo.ValidateRefreshToken(token)
	if err != nil {
		return "", "", errors.New("invalid refresh token")
	}

	_ = u.userRepo.DeleteRefreshToken(token)

	user, err := u.userRepo.FindByEmail(userID)
	if err != nil {
		return "", "", errors.New("user not found")
	}

	accessToken, err := security.GenerateToken(strconv.FormatInt(user.ID, 10), user.Roles, os.Getenv("ACCESS_TOKEN_EXPIRY"))
	if err != nil {
		return "", "", err
	}

	refreshToken, err := security.GenerateToken(strconv.FormatInt(user.ID, 10), user.Roles, os.Getenv("REFRESH_TOKEN_EXPIRY"))
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}
