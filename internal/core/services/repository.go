package services

import "garasystem/internal/core/ports"

type Repository struct {
	UserStore    ports.UserStore
	VehicleStore ports.VehicleOrderStore
}

func NewRepository(userStore ports.UserStore, vehicleStore ports.VehicleOrderStore) *Repository {
	return &Repository{userStore, vehicleStore}
}
