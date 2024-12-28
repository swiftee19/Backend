package services

import (
	"bone-backend/models"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService struct {
	db *gorm.DB
}

func NewUserService(db *gorm.DB) *UserService {
	return &UserService{db: db}
}

func hashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func comparePasswords(hashedPassword string, plainPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))
	return err == nil // Returns true if the passwords match
}

func (s *UserService) Signup(name, email, password string) (models.User, error) {
	// Check if user already exists

	exists, err := models.CheckUserExistsByEmail(s.db, email)
	if err != nil {
		return models.User{}, err
	}
	if exists {
		return models.User{}, errors.New("user already exists")
	}

	// Hash the password
	hashedPassword, err := hashPassword(password)
	if err != nil {
		return models.User{}, err
	}

	// Save user to DB
	newUser, err := models.CreateUser(s.db, name, email, hashedPassword)
	if err != nil {
		return models.User{}, err
	}

	return newUser, nil
}

func (s *UserService) Signin(email, password string) (models.User, error) {
	user, err := models.GetUserByEmail(s.db, email)
	if err != nil {
		return models.User{}, err
	}

	fmt.Println(user)
	fmt.Println(password)

	if !comparePasswords(user.Password, password) {
		return models.User{}, errors.New("invalid password")
	}

	return user, nil
}

func (s *UserService) UpdateUserLastQuestionnaireDate(userID uuid.UUID) (models.User, error) {
	user, err := models.GetUserByID(s.db, userID)

	if err != nil {
		return models.User{}, err
	}

	user.LastQuestionnaireDate = models.NullTime{sql.NullTime{Time: time.Now(), Valid: true}}

	return models.UpdateUser(s.db, user)
}

func (s *UserService) GetUserByID(userID uuid.UUID) (models.User, error) {
	return models.GetUserByID(s.db, userID)
}
