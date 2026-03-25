package handlers

import (
	"bookstore/models"
	"encoding/json"
	"net/http"
)

func GetAuthors(w http.ResponseWriter, r *http.Request) {
	var list []models.Author

	for _, a := range Authors {
		list = append(list, a)
	}

	json.NewEncoder(w).Encode(list)
}

func CreateAuthor(w http.ResponseWriter, r *http.Request) {
	var author models.Author

	json.NewDecoder(r.Body).Decode(&author)

	if author.Name == "" {
		http.Error(w, "Name required", http.StatusBadRequest)
		return
	}

	author.ID = AuthorID
	AuthorID++
	Authors[author.ID] = author

	json.NewEncoder(w).Encode(author)
}
