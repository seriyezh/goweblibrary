package store

import "github.com/seriyezh/goweblibrary/webapi/entities"

var Books []entities.Book

func init() {
	Books = append(Books, entities.Book{ID: "1", Title: "Война и Мир", Author: &entities.Author{Firstname: "Лев", Lastname: "Толстой"}})
	Books = append(Books, entities.Book{ID: "2", Title: "Преступление и наказание", Author: &entities.Author{Firstname: "Фёдор", Lastname: "Достоевский"}})
}
