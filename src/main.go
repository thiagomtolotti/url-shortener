package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"urlshortener.com/src/handlers"
	"urlshortener.com/src/repository"
	"urlshortener.com/src/service"
)

func main() {
	loadEnvironment()

	db, err := connectDB()
	if err != nil {
		panic(err)
	}

	r := repository.NewSQLRepository(db)
	s := service.NewService(r)

	handlers.RegisterRoutes(s)
}

func loadEnvironment() {
	godotenv.Load()
}

func connectDB() (*sql.DB, error) {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASS")
	dbname := os.Getenv("DB_NAME")

	psqlInfo := fmt.Sprintf(`
        host=%s
        port=%s
        user=%s
        password=%s
        dbname=%s
        sslmode=disable
    `, host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}

	fmt.Println("[DB] Connected. Pinging...")

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	fmt.Println("[DB] Connected successfully")

	return db, nil
}
