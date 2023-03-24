package controllers

import (
	"net/http"
	"strconv"

	"challenge-dua/models"
	"challenge-dua/repository"

	"github.com/gin-gonic/gin"
)

type BookController interface {
	GetAllBooks(ctx *gin.Context)
	GetBookByID(ctx *gin.Context)
	AddBook(ctx *gin.Context)
	UpdateBook(ctx *gin.Context)
	DeleteBook(ctx *gin.Context)
}

type bookController struct {
	repo repository.BookRepository
}

func NewBookController(repo repository.BookRepository) BookController {
	return &bookController{
		repo: repo,
	}
}

func (c *bookController) GetAllBooks(ctx *gin.Context) {
	books := c.repo.GetAllBooks()
	ctx.JSON(http.StatusOK, books)
}

func (c *bookController) GetBookByID(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book id"})
		return
	}

	book, err := c.repo.GetBookByID(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, book)
}

func (c *bookController) AddBook(ctx *gin.Context) {
	var book models.Book
	if err := ctx.ShouldBindJSON(&book); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.repo.AddBook(&book)
	ctx.JSON(http.StatusCreated, "Created")
}

func (c *bookController) UpdateBook(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book id"})
		return
	}

	var book models.Book
	if err := ctx.ShouldBindJSON(&book); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.repo.UpdateBook(id, &book); err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, "Updated")
}

func (c *bookController) DeleteBook(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book id"})
		return
	}

	if err := c.repo.DeleteBook(id); err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, "Deleted")
}
