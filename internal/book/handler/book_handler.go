package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/heise/myproject/Desktop/firstapp/internal/book/model"
	"github.com/heise/myproject/Desktop/firstapp/internal/book/service"
)

type BookHandler struct {
	service *service.BookService
}

func NewBookHandler(s *service.BookService) *BookHandler {
	return &BookHandler{service: s}
}

type AddBookRequestBody struct {
	Title       string `json:"title"`
	Author      string `json:"author"`
	Description string `json:"description"`
}
type UpdatedBookRequestBody struct {
	Title       string `json:"title"`
	Author      string `json:"author"`
	Description string `json:"description"`
}

func (h *BookHandler) AddBook(c *gin.Context) {
	var body AddBookRequestBody
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	book, err := h.service.CreateBook(body.Title, body.Author, body.Description)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, book)
}
func (h *BookHandler) DeleteBook(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}
	if err := h.service.DeleteBook(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error()})
	}
	c.Status(http.StatusOK)
}
func (h *BookHandler) GetBook(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid ID to fotmat"})
		return
	}
	book, err := h.service.GetBook(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}
	c.JSON(http.StatusOK, book)
}
func (h *BookHandler) GetBooks(c *gin.Context) {
	title := c.Query("title")
	author := c.Query("author")
	description := c.Query("description")
	books, err := h.service.GetBooks(title, author, description)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "cannot find books"})
		return
	}
	c.JSON(http.StatusOK, books)
}
func (h *BookHandler) UpdateBook(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid ID to fotmat"})
		return
	}
	var body UpdatedBookRequestBody
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	updatedBook := &model.Book{
		Title:       body.Title,
		Author:      body.Author,
		Description: body.Description,
	}
	result, err := h.service.UpdateBook(uint(id), updatedBook)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}
