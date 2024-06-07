package vehicleorder

import (
	"garasystem/internal/core/domain"
	"garasystem/internal/core/myerror"
)

func (s *Service) GetByID(id int64) (*domain.VehicleOrderModel, error) {
	vehicleOrder, err := s.repo.VehicleStore.GetByID(id)
	if err != nil {
		return nil, myerror.ErrVehicleOrderGet(nil)
	}

	vehicleOrderModel := vehicleOrder.MappingVehicleOrderModel()

	user, err := s.userService.GetByID(vehicleOrder.UserID)
	if err != nil {
		return nil, myerror.ErrVehicleOrderGetUser(err)
	}
	vehicleOrderModel.User = *user

	handler, err := s.userService.GetByID(vehicleOrder.UserID)
	if err != nil {
		return nil, myerror.ErrVehicleOrderGetHandler(err)
	}
	vehicleOrderModel.Handler = *handler

	histories, err := s.StatusHistories(vehicleOrder.ID)
	if err != nil {
		return nil, err
	}
	vehicleOrderModel.HistoryStatuses = histories

	return vehicleOrderModel, nil
}
