package service

import (
	"crypto/rand"
	"encoding/hex"

	"urlshortener.com/src/repository"
)

type Service struct {
	repo repository.Repository
}

func NewService(repo repository.Repository) Service {
	return Service{
		repo: repo,
	}
}

func (s *Service) CreateURL(originalURL string) string {
	id, err := newRandomId()
	if err != nil {
		panic(err)
	}

	s.repo.CreateURL(originalURL, id)

	return id
}

func (s *Service) GetURL(id string) (string, error) {
	return s.repo.GetURL(id)
}

func newRandomId() (string, error) {
	bytes := make([]byte, 4)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}

	return (hex.EncodeToString(bytes)), nil
}
