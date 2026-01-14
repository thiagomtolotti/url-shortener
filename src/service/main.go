package service

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"

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
	var id string
	for {
		id, err := newRandomId()
		if err != nil {
			panic(err)
		}

		exists, err := s.repo.Exists(id)
		if err != nil {
			panic(err)
		}

		if !exists {
			break
		} else {
			fmt.Println("Collision detected, generating new id")
		}
	}

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

func (s *Service) DeleteURL(id string) error {
	exists, err := s.repo.Exists(id)
	if err != nil {
		return err
	}

	if !exists {
		return errors.New("url not found")
	}

	return s.repo.DeleteURL(id)
}

func newRandomId() (string, error) {
	bytes := make([]byte, 4)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}

	return (hex.EncodeToString(bytes)), nil
}
