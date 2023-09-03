package repository

type Authorization interface {
}

type Car interface {
}

type Repository struct {
	Authorization
	Car
}

func NewRepository() *Repository {
	return &Repository{}
}
