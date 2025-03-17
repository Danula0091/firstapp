package controllers

import (
	"auth/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RegisterInput struct {
	Email    string `json:"email" binding:"required,unique"`
	Password string `json:"password" binding:"required"`
}

func (h *handler) Register(c *gin.Context) {
	var input RegisterInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	u := models.User{}
	u.Email = input.Email
	u.Password = input.Password
	savedUser, err := u.SaveUser(h.DB)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
	}

	if err := h.DB.Create(&u).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при сохранении пользователя"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "registration success",
		"user":    savedUser})

}
