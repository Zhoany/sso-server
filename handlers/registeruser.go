// /cas/handlers/user.go

package handlers

import (
	"github.com/gin-gonic/gin"
	"cas/service"
)

func RegisterUserHandler(c *gin.Context) {
	username := c.PostForm("username")
	email := c.PostForm("email")
	password := c.PostForm("password")

	err := service.RegisterUser(username, email, password)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "User registered successfully!",
	})
}
