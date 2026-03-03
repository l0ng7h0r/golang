package main

import (
	"log"
	"github.com/l0ng7h0r/pkg/config"
	"github.com/l0ng7h0r/pkg/database"
	"github.com/l0ng7h0r/internal/handler"
	"github.com/l0ng7h0r/internal/repository"
	"github.com/l0ng7h0r/internal/usecase"

	"github.com/gofiber/fiber/v3"
)

func main() {
	env := config.LoadEnv()

	db, err := database.Connect(env)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	userRepo := repository.NewUserRepository(db)
	authUC := usecase.NewAuthUsecase(userRepo)
	authHandler := handler.NewAthHandler(authUC)

	app := fiber.New()

	app.Post("/register", authHandler.Register)
	app.Post("/login", authHandler.Login)

	log.Println("Server started on :8080")
	app.Listen(":8080")
}