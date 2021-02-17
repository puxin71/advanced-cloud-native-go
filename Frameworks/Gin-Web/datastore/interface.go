package datastore

import "github.com/puxin71/advanced-cloudnative-go/Frameworks/Gin-Web/model"

type DataStore interface {
	AllBooks() []model.Book
	Add(b model.Book)
	Delete(isbn string)
	Get(isbns ...string) []model.Book
}
