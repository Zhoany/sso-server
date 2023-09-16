package handlers

import (

	"github.com/gin-gonic/gin"
	"cas/uuuid"
	
	"net/http"
	
	"cas/service"
	"cas/db"
	"cas/models"
	

)

func GetAndPostAuthorizeEndpoint(c *gin.Context) {

	if c.Request.Method == "GET" {
		// Save the parameters to the session
		responseType:=c.Query("response_type")
		clientID:=c.Query("client_id")
		redirectURI:=c.Query("redirect_uri")
		nonce:=c.Query("nonce")
		state:=c.Query("state")
		
        
		c.HTML(http.StatusOK, "login.tmpl", gin.H{
			"responseType": responseType,
			"clientID":     clientID,
			"redirectURI":  redirectURI,
			"nonce":        nonce,
			"state":        state,
		})
		return
	}

	// 如果您使用的是gin框架，这是一个示例handler函数
    

    // 从这里开始，您可以使用responseType, clientID和redirectURI进行后续的操作


	if c.Request.Method == "POST"{
		responseType:=c.Query("response_type")
		clientID:=c.Query("client_id")
		redirectURI:=c.Query("redirect_uri")
		nonce:=c.Query("nonce")
		state:=c.Query("state")
		scope:=c.Query("scope")
		if scope==""{
			scope="openid"
		}
		code:=uuuid.GenerateTimestampID(24)
		

		
	if !service.ValidateRequestParameters(responseType, clientID, redirectURI) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request parameters"})
		return
	}

	
	username := c.PostForm("username")
	password := c.PostForm("password")

	sub,flag:= service.AuthenticateUser(username, password)
	if !flag{
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication failed"})
		return
	}
	if sub==""{
		c.JSON(500, gin.H{"error": "Get user failed"})
		return
	}
	
	err:=service.GenAndWriteTokenToDB(sub ,nonce ,clientID ,scope,BaseURL)
	if err!=nil{
		c.JSON(500, gin.H{"error": "Gen token error"})
	}
	redissession:=&models.RedisSession{
		Code:code,
		Nonce:nonce,
	}
	db.CreateRedisSession(redissession)


	

  
	c.Redirect(http.StatusFound, redirectURI+"?code="+code+"&nonce="+nonce+"&state="+state)
}}

