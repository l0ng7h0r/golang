package repository

import (
	"database/sql"

	"github.com/l0ng7h0r/internal/domain"
)

type SellerRepository struct {
	db *sql.DB
}

func NewSellerRepository(db *sql.DB) *SellerRepository {
	return &SellerRepository{
		db: db,
	}
}

func (r *SellerRepository) CreateSeller(seller *domain.Seller) error {
	_, err := r.db.Exec(
		`INSERT INTO sellers (user_id, shop_name, description, phone) VALUES ($1, $2, $3, $4)`,
		seller.UserID, seller.ShopName, seller.Description, seller.Phone,
	)
	return err
}

func (r *SellerRepository) GetSellerByID(id int64) (*domain.Seller, error) {
	row := r.db.QueryRow(
		`SELECT id, user_id, shop_name, description, phone FROM sellers WHERE id = $1`, id,
	)
	seller := &domain.Seller{}
	err := row.Scan(&seller.ID, &seller.UserID, &seller.ShopName, &seller.Description, &seller.Phone)
	if err != nil {
		return nil, err
	}
	return seller, nil
}

func (r *SellerRepository) GetSellerByUserID(userID int64) (*domain.Seller, error) {
	row := r.db.QueryRow(
		`SELECT id, user_id, shop_name, description, phone FROM sellers WHERE user_id = $1`, userID,
	)
	seller := &domain.Seller{}
	err := row.Scan(&seller.ID, &seller.UserID, &seller.ShopName, &seller.Description, &seller.Phone)
	if err != nil {
		return nil, err
	}
	return seller, nil
}

func (r *SellerRepository) UpdateSeller(seller *domain.Seller) error {
	_, err := r.db.Exec(
		`UPDATE sellers SET shop_name = $1, description = $2, phone = $3 WHERE id = $4`,
		seller.ShopName, seller.Description, seller.Phone, seller.ID,
	)
	return err
}

func (r *SellerRepository) DeleteSeller(id int64) error {
	_, err := r.db.Exec(`DELETE FROM sellers WHERE id = $1`, id)
	return err
}