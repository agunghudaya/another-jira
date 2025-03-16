package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetUserHandler handles GET /users/:id
func GetUserHandler(c *gin.Context) {
	userID := c.Param("id")
	c.JSON(http.StatusOK, gin.H{"message": "User found", "id": userID})
}

// CreateUserHandler handles POST /users
func CreateUserHandler(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{"message": "User created"})
}

func GetHealthHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "OK"})
}
