package main

import (
	"bookstore/handlers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// 🟢 Главная страница (чтобы не было 404)
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Bookstore API is running"))
	}).Methods("GET")

	// 📚 BOOKS
	r.HandleFunc("/books", handlers.GetBooks).Methods("GET")
	r.HandleFunc("/books", handlers.CreateBook).Methods("POST")
	r.HandleFunc("/books/{id}", handlers.GetBookByID).Methods("GET")
	r.HandleFunc("/books/{id}", handlers.UpdateBook).Methods("PUT")
	r.HandleFunc("/books/{id}", handlers.DeleteBook).Methods("DELETE")

	// 👤 AUTHORS
	r.HandleFunc("/authors", handlers.GetAuthors).Methods("GET")
	r.HandleFunc("/authors", handlers.CreateAuthor).Methods("POST")

	// 🏷 CATEGORIES
	r.HandleFunc("/categories", handlers.GetCategories).Methods("GET")
	r.HandleFunc("/categories", handlers.CreateCategory).Methods("POST")

	// 🌱 тестовые данные (ВАЖНО)
	handlers.SeedData()

	log.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
