package storage

import (
	"database/sql"
	"eventhub/utils"
	"log"

	"golang.org/x/crypto/bcrypt"
)

var users = map[string]string{
	"admin":   "admin",
	"user123": "qwerty",
}

func ValidateUser(username, password string) bool {
	var hashedPassword string
	err := DB.QueryRow("SELECT password_hash FROM users WHERE username = $1", username).Scan(&hashedPassword)
	if err == sql.ErrNoRows {
		return false
	} else if err != nil {
		log.Println("Ошибка при запросе в БД: ", err)
		return false
	}

	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}

func IsUsernameTaken(username string) bool {
	var exists bool
	err := DB.QueryRow("SELECT EXISTS (SELECT 1 FROM users WHERE username = $1);", username).Scan(&exists)
	if err != nil {
		log.Println("Ошибка при проверке имени пользователя: ", err)
		return false
	}
	return exists
}

func RegisterUser(firstName, lastName, username, password string) error {
	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		return err
	}

	_, err = DB.Exec("INSERT INTO users (first_name, last_name, username, password_hash) VALUES ($1, $2, $3, $4);", firstName, lastName, username, string(hashedPassword))
	return err
}
