package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"log"
	"github.com/dgrijalva/jwt-go"
	"time"
	"cas/config"
	"cas/jwks"
	
)

func TokenEndpoint(c *gin.Context) {
	grantType := c.PostForm("grant_type")
	code := c.PostForm("code")
	reqClientID := c.PostForm("client_id")
	reqClientSecret := c.PostForm("client_secret")
	nonce:=c.Query("nonce")
	log.Println("\nnonce"+nonce)
	if grantType != "authorization_code" || reqClientID != config.ClientID || reqClientSecret != config.ClientSecret || code != config.Code{
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// 生成JWT
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"iss": "your_issuer",
		"sub": "subject",
		"nonce":reqnonce,
		"aud":config.ClientID,
		"exp": time.Now().Add(time.Hour * 1).Unix(),
		// 可以添加其他claims
	})
	token.Header["kid"] = "sample-key-id"
	mySigningKey, err := jwks.LoadRSAPrivateKeyFromFile("/home/codeserver/Code/go/oidcprovider/cert/private_key.pem")
	if err!=nil{

		return
	}

	tokenString, err := token.SignedString(mySigningKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}
	jwks.WriteJWTToFile(tokenString,"/home/codeserver/Code/go/oidcprovider/cert/jwt")
	response := map[string]string{
		"access_token":  tokenString, // 使用生成的JWT
		"refresh_token": tokenString,
		"id_token":      tokenString, 
		"token_type":    "bearer",
		"expires_in":    "3600",
	}
	

	c.JSON(http.StatusOK,response)
}
