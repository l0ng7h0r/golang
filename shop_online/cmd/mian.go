package main

import (
	"log"

	"github.com/gofiber/fiber/v3"

	"github.com/l0ng7h0r/golang/internal/handler"
	"github.com/l0ng7h0r/golang/internal/repository"
	"github.com/l0ng7h0r/golang/internal/usecase"
	"github.com/l0ng7h0r/golang/pkg/config"
	"github.com/l0ng7h0r/golang/pkg/database"
	"github.com/l0ng7h0r/golang/internal/middleware"
)

func main() {
	cfg := config.Load()

	db, err := database.NewPostgres(cfg.DBDsn)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}   

	userRepo := repository.NewUserRepository(db)
	authUsecase := usecase.NewAuthUsecase(userRepo)
	authHandler := handler.NewAuthHandler(authUsecase)
	authMiddleware := middleware.NewAuthMiddleware(authUsecase)

	app := fiber.New()

	api := app.Group("/api")

	//Public routes
	api.Post("/register", authHandler.Register)
	api.Post("/login", authHandler.Login)
	api.Post("/refresh", authHandler.Refresh)
	
	//Admin routes
	admin := api.Group("/admin")
	admin.Use(authMiddleware.Auth)
	admin.Use(authMiddleware.RequireRole("admin"))
	admin.Post("/create-user", authHandler.CreateUser)
	admin.Get("/users", authHandler.GetAllUsers)
	admin.Get("/users/:id", authHandler.GetUserByID)
	admin.Delete("/users/:id", authHandler.DeleteUser)

	app.Listen(":" + cfg.AppPort)
}