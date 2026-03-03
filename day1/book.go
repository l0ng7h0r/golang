package main

import (
	"strconv"

	"github.com/gofiber/fiber/v3"
)


// getBooks godoc
// @Summary Get all books
// @Tags Books
// @Security ApiKeyAuth
// @Produce json
// @Success 200 {array} Book
// @Router /get-books [get]
func getBooks(c fiber.Ctx) error {
	return c.JSON(books)
}

// getBooksById godoc
// @Summary Get book by ID
// @Tags Books
// @Produce json
// @Param id path int true "Book ID"
// @Success 200 {object} Book
// @Security ApiKeyAuth
// @Router /get-books/{id} [get]
func getBooksById(c fiber.Ctx) error {
	bookId, err := strconv.Atoi(c.Params("id"));

	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error());
	}

	for _, book := range books {
		if book.ID == bookId {
			return c.JSON(book);
		}
	}
	return c.SendStatus(fiber.StatusNotFound);
}

// createBook godoc
// @Summary Create a new book
// @Tags Books
// @Accept json
// @Produce json
// @Param book body Book true "New Book"
// @Success 200 {object} Book
// @Security ApiKeyAuth
// @Router /create [post]
func createBook(c fiber.Ctx) error {
	book := new(Book)
	if err := c.Bind().Body(book); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error());
	}
	books = append(books, *book)
	return c.Status(fiber.StatusCreated).JSON(book)
}


// updateBook godoc
// @Summary Update book
// @Tags Books
// @Accept json
// @Produce json
// @Param id path int true "Book ID"
// @Param book body Book true "Updated Book"
// @Success 200 {object} Book
// @Security ApiKeyAuth
// @Router /update/{id} [put]
func updateBook(c fiber.Ctx) error {
	bookId, err := strconv.Atoi(c.Params("id"));

	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error());
	}

	bookUpdated := new(Book)

	if err := c.Bind().Body(bookUpdated); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error());
	}

	for i, book := range books {
		if book.ID == bookId {
			books[i].Title = bookUpdated.Title
			books[i].Author = bookUpdated.Author
			return c.JSON(books[i])
		}
	}

	return c.SendStatus(fiber.StatusNotFound)
}

// deleteBook godoc
// @Summary Delete book
// @Tags Books
// @Param id path int true "Book ID"
// @Success 200 {string} string
// @Security ApiKeyAuth
// @Router /delete/{id} [delete]

func deleteBook(c fiber.Ctx) error {
	bookId, err := strconv.Atoi(c.Params("id"));

	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error());
	}

	for i, book := range books {
		if book.ID == bookId {
			books = append(books[:i], books[i+1:]...)
			return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Book deleted successfully", "book": book})
		}
	}

	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Book not found" })
}