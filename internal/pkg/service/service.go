package service

import "simple-db-crud/internal/pkg/repository"

type Authorization interface {
}

type Car interface {
}

type Service struct {
	Authorization
	Car
}

func NewService(repository *repository.Repository) *Service {
	return &Service{}
}
