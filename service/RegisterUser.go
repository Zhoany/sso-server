package service

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"cas/models"
	"cas/conn"
	"cas/uuuid"

	"cas/db"
)

func RegisterUser(username, email, password string) error {
	// 1. Input Validation (你可以扩展这部分)
	if username == "" || email == "" || password == "" {
		return errors.New("all fields are required")
	}

	// 2. Check if user already exists
	existingUser := new(models.User)
	has, err := conn.Engine.Where("username = ? OR email = ?", username, email).Get(existingUser)
	if err != nil {
		return err
	}
	if has {
		return errors.New("username or email already registered")
	}

	// 3. Password hashing
	hasher := sha256.New()
	hasher.Write([]byte(password))
	encryptedPassword := hex.EncodeToString(hasher.Sum(nil))
	sub:=uuuid.GenerateTimestampID(12)
	
	// 4. Database operation to save the user
	user := &models.User{
		Sub:sub,
		
		Username: username,
		Email:    email,
		Password: encryptedPassword, // Storing encrypted password
	}

	err=db.CreateUser(user)
	if err != nil {
		return err
	}

	return nil
}
