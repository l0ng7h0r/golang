package usecase

import (
	"errors"

	"github.com/l0ng7h0r/internal/domain"
	"github.com/l0ng7h0r/internal/repository"
)

type ProductUsecase struct {
	ProductRepo *repository.ProductRepository
}

func NewProductUsecase(repo *repository.ProductRepository) *ProductUsecase {
	return &ProductUsecase{
		ProductRepo: repo,
	}
}

func (u *ProductUsecase) CreateProduct(product *domain.Product) error {
	if product.Price <= 0 {
        return errors.New("price must be greater than 0")
    }
    if product.Stock < 0 {
        return errors.New("stock must be greater than 0")
    }
    return u.ProductRepo.CreateProduct(product)
}

func (u *ProductUsecase) GetProduct(id int64) (*domain.Product, error) {
	return u.ProductRepo.GetProduct(id)
}

func (u *ProductUsecase) GetProducts() ([]*domain.Product, error) {
	return u.ProductRepo.GetProducts()
}

func (u *ProductUsecase) UpdateProduct(product *domain.Product) error {
	if product.Price <= 0 {
		return errors.New("price must be greater than 0")
	}
	if product.Stock < 0 {
		return errors.New("stock must be greater than 0")
	}
	return u.ProductRepo.UpdateProduct(product)
}

func (u *ProductUsecase) DeleteProduct(id int64) error {
	return u.ProductRepo.DeleteProduct(id)
}