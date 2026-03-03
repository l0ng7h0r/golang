package main

import (
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/template/html/v2"
	"github.com/joho/godotenv"

	jwtware "github.com/gofiber/contrib/v3/jwt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/gofiber/swagger/v2"
	_ "github.com/long/fiber-test/docs"
)


type Book struct {
	ID int `json:"id"`
	Title string `json:"title"`
	Author string `json:"author"`
}

type User struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

var userMember  = User{
	Email: "long123@gmail.com", Password: "123456",
}

var books []Book

// @title Book API
// @version 1.0
// @description This is a sample server for a book API.
// @host localhost:8080
// @BasePath /
// @schemes http
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func checkMiddleware(c fiber.Ctx) error {
	return c.Next()
}


func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	engine := html.New("./views", ".html")

    app := fiber.New(fiber.Config{
        Views: engine,
    })

	// Swagger route
	app.Get("/swagger/*", swagger.HandlerDefault)


	books = append(books, Book{ID: 1, Title: "Book 1", Author: "Author 1"})
	books = append(books, Book{ID: 2, Title: "Book 2", Author: "Author 2"})
	books = append(books, Book{ID: 3, Title: "Book 3", Author: "Author 3"})


	app.Post("/login", login)

	app.Use(checkMiddleware)


	// TODO: Fix JWT middleware - need to find Fiber v3 compatible JWT package
	app.Use(jwtware.New(jwtware.Config{ 
		SigningKey: jwtware.SigningKey{ 
			Key: []byte(os.Getenv("JWT_SECRET")), 
			}, 
	}))


	app.Get("/get-books", getBooks);
	app.Get("/get-books/:id", getBooksById)
	app.Post("/create", createBook)
	app.Put("/update/:id", updateBook)
	app.Delete("/delete/:id", deleteBook)

	app.Post("/upload", uploadFile)
	app.Get("/test-html", testHTML)
	
	app.Get("/get-env", getENV);
	
	app.Listen(":8080")
}


// uploadFile godoc
// @Summary Upload image file
// @Tags File
// @Accept multipart/form-data
// @Produce json
// @Param image formData file true "Image file"
// @Success 200 {string} string
// @Router /upload [post]

func uploadFile(c fiber.Ctx) error {

	file, err := c.FormFile("image")

	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	err = c.SaveFile(file, "./upload/" + file.Filename)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.SendString("File upload complete!")
}

// testHTML
// @Summary Render example HTML template
// @Tags Html
// @Success 200
// @Router /test-html [get]

func testHTML(c fiber.Ctx) error {
	return c.Render("index", fiber.Map{
		"Title": "Hello, Word!",
	})
}

// getENV
// @Summary Get ENV variable (debug)
// @Tags System
// @Success 200 {object} map[string]string
// @Router /get-env [get]
func getENV(c fiber.Ctx) error {
	return c.JSON(
		fiber.Map{
			"SECRET": os.Getenv("JWT_SECRET"),
		},
	)
}

// login godoc
// @Summary Login user
// @Description Authenticates user and returns a JWT token
// @Tags Auth
// @Accept json
// @Produce json
// @Param credentials body User true "User login"
// @Success 200 {object} map[string]string
// @Router /login [post]
func login (c fiber.Ctx) error {
	user := new(User)

	if err := c.Bind().Body(user); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	if user.Email != userMember.Email || user.Password != userMember.Password {
		return fiber.ErrUnauthorized
	}

	// Generate JWT token
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["Email"] = userMember.Email
	claims["role"] = "admin"
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	t, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{
		"token": t,
	})
}
	

