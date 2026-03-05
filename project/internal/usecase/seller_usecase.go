package usecase

import (
	"errors"

	"github.com/l0ng7h0r/internal/domain"
	"github.com/l0ng7h0r/internal/repository"
)

type SellerUsecase struct {
	SellerRepo *repository.SellerRepository
}

func NewSellerUsecase(repo *repository.SellerRepository) *SellerUsecase {
	return &SellerUsecase{
		SellerRepo: repo,
	}
}

func (u *SellerUsecase) CreateSeller(seller *domain.Seller) error {
	if seller.ShopName == "" {
		return errors.New("shop name is required")
	}
	if seller.UserID == 0 {
		return errors.New("user id is required")
	}
	return u.SellerRepo.CreateSeller(seller)
}

func (u *SellerUsecase) GetSellerByID(id int64) (*domain.Seller, error) {
	return u.SellerRepo.GetSellerByID(id)
}

func (u *SellerUsecase) GetSellerByUserID(userID int64) (*domain.Seller, error) {
	return u.SellerRepo.GetSellerByUserID(userID)
}

func (u *SellerUsecase) UpdateSeller(seller *domain.Seller) error {
	if seller.ShopName == "" {
		return errors.New("shop name is required")
	}
	return u.SellerRepo.UpdateSeller(seller)
}

func (u *SellerUsecase) DeleteSeller(id int64) error {
	return u.SellerRepo.DeleteSeller(id)
}
