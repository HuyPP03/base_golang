package services

import (
	"errors"

	"github.com/HuyPP03/learn/src/database"
	"github.com/HuyPP03/learn/src/models"
	"github.com/HuyPP03/learn/src/utils"
)

func Register(username, email, password string) (*models.User, error) {
	var user models.User
	database.DB.Where("email = ?", email).First(&user)
	if user.ID != 0 {
		return nil, errors.New("Email already exists!")
	}

	hash, err := utils.HashPassword(password)
	if err != nil {
		return nil, errors.New("Invalid password!")
	}

	user.Username = username
	user.Email = email
	user.Password = string(hash)
	user.Role = "user"
	database.DB.Create(&user)

	return &user, nil
}

func Login(email, password string) (string, error) {
	var user models.User
	database.DB.Where("email = ?", email).First(&user)
	if user.ID == 0 {
		return "", errors.New("User not found!")
	}

	err := utils.ComparePassword(password, user.Password)
	if err != nil {
		return "", errors.New("Invalid password!")
	}

	token, err := utils.GenerateToken(user.ID, user.Email, user.Role)
	if err != nil {
		return "", errors.New("Could not generate token!")
	}

	return token, nil
}
