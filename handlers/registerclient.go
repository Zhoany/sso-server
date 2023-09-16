// /cas/handlers/user.go

package handlers

import (
	"github.com/gin-gonic/gin"
	"cas/service"
)

func RegisterClientHandler(c *gin.Context) {
Client_name:=c.Query("client_name")
Redirect_uri:=c.Query("redirect_uri")
id,sercet,err := service.Registerclient(Redirect_uri, Client_name)
if err != nil {
	c.JSON(400, gin.H{
		"error": err.Error(),
	})
	return
}

c.JSON(200, gin.H{
	"message": "Client registered successfully!",
	"client_id":id,
	"client_sercet":sercet,
})
}
