package repository

import (
	"log"
	"database/sql"
	"github.com/l0ng7h0r/internal/domain"
)

type UserRepository struct {
	db *sql.DB
}


func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) CreateUser(user *domain.User) error {
    tx, err := r.db.Begin()
    if err != nil {
        return err
    }
    defer tx.Rollback()

    var userID int64
    err = tx.QueryRow(
        `INSERT INTO users (email, password) VALUES ($1, $2) RETURNING id`,
        user.Email, user.Password,
    ).Scan(&userID)
    if err != nil {
        log.Println("Insert user error:", err)
        return err
    }
    log.Println("Inserted userID:", userID)

    for _, role := range user.Roles {
        log.Println("Inserting role:", role)
        _, err := tx.Exec(
            `INSERT INTO user_roles (user_id, role_id) SELECT $1, id FROM roles WHERE name = $2`,
            userID, role,
        )
        if err != nil {
            log.Println("Insert role error:", err)
            return err
        }
    }

    err = tx.Commit()
    log.Println("Commit error:", err)
    return err
}

func (r *UserRepository) FindUserByEmail(email string) (*domain.User, error) {
    row := r.db.QueryRow(`SELECT id, email, password FROM users WHERE email = $1`, email)

    var user domain.User
    err := row.Scan(&user.ID, &user.Email, &user.Password)
    if err != nil {
        return nil, err
    }

    // แก้ JION → JOIN และ FROM roles ให้ถูกต้อง
    rows, err := r.db.Query(
        `SELECT roles.name FROM roles JOIN user_roles ON roles.id = user_roles.role_id WHERE user_roles.user_id = $1`,
        user.ID,
    )
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    for rows.Next() {
        var role string
        rows.Scan(&role)
        user.Roles = append(user.Roles, role)
    }

    return &user, nil
}