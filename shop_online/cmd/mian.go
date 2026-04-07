package main

import (
	"log"

	"github.com/gofiber/fiber/v3"

	"github.com/l0ng7h0r/golang/internal/handler"
	"github.com/l0ng7h0r/golang/internal/repository"
	"github.com/l0ng7h0r/golang/internal/usecase"
	"github.com/l0ng7h0r/golang/pkg/config"
	"github.com/l0ng7h0r/golang/pkg/database"
)

func main() {
	cfg := config.Load()

	db, err := database.NewPostgres(cfg.DBDsn)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}   

	repo := repository.NewUserRepository(db)
	usecase := usecase.NewAuthUsecase(repo)
	handler := handler.NewAuthHandler(usecase)

	app := fiber.New()

	app.Post("/register", handler.Register)
	app.Post("/login", handler.Login)
	app.Post("/refresh", handler.Refresh)

	app.Listen(":" + cfg.AppPort)
}