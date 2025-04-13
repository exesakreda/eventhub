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
	authMW := middleware.AuthMiddleware

	e.POST("/login", handlers.LoginHandler)
	e.POST("/register", handlers.RegistrationHandler)

	e.POST("/event/create", handlers.CreateEventHandler, authMW)
	e.POST("/event/join", handlers.JoinEvent, authMW)
	e.POST("/event/update", handlers.UpdateEvent, authMW)
	e.POST("/event/quit", handlers.QuitEvent, authMW)

	e.POST("/organizations/create", handlers.CreateOrganizationHandler, authMW)
	e.POST("/organizations/join", handlers.JoinOrganizationHandler, authMW)
	e.POST("/organizations/quit", handlers.QuitOrganizationHandler, authMW)

	e.GET("/user/getinfo", handlers.GetUserData, authMW)
	e.GET("/user/getevents", handlers.GetEvents, authMW)

	e.Start(":3000")
}
