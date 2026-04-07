package repository

import (
	"database/sql"

	"github.com/l0ng7h0r/golang/internal/domain"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) CreateUser(user *domain.User) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	defer tx.Rollback()

	var userID string
	err = tx.QueryRow(`INSERT INTO users(email, password) VALUES ($1, $2) RETURNING id`, user.Email, user.Password).Scan(&userID)
	if err != nil {
		return err
	}

	for _, role := range user.Roles {
		_, err := tx.Exec(`INSERT INTO user_roles (user_id, role_id) SELECT $1, id FROM roles WHERE name = $2`, userID, role)
		if err != nil {
			return err
		}
	}
	err = tx.Commit()
	return err
}

func (r *UserRepository) FindByEmail(email string) (*domain.User, error) {
	row := r.db.QueryRow(`SELECT id, email, password FROM users WHERE email=$1`, email)
	var u domain.User
	err := row.Scan(&u.ID, &u.Email, &u.Password)
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func (r *UserRepository) SaveRefreshToken(userID, token string) error {
	_, err := r.db.Exec(`INSERT INTO refresh_tokens(user_id, token, expires_at) VALUES ($1,$2,NOW() + INTERVAL '7 days')`, userID, token)
	return err
}

func (r *UserRepository) ValidateRefreshToken(token string) (string, error) {
	row := r.db.QueryRow(`SELECT user_id FROM refresh_tokens WHERE token=$1 AND expires_at > NOW()`, token)
	var userID string
	err := row.Scan(&userID)
	return userID, err
}

func (r *UserRepository) DeleteRefreshToken(token string) error {
	_, err := r.db.Exec("DELETE FROM refresh_tokens WHERE token=$1", token)
	return err
}