package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/RedContritio/E-commerce-system/models"
	"github.com/RedContritio/E-commerce-system/services"
	"net/http"
)

type UserHandler struct {
	UserService *services.UserService
}

func (h *UserHandler) RegisterUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.UserService.CreateUser(user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully!"})
}