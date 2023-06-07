package controllers

import (
	"net/http"

	"github.com/chrisanlung/meditest/models"
	"github.com/chrisanlung/meditest/services"
	"github.com/gin-gonic/gin"
)

type Controller struct {
	BookService services.BookService
}

func (ctrl Controller) GetAllBooks(c *gin.Context) {
	books, err := ctrl.BookService.GetAllBooks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, books)
}

func (ctrl Controller) GetBookByID(c *gin.Context) {
	id := c.Param("id")
	book, err := ctrl.BookService.GetBookByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, book)
}

func (ctrl Controller) CreateBook(c *gin.Context) {
	var book models.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdBook, err := ctrl.BookService.CreateBook(book)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, createdBook)
}
