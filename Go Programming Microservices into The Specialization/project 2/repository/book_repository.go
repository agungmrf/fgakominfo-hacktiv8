package repository

import (
	"errors"
	"fmt"
	"github.com/jinzhu/copier"
	"time"

	"gorm.io/gorm"

	"project-2/model"
)

type BookRepository interface {
	GetAll() ([]*model.Book, error)
	GetById(id uint) (*model.Book, error)
	Create(book *model.Book) (*model.Book, error)
	Update(id uint, book *model.Book) (*model.Book, error)
	Delete(id uint) error
}

type bookRepository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) BookRepository {
	return &bookRepository{db}
}

func (repo *bookRepository) GetAll() ([]*model.Book, error) {
	var books []*model.Book
	if err := repo.db.Find(&books).Error; err != nil {
		return nil, fmt.Errorf("failed to get books: %v", err)
	}
	return books, nil
}

func (repo *bookRepository) GetById(id uint) (*model.Book, error) {
	var book model.Book
	if err := repo.db.First(&book, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("failed to get book: book not found")
		}
		return nil, fmt.Errorf("failed to get book: %v", err)
	}
	return &book, nil
}

func (repo *bookRepository) Create(book *model.Book) (*model.Book, error) {
	newBook := model.Book{
		Name:      book.Name,
		Author:    book.Author,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	if err := repo.db.Create(&newBook).Error; err != nil {
		return nil, fmt.Errorf("failed to create book: %v", err)
	}
	return &newBook, nil
}

func (repo *bookRepository) Update(id uint, book *model.Book) (*model.Book, error) {
	var existingBook model.Book
	result := repo.db.First(&existingBook, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("failed to update book: book not found")
		}
		return nil, fmt.Errorf("failed to update book: %v", result.Error)
	}
	if err := copier.Copy(&existingBook, book); err != nil {
		return nil, fmt.Errorf("failed to update book: %v", err)
	}
	existingBook.UpdatedAt = time.Now()
	if err := repo.db.Save(&existingBook).Error; err != nil {
		return nil, fmt.Errorf("failed to update book: %v", err)
	}
	return &existingBook, nil
}

func (repo *bookRepository) Delete(id uint) error {
	var book model.Book
	result := repo.db.First(&book, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return fmt.Errorf("failed to delete book: book not found")
		}
		return fmt.Errorf("failed to delete book: %v", result.Error)
	}
	if err := repo.db.Delete(&book).Error; err != nil {
		return fmt.Errorf("failed to delete book: %v", err)
	}
	return nil
}
