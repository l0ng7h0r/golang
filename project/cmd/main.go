package main

import (
	"log"

	"github.com/l0ng7h0r/internal/handler"
	"github.com/l0ng7h0r/internal/middlewares"
	"github.com/l0ng7h0r/internal/repository"
	"github.com/l0ng7h0r/internal/usecase"
	"github.com/l0ng7h0r/pkg/config"
	"github.com/l0ng7h0r/pkg/database"

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
	sellerRepo := repository.NewSellerRepository(db)

	// Auth (ต้องส่ง sellerRepo เข้าไปด้วย เพราะ register seller สร้าง seller profile ด้วย)
	authUC := usecase.NewAuthUsecase(userRepo, sellerRepo)
	authHandler := handler.NewAthHandler(authUC)

	//Products

	productRepo := repository.NewProductRepository(db)
	productUC := usecase.NewProductUsecase(productRepo)

	//Seller

	sellerUC := usecase.NewSellerUsecase(sellerRepo)
	sellerHandler := handler.NewSellerHandler(sellerUC)

	// ProductHandler ต้องใช้ SellerUsecase เพื่อหา seller_id จาก user_id
	productHandler := handler.NewProductHandler(productUC, sellerUC)

	//Order

	orderRepo := repository.NewOrderRepository(db)
	orderUC := usecase.NewOrderUsecase(orderRepo)
	orderHandler := handler.NewOrderHandler(orderUC)

	app := fiber.New()
	// Public routes (ไม่ต้อง login)
	app.Post("/register", authHandler.Register)
	app.Post("/login", authHandler.Login)

	// Protected routes (ต้อง login ทุก route)
	api := app.Group("/", middlewares.AuthMiddleware)

	// ทุก role ดูได้
	api.Get("/products", productHandler.GetProducts)
	api.Get("/products/:id", productHandler.GetProduct)

	// admin หรือ seller สร้าง/แก้ไขสินค้าได้
	api.Post("/products", middlewares.RoleMiddleware("admin", "seller"), productHandler.CreateProduct)
	api.Put("/products/:id", middlewares.RoleMiddleware("admin", "seller"), productHandler.UpdateProduct)
	api.Delete("/products/:id", middlewares.RoleMiddleware("admin"), productHandler.DeleteProduct)

	//order

	// orders - user
	api.Post("/orders", orderHandler.CreateOrders)
	api.Get("/orders/user/:id", orderHandler.GetOrdersByUserId)
	api.Get("/orders/:id", orderHandler.GetOrdersByID)
	api.Delete("/orders/:id", orderHandler.DeleteOrders)

	// orders - admin
	api.Put("/orders/:id", middlewares.RoleMiddleware("admin"), orderHandler.UpdateOrders)

	// seller routes
	api.Post("/sellers", middlewares.RoleMiddleware("seller"), sellerHandler.CreateSeller)
	api.Get("/sellers/me", middlewares.RoleMiddleware("seller"), sellerHandler.GetMySellerProfile)
	api.Get("/sellers/:id", sellerHandler.GetSeller)
	api.Put("/sellers/:id", middlewares.RoleMiddleware("seller"), sellerHandler.UpdateSeller)
	api.Delete("/sellers/:id", middlewares.RoleMiddleware("admin"), sellerHandler.DeleteSeller)

	log.Println("Server started on :8080")
	app.Listen(":8080")
}