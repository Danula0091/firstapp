package books

import (
	"diary-api/pkg/common/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h handler) DeleteBook(c *gin.Context) {
	id := c.Param("id")
	var book models.Book
	if result := h.DB.First(&book, id); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
	}
	h.DB.Delete(&book)
	c.Status(http.StatusOK)
}
