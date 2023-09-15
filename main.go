package main

import (
	
	"cas/middleware"
	"cas/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(middleware.PrintRequestBody())
	r.LoadHTMLGlob("templates/*")
	r.GET("/.well-known/openid-configuration", handlers.DiscoveryEndpoint)
	r.POST("/authorize/login", handlers.PostAuthorizeEndpoint)
	r.GET("/authorize", handlers.GetAuthorizeEndpoint)
	r.POST("/token", handlers.TokenEndpoint)
	r.GET("/userinfo", handlers.UserinfoEndpoint)
	r.GET("/jwks", handlers.JWKSEndpoint)
	r.Run(":8080")
}