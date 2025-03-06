package handlers

import (
	"net/http"
	"polyprep/internal/auth"

	"github.com/gin-gonic/gin"
)

type RegisterRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func RegisterUser(c *gin.Context) {
	var request RegisterRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	if !validateEmail(request.Email) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "not spbstu"})
		return
	}

	err := auth.CreateUser(request.Username, request.Email, request.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "registration failed: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"log": "good registration"})
}

func validateEmail(email string) bool {
	return len(email) > 13 && email[len(email)-14:] == "@edu.spbstu.ru"
}

func LoginUser(c *gin.Context) {
	var request LoginRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	token, err := auth.AuthenticateUser(request.Username, request.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication failed: " + err.Error()})

		return
	}

	c.JSON(http.StatusOK, gin.H{"access_token": token})
}
