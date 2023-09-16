package service

import(
	"crypto/sha256"
	"encoding/hex"
	"github.com/dgrijalva/jwt-go"
	"cas/models"
	"cas/conn"
	"time"
	"cas/jwks"
	"cas/db"
	"log"
	
	
	

	
)

func ValidateRequestParameters(responseType, clientID, redirectURI string) bool {
    // 先检查responseType
    if responseType != "code" {
		log.Println("not code")
        return false
    }

    client := new(models.Client)
	log.Println(clientID+"\n"+redirectURI)
    has, err := conn.Engine.Where("client_id = ? AND redirect_uri = ?", clientID, redirectURI).Get(client)
    
    // 如果查询出错或没有找到匹配的客户端，返回false
    if err != nil || !has {
		
        return false
    }

    return true
}


func AuthenticateUser(username, password string) (string,bool) {
	user := new(models.User)
    hasher := sha256.New()
	hasher.Write([]byte(password))
	encryptedPassword := hex.EncodeToString(hasher.Sum(nil))
	has, err := conn.Engine.Where("username = ? AND password= ?", username,encryptedPassword).Get(user)
	if err != nil {
		log.Println(err)
		return "",false
	}
	if !has {
		log.Println("has\n",has)
		
		return "",false
	}
	
	return user.Sub,true
}
func GenAndWriteTokenToDB(sub string,nonce string,clientID string,scope string,baseurl string )error{
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"iss":  baseurl,
		"sub":   sub,
		"nonce": nonce, // 使用从session中获取的nonce
		"aud":   clientID,
		"exp":   time.Now().Add(time.Hour * 1).Unix(),
		"iat":   time.Now().Unix(),
		"scope":scope,
		
	})
	token.Header["kid"] = "e06a642f-f66b-46bb-8272-145c6edb26ea"
	mySigningKey, err := jwks.LoadRSAPrivateKeyFromFile("./cert/private_key.pem")
	if err != nil {
		// 这里添加了日志输出来更好地跟踪错误
		log.Println("read sign err")
		return err
	}

	tokenString, err := token.SignedString(mySigningKey)
	if err != nil {
		log.Println("sign err")
		return err
	}
	dbtoken := &models.AccessToken{
		Token :tokenString,
		Sub   :sub,
		Aud   :clientID,
		Nonce :nonce,
		Exp:time.Unix(time.Now().Add(time.Hour * 1).Unix(), 0),
		Iat:time.Unix(time.Now().Unix(),0),
		Scope :scope,
	}
	log.Println(dbtoken)
	err = db.CreateAccessToken(dbtoken)
	if err != nil {
		log.Println("db create err")
		return err
	}
	return nil
}