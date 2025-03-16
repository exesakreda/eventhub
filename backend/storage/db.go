package storage

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() {
	var err error

	connStr := "host=localhost port=5432 user=sakreda password=qweqweqwE5! dbname=eventhub sslmode=disable"

	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Ошибка подключения к базе: ", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal("Не удалось подключиться к базе: ", err)
	}

	fmt.Println("Успешное подключение к базе.")
}
