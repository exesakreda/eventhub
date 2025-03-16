package handlers

import (
	"encoding/json"
	"eventhub/models"
	"eventhub/storage"
	"eventhub/utils"
	"log"
	"net/http"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var creds models.LoginCredentials
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		http.Error(w, "Ошибка при обработке запроса", http.StatusBadRequest)
		return
	}

	if !storage.ValidateUser(creds.Username, creds.Password) {
		http.Error(w, "Неверный логин или пароль", http.StatusUnauthorized)
		return
	}

	token, err := utils.GenerateJWT(creds.Username)
	if err != nil {
		http.Error(w, "Ошибка генерации токена", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"token": token})
}

func RegistrationHandler(w http.ResponseWriter, r *http.Request) {
	var creds models.RegistrationCredentials
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		http.Error(w, "Ошибка при обработке запроса", http.StatusBadRequest)
		return
	}

	if storage.IsUsernameTaken(creds.Username) {
		http.Error(w, "Пользователь с таким именем уже существует.", http.StatusUnauthorized)
	}

	err = storage.RegisterUser(creds.FirstName, creds.LastName, creds.Username, creds.Password)
	if err != nil {
		log.Println("Ошибка при сохранении пользователя:", err)
		http.Error(w, "Ошибка при создании пользователя", http.StatusInternalServerError)
		return
	}

	token, err := utils.GenerateJWT(creds.Username)
	if err != nil {
		http.Error(w, "Ошибка генерации токена", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Пользователь успешно зарегистрирован",
		"token":   token,
	})
}
