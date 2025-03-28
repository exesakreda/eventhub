package database

import (
	"errors"
	"eventhub/models"
	"eventhub/utils"
	"fmt"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var users = map[string]string{
	"admin":   "admin",
	"user123": "qwerty",
}

func ValidateUser(username, password string) (bool, error) {
	var user models.User

	err := DB.Where("username = ?", username).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, err
		}
		return false, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return false, nil
	}

	return true, nil
}

func IsUsernameTaken(username string) (bool, error) {
	var exists bool
	err := DB.Model(&models.User{}).Where("username = ?", username).Select("1").Scan(&exists).Error
	if err != nil {
		return false, err
	}
	return exists, nil
}

func RegisterUser(firstName, lastName, username, password string) error {
	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		return err
	}

	user := models.User{
		FirstName: firstName,
		LastName:  lastName,
		Username:  username,
		Password:  hashedPassword,
	}

	err = DB.Create(&user).Error
	if err != nil {
		return fmt.Errorf("Ошибка при создании пользователя: %v", err)
	}

	return nil
}
