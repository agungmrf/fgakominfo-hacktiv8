package repository

import (
	"errors"
	"sync"

	"challenge-dua/models"
)

type BookRepository interface {
	GetAllBooks() []*models.Book
	GetBookByID(id int) (*models.Book, error)
	AddBook(book *models.Book) int
	UpdateBook(id int, book *models.Book) error
	DeleteBook(id int) error
}

type bookRepository struct {
	books []*models.Book
	mu    sync.RWMutex
}

func NewBookRepository() BookRepository {
	return &bookRepository{
		books: []*models.Book{},
	}
}

func (r *bookRepository) GetAllBooks() []*models.Book {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return r.books
}

func (r *bookRepository) GetBookByID(id int) (*models.Book, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	for _, book := range r.books {
		if book.ID == id {
			return book, nil
		}
	}
	return nil, errors.New("book not found")
}

func (r *bookRepository) AddBook(book *models.Book) int {
	r.mu.Lock()
	defer r.mu.Unlock()
	id := len(r.books) + 1
	book.ID = id
	r.books = append(r.books, book)
	return id
}

func (r *bookRepository) UpdateBook(id int, book *models.Book) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	for i, b := range r.books {
		if b.ID == id {
			book.ID = id
			r.books[i] = book
			return nil
		}
	}
	return errors.New("book not found")
}

func (r *bookRepository) DeleteBook(id int) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	for i, book := range r.books {
		if book.ID == id {
			r.books = append(r.books[:i], r.books[i+1:]...)
			return nil
		}
	}
	return errors.New("book not found")
}
