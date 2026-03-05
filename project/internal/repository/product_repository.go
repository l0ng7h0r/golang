package repository

import (
	"database/sql"

	"github.com/l0ng7h0r/internal/domain"
)

type ProductRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{
		db: db,
	}
}

func (r *ProductRepository) CreateProduct(product *domain.Product) error {
	_, err := r.db.Exec(
		"INSERT INTO public.products(seller_id, name, price, stock) VALUES ($1, $2, $3, $4)",
		product.SellerID, product.Name, product.Price, product.Stock,
	)
	return err
}

func (r *ProductRepository) GetProduct(id int64) (*domain.Product, error) {
	row := r.db.QueryRow("SELECT id, seller_id, name, price, stock FROM public.products WHERE id = $1", id)
	product := &domain.Product{}
	err := row.Scan(&product.ID, &product.SellerID, &product.Name, &product.Price, &product.Stock)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (r *ProductRepository) GetProducts() ([]*domain.Product, error) {
	var products []*domain.Product
	rows, err := r.db.Query("SELECT id, seller_id, name, price, stock FROM public.products")
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		product := &domain.Product{}
		err := rows.Scan(&product.ID, &product.SellerID, &product.Name, &product.Price, &product.Stock)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}

func (r *ProductRepository) UpdateProduct(product *domain.Product) error {
	_, err := r.db.Exec("UPDATE public.products SET name=$1, price=$2, stock=$3 WHERE id=$4",
		product.Name, product.Price, product.Stock, product.ID)
	return err
}

func (r *ProductRepository) DeleteProduct(id int64) error {
	_, err := r.db.Exec("DELETE FROM public.products WHERE id=$1", id)
	return err
}
