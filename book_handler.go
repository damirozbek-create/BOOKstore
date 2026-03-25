package handlers

import (
	"bookstore/models"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var Books = make(map[int]models.Book)
var Authors = make(map[int]models.Author)
var Categories = make(map[int]models.Category)

var BookID = 1
var AuthorID = 1
var CategoryID = 1

func SeedData() {
	// 👤 AUTHORS
	Authors[1] = models.Author{ID: 1, Name: "Leo Tolstoy"}
	Authors[2] = models.Author{ID: 2, Name: "Fyodor Dostoevsky"}
	Authors[3] = models.Author{ID: 3, Name: "George Orwell"}
	Authors[4] = models.Author{ID: 4, Name: "J.K. Rowling"}
	Authors[5] = models.Author{ID: 5, Name: "Ernest Hemingway"}

	// 🏷 CATEGORIES
	Categories[1] = models.Category{ID: 1, Name: "Fiction"}
	Categories[2] = models.Category{ID: 2, Name: "Philosophy"}
	Categories[3] = models.Category{ID: 3, Name: "Dystopia"}
	Categories[4] = models.Category{ID: 4, Name: "Fantasy"}
	Categories[5] = models.Category{ID: 5, Name: "Classic"}

	// 📚 BOOKS
	Books[1] = models.Book{
		ID: 1, Title: "War and Peace", AuthorID: 1, CategoryID: 1, Price: 15,
	}

	Books[2] = models.Book{
		ID: 2, Title: "Anna Karenina", AuthorID: 1, CategoryID: 5, Price: 12,
	}

	Books[3] = models.Book{
		ID: 3, Title: "Crime and Punishment", AuthorID: 2, CategoryID: 2, Price: 14,
	}

	Books[4] = models.Book{
		ID: 4, Title: "The Idiot", AuthorID: 2, CategoryID: 2, Price: 13,
	}

	Books[5] = models.Book{
		ID: 5, Title: "1984", AuthorID: 3, CategoryID: 3, Price: 10,
	}

	Books[6] = models.Book{
		ID: 6, Title: "Animal Farm", AuthorID: 3, CategoryID: 3, Price: 9,
	}

	Books[7] = models.Book{
		ID: 7, Title: "Harry Potter 1", AuthorID: 4, CategoryID: 4, Price: 20,
	}

	Books[8] = models.Book{
		ID: 8, Title: "Harry Potter 2", AuthorID: 4, CategoryID: 4, Price: 22,
	}

	Books[9] = models.Book{
		ID: 9, Title: "The Old Man and the Sea", AuthorID: 5, CategoryID: 1, Price: 11,
	}

	Books[10] = models.Book{
		ID: 10, Title: "A Farewell to Arms", AuthorID: 5, CategoryID: 1, Price: 13,
	}

	// 🔢 обновляем ID счётчики
	BookID = 11
	AuthorID = 6
	CategoryID = 6
}
func GetBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	category := r.URL.Query().Get("category")
	pageStr := r.URL.Query().Get("page")

	page := 1
	limit := 5

	if pageStr != "" {
		p, _ := strconv.Atoi(pageStr)
		if p > 0 {
			page = p
		}
	}

	var result []models.Book

	for _, book := range Books {
		if category != "" {
			catID, _ := strconv.Atoi(category)
			if book.CategoryID != catID {
				continue
			}
		}
		result = append(result, book)
	}

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

func CreateBook(w http.ResponseWriter, r *http.Request) {
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

	book.ID = BookID
	BookID++
	Books[book.ID] = book

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(book)
}

func GetBookByID(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	book, ok := Books[id]
	if !ok {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(book)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
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
