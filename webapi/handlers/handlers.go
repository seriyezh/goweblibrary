package handlers

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/seriyezh/goweblibrary/webapi/entities"
	"github.com/seriyezh/goweblibrary/webapi/store"

	"github.com/gorilla/mux"
)

func GetBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(store.Books)
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range store.Books {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&entities.Book{})
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book entities.Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(1000000))
	store.Books = append(store.Books, book)
	json.NewEncoder(w).Encode(book)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range store.Books {
		if item.ID == params["id"] {
			store.Books = append(store.Books[:index], store.Books[index+1:]...)
			var book entities.Book
			_ = json.NewDecoder(r.Body).Decode(&book)
			book.ID = params["id"]
			store.Books = append(store.Books, book)
			json.NewEncoder(w).Encode(book)
			return
		}
	}
	json.NewEncoder(w).Encode(store.Books)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range store.Books {
		if item.ID == params["id"] {
			store.Books = append(store.Books[:index], store.Books[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(store.Books)
}
