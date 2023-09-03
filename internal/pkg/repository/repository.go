package repository

import "github.com/jmoiron/sqlx"

type Authorization interface {
}

type Car interface {
}

type Repository struct {
	Authorization
	Car
}

func NewRepository(database *sqlx.DB) *Repository {
	return &Repository{}
}
