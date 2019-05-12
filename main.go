// You can edit this code!
// Click here and start typing.
package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Year   string `json:"year"`
}

var books []Book

func main() {
	router := mux.NewRouter()

	books = append(books, Book{ID: 1, Title: "abc", Author: "thinh", Year: "2010"},
		Book{ID: 2, Title: "abc", Author: "thinh", Year: "2010"})
	router.HandleFunc("/books", getBooks).Methods("Get")
	router.HandleFunc("/books/{id}", getBook).Methods("Get")
	router.HandleFunc("/books", addbook).Methods("POST")
	router.HandleFunc("/books", updatebook).Methods("PUT")
	router.HandleFunc("/books/{id}", removeBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", router))
}

func getBooks(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(books)
	log.Println("get all book")
}
func getBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	//log.println(reflect.TypeOf(params["id"]))
	i, _ := strconv.Atoi(params["id"])
	for _, book := range books {
		if book.ID == i {
			json.NewEncoder(w).Encode(&book)
		}
	}
}

func addbook(w http.ResponseWriter, r *http.Request) {
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	books = append(books, book)
	json.NewEncoder(w).Encode(books)
}

func removeBook(w http.ResponseWriter, r *http.Request) {
	log.Println("get all book")
}

func updatebook(w http.ResponseWriter, r *http.Request) {
	log.Println("get all book")
}
