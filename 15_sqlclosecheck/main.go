package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq" // PostgreSQL driver
)

func main() {
	// Подключение к базе данных
	const (
		host     = "localhost"
		port     = 5432
		user     = "your_username"
		password = "your_password"
		dbname   = "your_dbname"
	)
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	// Проверка подключения к базе данных
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully connected!")

	// Выполнение запроса
	rows, err := db.Query("SELECT id, name FROM your_table")
	if err != nil {
		log.Fatal(err)
	}

	// Rows/Stmt/NamedStmt was not closed (sqlclosecheck)
	//defer rows.Close()

	// Перебор результатов
	for rows.Next() {
		var id int
		var name string
		err = rows.Scan(&id, &name)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("id = %d, name = %s\n", id, name)
	}

	// Проверка на ошибки при переборе
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
}
