package controllers

import (
	"go-book-management/database"
	"go-book-management/helpers"
	"go-book-management/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BookController struct{}

type AddBookBody struct {
	AnimalId int64  `json:"animal_id"`
	Name     string `json:"name"`
	Author   string `json:"author"`
	Genre    string `json:"genre"`
}

func (ctrl BookController) GetBooks(context *gin.Context) {
	var books []*models.Book
	err := database.DB.Table("\"book\" b").
		Select("b.id, b.name, b.author, b.genre").
		Find(&books).Error
	helpers.HandleErr(err)

	context.JSON(http.StatusOK, gin.H{"response": books})
}

func (ctrl BookController) AddBook(context *gin.Context) {
	var book []*models.Book

	body := AddBookBody{}
	err_body := context.BindJSON(&body)
	helpers.HandleErr(err_body)

	err := database.DB.Raw("INSERT INTO public.book(name, author, genre) VALUES (?, ?, ?);",
		body.Name, body.Author, body.Genre).Scan(&book).Error
	helpers.HandleErr(err)

	context.JSON(http.StatusOK, gin.H{"response": book})
}
