package repository

import(
	"database/sql"
	"github.com/l0ng7h0r/golang/internal/domain"
)

type SellerRepository struct {
	db *sql.DB
}

func NewSellerRepository(db *sql.DB) *SellerRepository {
	return &SellerRepository{db: db}
}

func (r *SellerRepository) CreateSeller(seller *domain.Seller) error {
	return nil
}