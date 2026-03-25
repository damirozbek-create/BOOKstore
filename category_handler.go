package handlers

import (
	"bookstore/models"
	"encoding/json"
	"net/http"
)

func GetCategories(w http.ResponseWriter, r *http.Request) {
	var list []models.Category

	for _, c := range Categories {
		list = append(list, c)
	}

	json.NewEncoder(w).Encode(list)
}

func CreateCategory(w http.ResponseWriter, r *http.Request) {
	var category models.Category

	json.NewDecoder(r.Body).Decode(&category)

	if category.Name == "" {
		http.Error(w, "Name required", http.StatusBadRequest)
		return
	}

	category.ID = CategoryID
	CategoryID++
	Categories[category.ID] = category

	json.NewEncoder(w).Encode(category)
}
