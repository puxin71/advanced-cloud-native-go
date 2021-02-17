package handler

import (
	"net/http"

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
	g := h.Engine.Group(h.Config.BaseURL)

	// Set the timeout period in seconds for each endpoint
	g.Use(middleware.Timeout(h.Config.TimeoutDuration, model.NewServiceUnavailable()))
	g.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
	g.GET("/books", func(c *gin.Context) {
		books := h.BookService.AllBooks()
		if len(books) == 0 {
			c.JSON(http.StatusNoContent, nil)
			return
		}
		c.JSON(http.StatusOK, books)
	})
	/*
		g.POST("/books", func(c *gin.Context) {
			var book model.Book
			if c.BindJSON(&book) == nil {
				isbn, created := h.BookService.AddBook(book)
			}
		})
	*/
}
