package db

import (
	"cas/models"
    "fmt"
	"cas/conn"
	"log"
)

// MySQL CRUD for User

func CreateUser(user *models.User) error {
	_, err := conn.Engine.Insert(user)
	return err
}

func GetUserByID(id int) (*models.User, error) {
	user := &models.User{}
	has, err := conn.Engine.ID(id).Get(user)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, fmt.Errorf("User with ID %d not found", id)
	}
	return user, nil
}

func UpdateUser(user *models.User) error {
	_, err := conn.Engine.ID(user.ID).Update(user)
	return err
}

func DeleteUser(id int) error {
	_, err := conn.Engine.ID(id).Delete(&models.User{})
	return err
}


// MySQL CRUD for AccessToken

func CreateAccessToken(token *models.AccessToken) error {
	_, err := conn.Engine.Insert(token)
	log.Println(err)
	return err
}

func Finduserinfobysub(uuid string) (*models.User, error) {
	var User models.User
	has, err := conn.Engine.Where("sub = ?", uuid).Get(&User)
	if err != nil {
		return nil, fmt.Errorf("error getuserinfo: %v", err)
	}

	if !has {
		return nil, fmt.Errorf("no record found user")
	}

	return &User, nil
}
func FindAccessTokenByToken(token string) (*models.AccessToken, error) {
	var accessToken models.AccessToken
	has, err := conn.Engine.Where("token = ?", token).Get(&accessToken)
	if err != nil {
		return nil, fmt.Errorf("error retrieving token: %v", err)
	}

	if !has {
		return nil, fmt.Errorf("no record found for token: %s", token)
	}

	return &accessToken, nil
}
func GetAccessTokensByAudSub(aud, nonce string) (string, error) {
	var token models.AccessToken
	log.Println(aud,nonce)
	
	has,err := conn.Engine.Where("client_id = ? AND nonce = ?", aud, nonce).Get(&token)
	log.Println(err)
	if err != nil {
		return "", fmt.Errorf("find tokens error: %v", err)
	}
	if !has{
		return "", fmt.Errorf("no token in db")
	}


	
	
	return token.Token, nil
}




func DeleteAccessToken(tokenID string) error {
	_, err := conn.Engine.Where("token_id = ?", tokenID).Delete(&models.AccessToken{})
	return err
}

// MySQL CRUD for RefreshToken





// MySQL CRUD for Client

func CreateClient(client *models.Client) error {
	_, err := conn.Engine.Insert(client)
	return err
}

func GetClientByID(clientID string) (*models.Client, error) {
	client := &models.Client{}
	has, err := conn.Engine.Where("id = ?", clientID).Get(client)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, fmt.Errorf("Client with ID %s not found", clientID)
	}
	return client, nil
}

func UpdateClient(client *models.Client) error {
	_, err := conn.Engine.ID(client.ID).Update(client)
	return err
}

func DeleteClient(clientID string) error {
	_, err := conn.Engine.Where("id = ?", clientID).Delete(&models.Client{})
	return err
}
