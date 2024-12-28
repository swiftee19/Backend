package handlers

import (
	"encoding/json"
	"net/http"

	"gorm.io/gorm"

	"bone-backend/services"
)

func GetAllArticle(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	articles, err := services.GetAllArticle(db)
	if err != nil {
		http.Error(w, "Error fetching articles", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(articles)
}

func GetLatestArticle(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	article, err := services.GetLatestArticle(db)
	if err != nil {
		http.Error(w, "Error fetching latest article", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(article)
}
