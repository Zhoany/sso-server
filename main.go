package main

import (
	"cas/handlers"
	"cas/middleware"
	"cas/conn"
	"log"
	"github.com/gin-gonic/gin"
	"cas/jwks"
)

func main() {
	jwks.InitJWKS()
	err:=conn.InitDB()
	if err!=nil{
		log.Println(err)
		return 
	}
	err=conn.InitRedis()
	if err!=nil{
		log.Println(err)
		return 
	}
	r := gin.Default()
	r.Use(middleware.PrintRequestBody())

	// 创建新的cookie存储实例
	

	
	
	r.LoadHTMLGlob("templates/*")
	r.GET("/.well-known/openid-configuration", handlers.DiscoveryEndpoint)
	//r.POST("/authorize/login", handlers.PostAuthorizeEndpoint)
	//r.GET("/authorize", handlers.GetAuthorizeEndpoint)
	r.Any("/authorize", handlers.GetAndPostAuthorizeEndpoint)
	r.POST("/token", handlers.TokenEndpoint)
	r.GET("/userinfo", handlers.UserinfoEndpoint)
	r.GET("/jwks", handlers.JWKSEndpoint)
	r.POST("/register/user",handlers.RegisterUserHandler)
	r.GET("/register/client",handlers.RegisterClientHandler)
	r.Run(":8080")
}
