package books

import (
	"diary-api/pkg/common/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h handler) GetBooks(c *gin.Context) {
	var books []models.Book
	if result := h.DB.Find(&books); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}
	c.JSON(http.StatusOK, &books)
}
