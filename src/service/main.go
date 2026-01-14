package service

import (
	"crypto/rand"
	"encoding/hex"
)

type ID string

var db map[ID]string = make(map[ID]string, 0)

func CreateURL(originalURL string) ID {
	id, err := newRandomId()
	if err != nil {
		panic(err)
	}

	db[id] = originalURL

	return id
}

func GetURL(id string) string {
	return db[ID(id)]
}

func newRandomId() (ID, error) {
	bytes := make([]byte, 4)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}

	return ID(hex.EncodeToString(bytes)), nil
}
