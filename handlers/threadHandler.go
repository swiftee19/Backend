package handlers

import (
	"bone-backend/services"
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func GetAllThread(db *gorm.DB, w http.ResponseWriter, r *http.Request, userID uuid.UUID) {
	threads, err := services.GetAllThread(db, userID)
	if err != nil {
		http.Error(w, "Error fetching threads", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(threads)
}

func ThreadLike(db *gorm.DB, w http.ResponseWriter, r *http.Request, userID uuid.UUID) {
	type requestBody struct {
		ThreadID uuid.UUID
	}

	var req requestBody

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	err = services.ThreadLike(db, userID, req.ThreadID)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
