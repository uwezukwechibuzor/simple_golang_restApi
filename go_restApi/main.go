package main

import (
	"encoding/json"
	"math/rand"
	"strconv"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//book structs(Model)
type Book struct {
	ID     string  `json:"id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"id"`
}

//Author Struct
type Author struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

//Init books var as a slice book struct
var books []Book

//get all books
func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

//get a sungle book
func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
    params := mux.Vars(r) //get params id, then loop through books and find with id
	for _, item := range books {
		if item.ID == params["id"]{
           json.NewEncoder(w).Encode(item)
		   return
		}
	}

	json.NewEncoder(w).Encode(&Book{}) 
}

//for creating books
func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
   var book Book
   _= json.NewDecoder(r.Body).Decode(&book)
   book.ID = strconv.Itoa(rand.Intn(1000000))
   books = append(books, book)
   json.NewEncoder(w).Encode(book)
}
 
//updating book
func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
    params := mux.Vars(r)
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			var book Book

			_= json.NewDecoder(r.Body).Decode(&book)
			book.ID = params["id"]
			books = append(books, book)
			json.NewEncoder(w).Encode(book)
			return
		}
	}
}

//delete book
func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
    params := mux.Vars(r)
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]... )
			break
		}
	}
	json.NewEncoder(w).Encode(books)
}

func main() {
	//init Router
	r := mux. NewRouter()

	//mock data - @todo - implement DB
	books = append(books, Book{ID: "1", Isbn: "123", Title: "Blockchain Technology", Author: &Author{FirstName: "Daniel", LastName: "Chibuzor"}})

	books = append(books, Book{ID: "1", Isbn: "123", Title: "Defi", Author: &Author{FirstName: "Daniel", LastName: "Chibuzor"}})

	books = append(books, Book{ID: "1", Isbn: "123", Title: "Power of Web3 Technology", Author: &Author{FirstName: "Daniel", LastName: "Chibuzor"}})

	//Route handlers/Endpoints
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("Delete")

	//run the server
	log.Fatal(http.ListenAndServe(": 4000", r))
}
