package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"log"
	"cas/models"
	"cas/conn"
	"cas/db"
	
)

func TokenEndpoint(c *gin.Context) {
	grantType := c.PostForm("grant_type")
	code := c.PostForm("code")
	reqClientID := c.PostForm("client_id")
	reqClientSecret := c.PostForm("client_secret")
	if grantType != "authorization_code" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid grant type"})
		return
	}

	// 验证客户端ID和客户端密钥
	client := new(models.Client)
	has, err := conn.Engine.Where("client_id = ? AND client_secret = ?", reqClientID, reqClientSecret).Get(client)
	if err != nil {
		log.Println("Database error:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}
	if !has {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid client credentials"})
		return
	}
	sessions,err0:=db.GetRedisSessionByCode(code)
	if err0!=nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": "code error"})
		return
	}
	sendtoken,err3:=db.GetAccessTokensByAudSub(reqClientID,sessions.Nonce)
	log.Println(sendtoken)
	if err3!=nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": "getaccesstoken error"})
		return
	}
	response := map[string]string{
		"access_token":  sendtoken, // 使用生成的JWT
		"refresh_token": sendtoken,
		"id_token":      sendtoken, 
		"token_type":    "bearer",
		"expires_in":    "3600",
	}
	c.JSON(http.StatusOK,response)


}
