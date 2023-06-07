package services

import (
	"github.com/chrisanlung/meditest/models"
	"gorm.io/gorm"
)

type BookService struct {
	DB *gorm.DB
}

func (service *BookService) GetAllBooks() ([]models.Book, error) {
	var books []models.Book
	err := service.DB.Find(&books).Error
	if err != nil {
		return nil, err
	}
	return books, nil
}

func (service *BookService) CreateBook(book models.Book) (models.Book, error) {
	err := service.DB.Create(&book).Error
	if err != nil {
		return models.Book{}, err
	}
	return book, nil
}

func (service *BookService) GetBookByID(id string) (models.Book, error) {
	var book models.Book
	err := service.DB.First(&book, id).Error
	if err != nil {
		return models.Book{}, err
	}
	return book, nil
}
