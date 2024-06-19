package services

import "garasystem/internal/core/ports"

type Repository struct {
	UserStore        ports.UserStore
	VehicleStore     ports.VehicleOrderStore
	CategoryStore    ports.CategoryStore
	ItemStore        ports.ItemStore
	VehicleItemStore ports.VehicleOrderItemStore
}

func NewRepository(userStore ports.UserStore,
	vehicleStore ports.VehicleOrderStore,
	categoryStore ports.CategoryStore,
	itemStore ports.ItemStore,
	vehicleOrderItemStore ports.VehicleOrderItemStore,
) *Repository {
	return &Repository{
		userStore,
		vehicleStore,
		categoryStore,
		itemStore,
		vehicleOrderItemStore,
	}
}
