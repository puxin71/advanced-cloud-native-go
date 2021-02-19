package handler

import (
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
	"github.com/puxin71/advanced-cloudnative-go/Frameworks/Gin-Web/handler/middleware"
	"github.com/puxin71/advanced-cloudnative-go/Frameworks/Gin-Web/model"
	"github.com/puxin71/advanced-cloudnative-go/Frameworks/Gin-Web/service"
)

// Handler struct holds required services and configs for handler to function
type Handler struct {
	Engine      *gin.Engine
	Config      *Config
	BookService service.BookService
}

func NewHandler(engine *gin.Engine, config *Config, service service.BookService) *Handler {
	return &Handler{
		Engine:      engine,
		Config:      config,
		BookService: service,
	}
}

// Provision HTTP endpoints
func (h *Handler) CreateEndpoints() {
	// Ping is used by ks8 to verify if the POD is ready and healthy
	h.Engine.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
	g := h.Engine.Group(h.Config.BaseURL)

	// Set the timeout period in seconds for each endpoint
	g.Use(middleware.Timeout(h.Config.TimeoutDuration, model.NewServiceUnavailable()))
	g.GET("/books", func(c *gin.Context) {
		var isbns []string
		var books []model.Book
		var ok bool

		if isbns, ok = c.GetQueryArray("isbns"); !ok || len(isbns) == 0 {
			books = h.BookService.AllBooks()
		} else {
			books = h.BookService.GetBooks(isbns...)
		}

		if len(books) == 0 {
			c.Status(http.StatusNotFound)
			return
		}
		c.JSON(http.StatusOK, books)
	})
	g.POST("/books", func(c *gin.Context) {
		var book model.Book
		if c.BindJSON(&book) == nil {
			isbn, err := h.BookService.AddBook(book)
			if err == nil {
				c.Header("location", g.BasePath()+"/books/"+isbn)
				c.Status(http.StatusCreated)
			} else {
				log.WithFields(log.Fields{
					"book":  book,
					"error": err,
				}).Error("book is not added")
				c.Status(http.StatusConflict)
			}
		}
	})
	g.GET("/books/:isbn", func(c *gin.Context) {
		isbn := c.Params.ByName("isbn")
		book, err := h.BookService.GetBook(isbn)
		if err == nil {
			c.JSON(http.StatusFound, book)
		} else {
			log.WithFields(log.Fields{
				"isbn":  isbn,
				"error": err,
			}).Error("book is not found")
			c.Status(http.StatusNotFound)
		}
	})
}
