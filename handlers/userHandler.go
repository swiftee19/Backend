package handlers

import (
	"bone-backend/models"
	"bone-backend/services"
	"bone-backend/utilities"
	"encoding/json"
	"log"
	"net/http"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func GetUsers(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	users, err := models.GetAllUser(db)
	if err != nil {
		log.Fatalf("Error fetching users: %v", err)
		http.Error(w, "Error fetching users", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func GetUserByID(db *gorm.DB, w http.ResponseWriter, r *http.Request, userID uuid.UUID) {
	user, err := models.GetUserByID(db, userID)
	if err != nil {
		log.Fatalf("Error fetching user: %v", err)
		http.Error(w, "Error fetching user", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func Signup(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	type SignupInput struct {
		Name, Email, Password string
	}

	var input SignupInput

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	userService := services.NewUserService(db)
	user, err := userService.Signup(input.Name, input.Email, input.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	token, err := utilities.GenerateJWT(user.ID.String())
	if err != nil {
		http.Error(w, "Error generating token", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"token": token,
	})
}

func Signin(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	type SigninInput struct {
		Email, Password string
	}

	var input SigninInput

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	userService := services.NewUserService(db)
	user, err := userService.Signin(input.Email, input.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	token, err := utilities.GenerateJWT(user.ID.String())
	if err != nil {
		http.Error(w, "Error generating token", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"token": token,
	})
}

func UpdateUserLastQuestionnaireDate(db *gorm.DB, w http.ResponseWriter, r *http.Request, userID uuid.UUID) {
	userService := services.NewUserService(db)
	user, err := userService.UpdateUserLastQuestionnaireDate(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}
