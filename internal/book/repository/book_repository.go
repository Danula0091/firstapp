package repository

import (
	"errors"
	"fmt"
	"log"

	"github.com/heise/myproject/Desktop/firstapp/internal/book/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type BookRepository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) *BookRepository {
	return &BookRepository{db: db}
}
func (r *BookRepository) CreateBook(book *model.Book) error {
	return r.db.Create(book).Error
}
func InitDB(url string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	log.Println("Connected to the database successfully")
	return db, nil
}
func (r *BookRepository) GetBooks(title, author, description string) ([]*model.Book, error) {
	var books []*model.Book
	query := r.db.Model(&model.Book{})
	if title != "" {
		query = query.Where("title = ?", title)
	}
	if author != "" {
		query = query.Where("author = ?", author)
	}
	if description != "" {
		query = query.Where("description LIKE '%" + description + "%'")
	}
	if err := query.Find(&books).Error; err != nil {
		return nil, err
	}
	return books, nil
}
func (r *BookRepository) GetBook(id uint) (*model.Book, error) {
	var book model.Book
	if err := r.db.First(&book, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("book with id %d not found", id)
		}
		return nil, err
	}
	return &book, nil
}
func (r *BookRepository) DeleteBook(id uint) error {
	if err := r.db.Delete(&model.Book{}, id).Error; err != nil {
		return fmt.Errorf("delete is failed with book id %d", id)
	}
	return nil
}
func (r *BookRepository) UpdateBook(id uint, updatedBook *model.Book) (*model.Book, error) {
	var book model.Book
	if err := r.db.First(&book, id).Error; err != nil {
		return nil, fmt.Errorf("book id %d not found", id)
	}
	if err := r.db.Model(&book).Updates(map[string]interface{}{
		"title":       updatedBook.Title,
		"author":      updatedBook.Author,
		"description": updatedBook.Description,
	}).Error; err != nil {
		return nil, err
	}

	return &book, nil
}
