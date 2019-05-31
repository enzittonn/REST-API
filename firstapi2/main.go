package main

import (
	"encoding/json"
	"log"
	"net/http"
	"fmt"
	"math/rand"
	"strconv"

	"github.com/gorilla/mux"
)


type Book struct {
	ID          string `json:"id"`
	Isbn 	    string `json:"isbn"`
	Title       string `json:"title"`
	Author 	    *Author `json:"author"`
}

type Author struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

var books []Book

func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Book{})
}

func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(9999))
	books = append(books, book)
	json.NewEncoder(w).Encode(book)
}

func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			var book Book
			_ = json.NewDecoder(r.Body).Decode(&book)
			book.ID = params["id"]
			books = append(books, book)
			json.NewEncoder(w).Encode(book)
			return
		}
	}
	json.NewEncoder(w).Encode(books)
}

func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range books {
		if item.ID == params["id"]{
			books = append(books[:index], books[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(books)
}

func welcome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to my shop")
}


func main(){
	
	books = append(books, Book{ID: "1", Isbn: "978-1-449-31050-9", Title: "REST API - Design Book", Author: &Author {FirstName: "Mark", LastName: "Massé"}})
	
	books = append(books, Book{ID: "2", Isbn: "978-1-4842-2691-9", Title: "Network Programming with Go", Author: &Author {FirstName: "Jan", LastName: "Newmarch"}})
	
	books = append(books, Book{ID: "3", Isbn: "978-1-491-919712-2", Title: "Docker Cookbook", Author: &Author {FirstName: "Sébastien", LastName: "Goasguen"}})
	
	books = append(books, Book{ID: "4", Isbn: "978-1-491-956250-4", Title: "Microservice Architecture", Author: &Author {FirstName: "Irakli", LastName: "Nadareishvili"}})


	
	router := mux.NewRouter()


	router.HandleFunc("/", welcome)
	router.HandleFunc("/books", getBooks).Methods("GET")
	router.HandleFunc("/book/{id}", getBook).Methods("GET")
	router.HandleFunc("/book", createBook).Methods("POST")
	router.HandleFunc("/book/{id}", updateBook).Methods("PUT")
	router.HandleFunc("/book/{id}", deleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", router))
}
