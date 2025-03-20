package service

import (
	"errors"

	"github.com/heise/myproject/Desktop/firstapp/internal/book/model"
	"github.com/heise/myproject/Desktop/firstapp/internal/book/repository"
)

type BookService struct {
	repo *repository.BookRepository
}

func NewBookService(repo *repository.BookRepository) *BookService {
	return &BookService{repo: repo}
}
func (s *BookService) CreateBook(title, author, description string) (*model.Book, error) {
	if title == "" || author == "" {
		return nil, errors.New("title and author are required")
	}
	book := &model.Book{
		Title:       title,
		Author:      author,
		Description: description,
	}
	if err := s.repo.CreateBook(book); err != nil {
		return nil, err
	}
	return book, nil
}
func (s *BookService) GetBooks(title, author, description string) ([]*model.Book, error) {
	if title == "" && author == "" {
		return nil, errors.New("title or author are required")
	}
	books, err := s.repo.GetBooks(title, author, description)
	if err != nil {
		return nil, err
	}
	return books, nil
}
func (s *BookService) GetBook(id uint) (*model.Book, error) {
	if id == 0 {
		return nil, errors.New("id is required")
	}
	book, err := s.repo.GetBook(id)
	if err != nil {
		return nil, err
	}
	return book, nil
}
func (s *BookService) DeleteBook(id uint) error {
	return s.repo.DeleteBook(id)
}
func (s *BookService) UpdateBook(id uint, updatedBook *model.Book) (*model.Book, error) {
	return s.repo.UpdateBook(id, updatedBook)
}
