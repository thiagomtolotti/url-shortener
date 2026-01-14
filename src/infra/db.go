package infra

import (
	"database/sql"
	"fmt"
)

func ConnectToDB() (*sql.DB, error) {
	host := DB_HOST
	port := DB_PORT
	user := DB_USER
	password := DB_PASS
	dbname := DB_NAME

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
