package main

import (
	"bookstore/handlers"
	"bookstore/models"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	// 📚 Routes
	router.HandleFunc("/books", handlers.GetBooks).Methods("GET")
	router.HandleFunc("/books", handlers.CreateBook).Methods("POST")
	router.HandleFunc("/books/{id}", handlers.GetBooks).Methods("GET")
	router.HandleFunc("/books/{id}", handlers.UpdateBook).Methods("PUT")
	router.HandleFunc("/books/{id}", handlers.DeleteBook).Methods("DELETE")

	router.HandleFunc("/authors", handlers.GetAuthors).Methods("GET")
	router.HandleFunc("/authors", handlers.CreateAuthor).Methods("POST")

	router.HandleFunc("/categories", handlers.GetCategories).Methods("GET")
	router.HandleFunc("/categories", handlers.CreateCategory).Methods("POST")

	// 📌 SEED DATA (стартовые данные)
	handlers.Authors[1] = models.Author{ID: 1, Name: "George Orwell"}
	handlers.Authors[2] = models.Author{ID: 2, Name: "Robert Martin"}
	handlers.Authors[3] = models.Author{ID: 3, Name: "Elon Musk"}

	handlers.Categories[1] = models.Category{ID: 1, Name: "Fiction"}
	handlers.Categories[2] = models.Category{ID: 2, Name: "Technology"}
	handlers.Categories[3] = models.Category{ID: 3, Name: "Business"}

	handlers.Books[1] = models.Book{
		ID: 1, Title: "1984", AuthorID: 1, CategoryID: 1, Price: 12.99,
	}
	handlers.Books[2] = models.Book{
		ID: 2, Title: "Clean Code", AuthorID: 2, CategoryID: 2, Price: 25.50,
	}
	handlers.Books[3] = models.Book{
		ID: 3, Title: "Tesla Story", AuthorID: 3, CategoryID: 3, Price: 18.00,
	}

	log.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
