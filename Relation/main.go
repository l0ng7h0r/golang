package main
import (
	"fmt"
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
	"log"
	"os"
	"time"
	"gorm.io/gorm/logger"
)


const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "long1122"
	dbname   = "test_go"
)

type Book struct{
	gorm.Model
	Name string `json:"name"`
	Author string `json:"author"`
	Description string `json:"description"`
	PublisherID uint
	publisher Publisher
	Authors []Author `gorm:"many2many:book_authors;"`
}

type Publisher struct{
	gorm.Model
	Details string
	Name string 
}

type Author struct{
	gorm.Model
	Name string `json:"name"`
	Book []Book `gorm:"many2many:book_authors;"`
}

type BookAuthor struct{
	Author Author
	BookID uint
	AuthorID uint
	Book Book
}

func main() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", host, user, password, dbname, port)

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel: logger.Info,
			Colorful: true,
		},
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})

	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&Book{}, &Publisher{}, &Author{}, &BookAuthor{})

	publisher := Publisher{
		Name: "long",
		Details: "bht",
	}

	_ = createPublisher(db, &publisher)

	author1 := Author{
		Name: "author",
	}

	_ = createAuthor(db, &author1)

	author2 := Author{
		Name: "author",
	}

	_ = createAuthor(db, &author2)

	book := Book{
		Name: "aaaa",
		Author: "author1",
		Description: "description",
		PublisherID: publisher.ID,
	}
	_ = createBookWithAuthor(db, &book, []uint{author1.ID, author2.ID})

}

func createPublisher(db *gorm.DB, publisher *Publisher) error {
  result := db.Create(publisher)
  if result.Error != nil {
    return result.Error
  }
  return nil
}

func createAuthor(db *gorm.DB, author *Author) error {
  result := db.Create(author)
  if result.Error != nil {
    return result.Error
  }
  return nil
}

func createBookWithAuthor(db *gorm.DB, book *Book, authorIDs []uint) error {
  // First, create the book
  if err := db.Create(book).Error; err != nil {
    return err
  }

  return nil
}