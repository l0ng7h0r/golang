package usecase

import (
	"errors"

	"github.com/l0ng7h0r/internal/domain"
	"github.com/l0ng7h0r/internal/repository"
	"github.com/l0ng7h0r/pkg/security"
)

type AuthUsecase struct {
	UserRepo   *repository.UserRepository
	SellerRepo *repository.SellerRepository
}

func NewAuthUsecase(userRepo *repository.UserRepository, sellerRepo *repository.SellerRepository) *AuthUsecase {
	return &AuthUsecase{
		UserRepo:   userRepo,
		SellerRepo: sellerRepo,
	}
}

func (u *AuthUsecase) Register(email string, password string, roles []string, seller *domain.Seller) error {
	hashed, err := security.HashPassword(password)
	if err != nil {
		return err
	}

	// ถ้า role เป็น seller ต้องมีข้อมูลร้าน
	isSeller := false
	for _, role := range roles {
		if role == "seller" {
			isSeller = true
			break
		}
	}

	if isSeller && (seller == nil || seller.ShopName == "") {
		return errors.New("shop_name is required for seller registration")
	}

	user := &domain.User{
		Email:    email,
		Password: hashed,
		Roles:    roles,
	}

	err = u.UserRepo.CreateUser(user)
	if err != nil {
		return err
	}

	// ถ้าเป็น seller ให้สร้าง seller profile ด้วย
	if isSeller && seller != nil {
		// หา user ที่เพิ่งสร้างเพื่อเอา ID
		createdUser, err := u.UserRepo.FindUserByEmail(email)
		if err != nil {
			return err
		}
		seller.UserID = int64(createdUser.ID)
		err = u.SellerRepo.CreateSeller(seller)
		if err != nil {
			return err
		}
	}

	return nil
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