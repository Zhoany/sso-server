package handlers

import (
	"github.com/gin-gonic/gin"
	"cas/config"
	"net/http"
	"log"
	

)
var (responseType ="" 
	reqClientID = ""
	reqRedirectURI ="" 
	reqnonce = ""
	reqstate ="")

func PostAuthorizeEndpoint(c *gin.Context) {
	// 1. 验证请求参数
	
	if responseType != "code" || reqClientID != config.ClientID || reqRedirectURI != config.RedirectURI {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request parameters"})
		return
	}

	// 2. 用户身份验证
	// 在此示例中，我们将假设一个简单的用户验证过程，您可以通过数据库或其他方式来实现真正的身份验证
	username := c.PostForm("username")
	password := c.PostForm("password")
	
	if username != config.LoginUsername || password != config.LoginPassword {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication failed"})
		return
	}

	
	code := config.Code
	log.Println("nonce:"+reqnonce+"\n"+"state:"+reqstate)
	c.Redirect(http.StatusFound, reqRedirectURI+"?code="+code+"&nonce="+reqnonce+"&state="+reqstate)
}

func GetAuthorizeEndpoint(c *gin.Context) {
	responseType = c.Query("response_type")
	reqClientID = c.Query("client_id")
	reqRedirectURI = c.Query("redirect_uri")
	reqnonce = c.Query("nonce")
	reqstate = c.Query("state")
	c.HTML(http.StatusOK, "login.tmpl",nil)
}
