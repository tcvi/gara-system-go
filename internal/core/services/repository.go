package services

import "garasystem/internal/core/ports"

type Repository struct {
	UserStore     ports.UserStore
	VehicleStore  ports.VehicleOrderStore
	CategoryStore ports.CategoryStore
}

func NewRepository(userStore ports.UserStore, vehicleStore ports.VehicleOrderStore, categoryStore ports.CategoryStore) *Repository {
	return &Repository{userStore, vehicleStore, categoryStore}
}
