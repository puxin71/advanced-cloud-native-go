package service

import (
	"github.com/puxin71/advanced-cloudnative-go/Frameworks/Gin-Web/datastore"
	"github.com/puxin71/advanced-cloudnative-go/Frameworks/Gin-Web/model"
)

// Collections of books stored in memory
type InMemBookService struct {
	store *datastore.InMemStore
}

// Initialize a memBookService with a collection of books
func NewInMemBookService(store *datastore.InMemStore) *InMemBookService {
	return &InMemBookService{store: store}
}

// Return the slice of books stored in memory. Note that the
// returned book slice can be empty
func (s *InMemBookService) AllBooks() []model.Book {
	return s.store.AllBooks()
}

// Add an book to memory
func (s *InMemBookService) AddBook(b model.Book) (string, error) {
	s.store.Add(b)
	return b.ISBN, nil
}

// Get a book by its ISBN from memory
func (s *InMemBookService) GetBook(isbn string) (model.Book, error) {
	books := s.store.Get(isbn)
	if len(books) == 0 {
		return model.Book{}, model.NewNotFound("book", isbn)
	}
	return books[0], nil
}

// Delete a book from memory
func (s *InMemBookService) DeleteBook(isbn string) error {
	s.store.Delete(isbn)
	return nil
}

// Get one or more books from memory
func (s *InMemBookService) GetBooks(isbns ...string) []model.Book {
	return s.store.Get(isbns...)
}
