package handlers

import (
	"github.com/gin-gonic/gin"
	
	"net/http"
	//"cas/config"
	"cas/db"
)

func UserinfoEndpoint(c *gin.Context) {
    // 这只是为了示范。在实际应用中，您需要验证访问令牌。
    token := c.Request.Header.Get("Authorization")
	
	
  tokenconfig,err:=db.FindAccessTokenByToken(token[7:])
	if err!=nil{
		c.JSON(500, gin.H{"error":"token问题"})
	}
	User,err:=db.Finduserinfobysub(tokenconfig.Sub)
	if err!=nil{
		c.JSON(500, gin.H{"error":"找不到用户"})
		return 
	}

    userInfo := gin.H{
        "sub": User.Sub, // 用户的唯一标识符
        "name": User.Username,
        "email": User.Email,
        "picture": "https://example.com/johndoe.jpg",
    }

    c.JSON(http.StatusOK, userInfo)
}
