package auth

import (
	"errors"
	"test/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Register(db *gorm.DB, email, password string) error {
	var user models.User

	if err := db.Where("email = ?", email). First(&user).Error; err == nil {
		return errors.New("email already exists")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return err
	}

	user = models.User{
		Email:    email,
		Password: string(hash),
	}

	return db.Create(&user).Error
}
func Login(db *gorm.DB, email, password string) (*models.User, error) {
	var user models.User

	if err := db.Where("email = ?", email).First(&user).Error; err != nil {
		return  nil, errors.New("invalid credentails") 
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return  nil, errors.New("invalid credentails")
	}

	return &user, nil
}