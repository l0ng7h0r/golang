package main

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"

	"github.com/gofiber/fiber/v3"

	_ "github.com/lib/pq"
)

const (
	host = "localhost"
	port = 5432
	database = "postgres"
	dbname = "gotest"
	password = "long1122"
)

var db *sql.DB

type Product struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Price int `json:"price"`
}
func main() {
	psqlInfo := fmt.Sprintf(
    "host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
    host, port, database, password, dbname,
	)
	dbs, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		log.Fatal(err)
	}

	db = dbs
	err = db.Ping()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to database")

	app := fiber.New()
	app.Get("/products", getAllProductsHandler)
	app.Get("/get-product/:id", getProductByid)
	app.Post("/createProduct", createProductHandler)
	app.Put("/updateProduct/:id", updateProductHandler)
	app.Delete("/deleteProduct/:id", deleteProductHandler)

	app.Listen(":8080")

	
	product, err := getAllProducts()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Products: ", product)
}


func getProductByid(c fiber.Ctx) error {
	productId, err := strconv.Atoi(c.Params("id"))
	
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	product, err := getProduct(productId)

	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	return c.JSON(product)
}


func createProductHandler(c fiber.Ctx) error {
	p := new(Product)

	err := c.Bind().Body(p)

	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	err = createProduct(p)

	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	return c.JSON(p)
	
}

func updateProductHandler(c fiber.Ctx) error {
	productId, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	p := new(Product)

	err = c.Bind().Body(p)

	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	product, err := updateProduct(productId, p)

	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	return c.JSON(product)
}

func deleteProductHandler(c fiber.Ctx) error {
	productId, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	err = deleteProduct(productId)

	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	return c.SendStatus(fiber.StatusOK)
}

func getAllProductsHandler(c fiber.Ctx) error {
	products, err := getAllProducts()

	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	return c.JSON(products)
}