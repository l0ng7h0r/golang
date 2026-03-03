package database

import (
	"database/sql"
	"fmt"
	"github.com/l0ng7h0r/pkg/config"
	_ "github.com/lib/pq"
)

func Connect(env *config.Config) (*sql.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", env.DB_HOST, env.DB_PORT, env.DB_USER, env.DB_PASS, env.DB_NAME)
	return sql.Open("postgres", dsn)
}