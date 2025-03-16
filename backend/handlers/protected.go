package handlers

import (
	"database/sql"
	"encoding/json"
	"eventhub/models"
	"eventhub/storage"
	"fmt"
	"net/http"
)

func ProtectedHandler(w http.ResponseWriter, r *http.Request) {
	username := r.Context().Value("username").(string)
	fmt.Fprintf(w, "Привет, %s! Это защищенная зона.", username)
}

func GetUserData(w http.ResponseWriter, r *http.Request) {
	username, ok := r.Context().Value("username").(string)
	if !ok {
		http.Error(w, "Неавторизованный доступ", http.StatusUnauthorized)
		return
	}

	var user models.User
	err := storage.DB.QueryRow("SELECT first_name, last_name FROM users WHERE username = $1", username).Scan(&user.FirstName, &user.LastName)
	user.Username = username
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Пользователь не найден", http.StatusNotFound)
		} else {
			http.Error(w, string(err.Error()), http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}
