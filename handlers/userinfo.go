package handlers

import (
	"github.com/gin-gonic/gin"
	"cas/jwks"
	"log"
	"net/http"
	"cas/config"
)

func UserinfoEndpoint(c *gin.Context) {
    // 这只是为了示范。在实际应用中，您需要验证访问令牌。
    token := c.Request.Header.Get("Authorization")
	
	
	jwtString, err :=jwks.ReadJWTFromFile("/home/codeserver/Code/go/oidcprovider/cert/jwt")
	if err != nil {
		log.Fatalf("Error reading JWT from file: %v", err)
	}
	access:="bearer "+jwtString
	log.Println(access)
	log.Println(token)
    if token == "" || token != access {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
        return
    }
	
    userInfo := gin.H{
        "sub": "1234567890", // 用户的唯一标识符
        "name": config.Username,
        "email": config.Email,
        "picture": "https://example.com/johndoe.jpg",
    }

    c.JSON(http.StatusOK, userInfo)
}
