package ports

import (
	"garasystem/internal/core/domain"
	"github.com/labstack/echo/v4"
)

type VehicleOrderItemStore interface {
	GetList(query interface{}, args ...interface{}) ([]domain.VehicleOrderItem, error)
	Updates(vehicleOderID int64, items []domain.VehicleOrderItem) error
	Create(items []domain.VehicleOrderItem) error
}

type VehicleOrderItemHandler interface {
	CreateOrderItem(c echo.Context) error
	UpdateOrderItem(c echo.Context) error
}

type VehicleOrderItemService interface {
	GetList(vehicleOderID int64) ([]domain.VehicleOrderItemModel, error)
	Updates(domain.UpdateVehicleOrderItemRequest) error
	Create(domain.CreateVehicleOrderItemRequest) error
}
