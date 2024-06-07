package ports

import (
	"garasystem/internal/core/domain"
	"github.com/labstack/echo/v4"
)

type VehicleOrderStore interface {
	GetByID(id int64) (*domain.VehicleOrder, error)
	GetList(*domain.FilterRequest) ([]domain.VehicleOrder, error)
	Create(*domain.VehicleOrder) error
	Update(vehicleOrder *domain.VehicleOrder, isLogStatusHistory bool) error
	StatusHistories(vehicleID int64) ([]domain.VehicleOrderStatusHistory, error)
}

type VehicleOrderService interface {
	GetByID(id int64) (*domain.VehicleOrderModel, error)
	GetList(req domain.FilterRequest) ([]domain.VehicleOrderModel, error)
	Create(req domain.CreateVehicleOrderRequest) error
	Update(req domain.UpdateVehicleOrderRequest) error
	StatusHistories(vehicleID int64) ([]domain.VehicleOrderStatusHistoryModel, error)
}

type VehicleOrderHandler interface {
	Create(c echo.Context) error
	Update(c echo.Context) error
	GetByID(c echo.Context) error
	GetList(c echo.Context) error
}
