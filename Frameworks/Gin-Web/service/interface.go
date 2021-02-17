package service

import "github.com/puxin71/advanced-cloudnative-go/Frameworks/Gin-Web/model"

// BookService defines methods that the HTTP handler expects
type BookService interface {
	// Get all books from the data store
	AllBooks() []model.Book
	// Add a book to the data store
	AddBook(b model.Book) (string, error)
	// Get a book by isbn from the data store
	GetBook(isbn string) (model.Book, error)
	// Delete a book by isbn from the data store
	DeleteBook(isbn string) error
	// Get a slice of books by a list of isbns from the data store
	GetBooks(isbn ...string) []model.Book
}
