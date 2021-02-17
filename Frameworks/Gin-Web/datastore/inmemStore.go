package datastore

import (
	"sync"

	"github.com/puxin71/advanced-cloudnative-go/Frameworks/Gin-Web/model"
)

// Thread-safe collection of books stored in a map using the ISBN as the key
type InMemStore struct {
	books sync.Map
}

// Create a memStore to keep the books by their ISBN number
func NewInMemStore() *InMemStore {
	return &InMemStore{}
}

// Return the slice of books stored in memory. Note that the
// returned book slice can be empty
func (s *InMemStore) AllBooks() []model.Book {
	books := make([]model.Book, 0)
	s.books.Range(func(key, value interface{}) bool {
		books = append(books, value.(model.Book))
		return true
	})
	return books
}

// Add an book to memory
func (s *InMemStore) Add(b model.Book) {
	s.books.Store(b.ISBN, b)
}

// Delete a book from memory
func (s *InMemStore) Delete(isbn string) {
	s.books.Delete(isbn)
}

// Get one or more books from memory. Note that the returned slice
// of book can be empty
func (s *InMemStore) Get(isbns ...string) []model.Book {
	books := make([]model.Book, 0)
	for _, v := range isbns {
		if value, ok := s.books.Load(v); ok {
			books = append(books, value.(model.Book))
		}
	}
	return books
}
