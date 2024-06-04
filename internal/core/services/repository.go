package services

import "garasystem/internal/core/ports"

type Repository struct {
	UserStore ports.UserStore
}

func NewRepository(userStore ports.UserStore) *Repository {
	return &Repository{userStore}
}
