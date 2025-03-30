package main

import (
	"eventhub/database"
	"eventhub/handlers"
	"eventhub/middleware"
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	database.ConnectToDB()
	database.MigrateDB()
	fmt.Println("Успешное подключение к БД!")

	e := echo.New()

	e.POST("/login", handlers.LoginHandler)
	e.POST("/register", handlers.RegistrationHandler)
	e.POST("/create_event", handlers.CreateEventHandler, middleware.AuthMiddleware)
	e.POST("/join_event", handlers.JoinEvent, middleware.AuthMiddleware)
	e.POST("/quit_event", handlers.QuitEvent, middleware.AuthMiddleware)

	e.GET("/getuserdata", handlers.GetUserData, middleware.AuthMiddleware)
	e.GET("/getevents", handlers.GetEvents, middleware.AuthMiddleware)

	e.Start(":8080")
}
