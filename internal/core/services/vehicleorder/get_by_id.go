package vehicleorder

import (
	"garasystem/internal/core/domain"
	"garasystem/internal/core/myerror"
)

func (u *Service) GetByID(id int64) (*domain.VehicleOrderModel, error) {
	vehicleOrder, err := u.repo.VehicleStore.GetByID(id)
	if err != nil {
		return nil, myerror.ErrVehicleOrderGet(nil)
	}

	vehicleOrderModel := vehicleOrder.MappingVehicleOrderModel()

	user, err := u.userService.GetByID(vehicleOrder.UserID)
	if err != nil {
		return nil, myerror.ErrVehicleOrderGetUser(err)
	}
	vehicleOrderModel.User = *user

	handler, err := u.userService.GetByID(vehicleOrder.UserID)
	if err != nil {
		return nil, myerror.ErrVehicleOrderGetHandler(err)
	}
	vehicleOrderModel.Handler = *handler

	return vehicleOrderModel, nil
}
