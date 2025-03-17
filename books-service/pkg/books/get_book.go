package books

import (
	"diary-api/pkg/common/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h handler) GetBook(c *gin.Context) {
	id := c.Param("id")
	var book models.Book
	if result := h.DB.Find(&book, id); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}
	c.JSON(http.StatusOK, &book)
}
