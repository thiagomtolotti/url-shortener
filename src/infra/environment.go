package infra

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

var (
	PORT,
	DB_HOST,
	DB_PORT,
	DB_USER,
	DB_PASS,
	DB_NAME string
)

func LoadEnvironment() error {
	err := godotenv.Load()
	if err != nil {
		return err
	}

	vars := map[string]*string{
		"PORT":    &PORT,
		"DB_HOST": &DB_HOST,
		"DB_PORT": &DB_PORT,
		"DB_USER": &DB_USER,
		"DB_PASS": &DB_PASS,
		"DB_NAME": &DB_NAME,
	}

	for key, target := range vars {
		value := os.Getenv(key)
		if value == "" {
			return fmt.Errorf("could not find env variable %s", key)
		}
		*target = value
	}

	return nil
}
