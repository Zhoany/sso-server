package service

import (
	"errors"
	"cas/models"
	"cas/conn"
	"cas/uuuid"
	"cas/db"
)

func Registerclient(Redirect_uri string, Client_name string) (string, string, error) {
	if Redirect_uri == "" || Client_name == "" {
		return "", "", errors.New("all fields are required")
	}

	existingClient := new(models.Client)
	has, err := conn.Engine.Where("client_name = ? and redirect_uri  = ?", Client_name,Redirect_uri).Get(existingClient)
	if err != nil {
		return "", "", err
	}
	if has {
		return "", "", errors.New("client already registered")
	}

	client_id:= uuuid.GenerateTimestampID(12)
	

	client_secret:= uuuid.GenerateTimestampID(28)
	

	client := &models.Client{
		ID:          client_id,
		Secret:      client_secret,
		RedirectURI: Redirect_uri,
		Client_name: Client_name,
	}

	err = db.CreateClient(client)
	if err != nil {
		return "", "", err
	}

	return client_id, client_secret, nil
}
