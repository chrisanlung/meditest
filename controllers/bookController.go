package controllers

import (
	"net/http"

	"github.com/chrisanlung/meditest/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Controller struct {
	DB *gorm.DB
}

func (ctrl Controller) GetAllBooks(c *gin.Context) {
	var books []models.Book
	ctrl.DB.Find(&books)
	c.JSON(http.StatusOK, books)
}

func (ctrl Controller) GetBookByID(c *gin.Context) {
	var book models.Book
	id := c.Param("id")
	if err := ctrl.DB.First(&book, id).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Book not found!"})
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
	ctrl.DB.Create(&book)
	c.JSON(http.StatusOK, book)
}
