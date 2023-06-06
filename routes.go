package main

import (
	"github.com/chrisanlung/meditest/controllers"
	"github.com/gin-gonic/gin"
)

func InitializeRoutes(r *gin.Engine, controller controllers.Controller) {
	r.GET("/books", controller.GetAllBooks)
	r.GET("/books/:id", controller.GetBookByID)
	r.POST("/books", controller.CreateBook)
}
