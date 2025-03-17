package controllers

import (
	"auth/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginInput struct {
	Email    string `json:"email" binding:"required,unique"`
	Password string `json:"password" binding:"required"`
}

func (h *handler) Login(c *gin.Context) {
	var input LoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	u := models.User{}
	u.Email = input.Email
	u.Password = input.Email
	token, err := u.LoginCheck(h.DB, input.Email, input.Password)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "user not found, username or password is incorrect"})
		return
	}
	c.JSON(http.StatusFound, gin.H{
		"token": token})
}
