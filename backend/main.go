package main

import (
	"eventhub/handlers"
	"eventhub/middleware"
	"eventhub/storage"
	"fmt"
	"net/http"
)

func main() {
	storage.InitDB()

	PORT := 8080
	addr := fmt.Sprintf(":%d", PORT)

	mux := http.NewServeMux()

	mux.HandleFunc("/login", handlers.LoginHandler)
	mux.HandleFunc("/register", handlers.RegistrationHandler)
	mux.Handle("/getuserdata", middleware.AuthMiddleware(http.HandlerFunc(handlers.GetUserData)))

	fmt.Println("Сервер запущен на порту ", PORT)
	http.ListenAndServe(addr, mux)
}
