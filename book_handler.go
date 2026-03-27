package handlers

import (
	"bookstore/models"
	"encoding/json"
	"net/http"
	"sort"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

var Books = make(map[int]models.Book)
var Authors = make(map[int]models.Author)
var Categories = make(map[int]models.Category)

var BookID = 1
var AuthorID = 1
var CategoryID = 1

// 🌱 SEED DATA
func SeedData() {
	Authors[1] = models.Author{ID: 1, Name: "Leo Tolstoy"}
	Authors[2] = models.Author{ID: 2, Name: "Fyodor Dostoevsky"}

	Categories[1] = models.Category{ID: 1, Name: "Fiction"}
	Categories[2] = models.Category{ID: 2, Name: "Philosophy"}

	Books[1] = models.Book{ID: 1, Title: "War and Peace", AuthorID: 1, CategoryID: 1, Price: 15}
	Books[2] = models.Book{ID: 2, Title: "Crime and Punishment", AuthorID: 2, CategoryID: 2, Price: 14}

	BookID = 3
	AuthorID = 3
	CategoryID = 3
}

// 📚 GET /books
func GetBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	categoryName := r.URL.Query().Get("category")
	pageStr := r.URL.Query().Get("page")

	page := 1
	limit := 5

	if pageStr != "" {
		p, _ := strconv.Atoi(pageStr)
		if p > 0 {
			page = p
		}
	}

	var categoryID int
	if categoryName != "" {
		for _, c := range Categories {
			if strings.EqualFold(c.Name, categoryName) {
				categoryID = c.ID
				break
			}
		}
	}

	var result []models.Book

	for _, book := range Books {
		if categoryName != "" && book.CategoryID != categoryID {
			continue
		}
		result = append(result, book)
	}

	// сортировка по цене
	sort.Slice(result, func(i, j int) bool {
		return result[i].Price < result[j].Price
	})

	start := (page - 1) * limit
	end := start + limit

	if start > len(result) {
		start = len(result)
	}
	if end > len(result) {
		end = len(result)
	}

	json.NewEncoder(w).Encode(result[start:end])
}

// ➕ CREATE BOOK
func CreateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var book models.Book

	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if book.Title == "" || book.Price <= 0 {
		http.Error(w, "Title and price required", http.StatusBadRequest)
		return
	}

	if _, ok := Authors[book.AuthorID]; !ok {
		http.Error(w, "Author not found", http.StatusBadRequest)
		return
	}

	if _, ok := Categories[book.CategoryID]; !ok {
		http.Error(w, "Category not found", http.StatusBadRequest)
		return
	}

	book.ID = BookID
	BookID++
	Books[book.ID] = book

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(book)
}

// 🔍 GET BY ID
func GetBookByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	book, ok := Books[id]
	if !ok {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(book)
}

// ✏️ UPDATE
func UpdateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	_, ok := Books[id]
	if !ok {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}

	var book models.Book
	json.NewDecoder(r.Body).Decode(&book)

	if book.Title == "" || book.Price <= 0 {
		http.Error(w, "Invalid data", http.StatusBadRequest)
		return
	}

	book.ID = id
	Books[id] = book

	json.NewEncoder(w).Encode(book)
}

// ❌ DELETE
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	_, ok := Books[id]
	if !ok {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}

	delete(Books, id)
	w.WriteHeader(http.StatusNoContent)
}
