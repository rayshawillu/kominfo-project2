package controllers

import (
	"book-challenge/database"
	"book-challenge/models"
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"gorm.io/gorm"
)

type Book struct {
	BookID    string `json:"id"`
	Title     string `json:"name_book"`
	Author    string `json:"author"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func GetBooks(ctx *gin.Context) {
	var results = []Book{}
	var requestBook = Book{}

	db := database.GetDB()

	books := []models.Book{}

	rows := db.Find(&books)

	if rows.Error != nil {
		ctx.AbortWithError(http.StatusInternalServerError, rows.Error)
		return
	}

	for _, data := range books {
		requestBook = convertToJson(data)
		results = append(results, requestBook)
	}

	ctx.JSON(http.StatusOK, results)
}

func GetBookById(ctx *gin.Context) {
	bookID := ctx.Param("bookID")
	book := models.Book{}

	db := database.GetDB()

	err := db.First(&book, "id=?", bookID).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"message": "Book data not found",
			})
			return
		}
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, convertToJson(book))
}

func CreateBook(ctx *gin.Context) {
	var requestBook Book

	if err := ctx.ShouldBindJSON(&requestBook); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	db := database.GetDB()

	newBook := models.Book{
		Name:   requestBook.Title,
		Author: requestBook.Author,
	}

	err := db.Create(&newBook).Error

	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, convertToJson(newBook))
}

func UpdateBook(ctx *gin.Context) {
	bookID := ctx.Param("bookID")
	var requestBook Book

	if err := ctx.ShouldBindJSON(&requestBook); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	db := database.GetDB()

	var book models.Book
	err := db.First(&book, "id=?", bookID).Find(&book).Error
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"message": "Book data not found",
		})
		return
	}
	book.Name = requestBook.Title
	book.Author = requestBook.Author
	result := db.Save(&book)

	if result.Error != nil {
		ctx.AbortWithError(http.StatusInternalServerError, result.Error)
		return
	}

	if result.RowsAffected == 0 {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"message": "Book data not found",
		})
		return
	}

	ctx.JSON(http.StatusOK, convertToJson(book))
}

func DeleteBook(ctx *gin.Context) {
	bookID := ctx.Param("bookID")

	db := database.GetDB()

	book := models.Book{}

	result := db.Model(&book).Where("id = ?", bookID).Delete(&book)

	if result.Error != nil {
		ctx.AbortWithError(http.StatusInternalServerError, result.Error)
		return
	}

	if result.RowsAffected == 0 {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"message": "Book data not found",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Book deleted successfully",
	})
}

func convertToJson(data models.Book) Book {
	return Book{
		BookID:    strconv.Itoa(int(data.ID)),
		Title:     data.Name,
		Author:    data.Author,
		CreatedAt: data.CreatedAt.Format(time.RFC3339),
		UpdatedAt: data.UpdatedAt.Format(time.RFC3339),
	}
}
