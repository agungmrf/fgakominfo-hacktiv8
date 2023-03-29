package controller

import (
	"errors"
	"time"

	"project-2/model"
	"project-2/repository"
)

type BookController interface {
	GetAllBooks() ([]*model.Book, error)
	GetBookById(id uint) (*model.Book, error)
	CreateBook(bookInput *BookInput) (*model.Book, error)
	UpdateBook(id uint, bookInput *BookInput) (*model.Book, error)
	DeleteBook(id uint) error
}

type bookController struct {
	repo repository.BookRepository
}

type BookInput struct {
	Name   string `json:"name_book"`
	Author string `json:"author"`
}

func NewBookController(repo repository.BookRepository) BookController {
	return &bookController{
		repo: repo,
	}
}

func (c *bookController) GetAllBooks() ([]*model.Book, error) {
	return c.repo.GetAll()
}

func (c *bookController) GetBookById(id uint) (*model.Book, error) {
	return c.repo.GetById(id)
}

func (c *bookController) CreateBook(bookInput *BookInput) (*model.Book, error) {
	if bookInput.Name == "" {
		return nil, errors.New("book name cannot be empty")
	}

	book := &model.Book{
		Name:      bookInput.Name,
		Author:    bookInput.Author,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	return c.repo.Create(book)
}

func (c *bookController) UpdateBook(id uint, bookInput *BookInput) (*model.Book, error) {
	book, err := c.repo.GetById(id)
	if err != nil {
		return nil, err
	}

	if bookInput.Name != "" {
		book.Name = bookInput.Name
	}

	if bookInput.Author != "" {
		book.Author = bookInput.Author
	}

	book.UpdatedAt = time.Now()

	return c.repo.Update(id, book)
}

func (c *bookController) DeleteBook(id uint) error {
	return c.repo.Delete(id)
}
