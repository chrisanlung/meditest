package main

import (
	"github.com/chrisanlung/meditest/controllers"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Title  string
	Author string
}

func main() {
	dsn := "root:@tcp(localhost:3306)/pindah?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&Book{})

	r := gin.Default()

	controller := controllers.Controller{DB: db}

	InitializeRoutes(r, controller)

	r.Run()
}
