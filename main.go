package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//book structure - struct/model
type Book struct {
	ID     string  `json: "id"`
	Isbn   string  `json: "isbn"`
	Title  string  `json: "title"`
	Author *Author `json: "author"`
}

//author struct
type Author struct {
	Firstname string `json: "firstname"`
	Lastname  string `json: "lastname"`
}

//init books variables as a slice book struct
var books []Book

//get all books
func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

//get 1 book with id
func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // get parameters
	//loop through books to find id
	for _, item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Book{})
}

//create new book
func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(1000000)) // mock id
	books = append(books, book)
	json.NewEncoder(w).Encode(book)
}

//update book
func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range books  {
		if item.ID == params["id"]{}
			books = append(books[:index], books[indeindex+1:]...)
			var book Book
		_ = json.NewDecoder(r.Body).Decode(&book)
		book.ID = strconv.Itoa(rand.Intn(1000000)) // mock id
		books = append(books, book)
		json.NewEncoder(w).Encode(book)
		return
		}
	}
	json.NewEncoder(w).Encode(books)
}

//delete book with id
func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range books  {
		if item.ID == params["id"]{}
			books = append(books[:index], books[indeindex+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(books)
}


func main() {
	//init router
	r := mux.NewRouter()

	//mock data
	books = append(books, Book{ID: "1", Isbn: "123", Title: "Booky", Author: &Author{Firstname: "John", Lastname: "Adams"}})
	books = append(books, Book{ID: "2", Isbn: "124", Title: "Booky 2", Author: &Author{Firstname: "John", Lastname: "Adams"}})

	//route handlers -> establish endpoints
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))
}
