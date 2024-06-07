package vehicleorder

import (
	"garasystem/internal/core/domain"
	"garasystem/internal/core/myerror"
)

func validateCreate(s *Service, req domain.CreateVehicleOrderRequest) error {
	if req.Note == "" || req.UserID <= 0 || req.HandlerID <= 0 {
		return myerror.ErrVehicleDataInvalid(nil)
	}

	_, err := s.userService.GetByID(req.UserID)
	if err != nil {
		return myerror.ErrVehicleUserNotFound(nil)
	}

	_, err = s.userService.GetByID(req.HandlerID)
	if err != nil {
		return myerror.ErrVehicleHandlerNotFound(nil)
	}

	return nil
}

func (u *Service) Create(req domain.CreateVehicleOrderRequest) error {
	err := validateCreate(u, req)
	if err != nil {
		return err
	}

	vehicleOrder := &domain.VehicleOrder{
		UserID:    req.UserID,
		HandlerID: req.HandlerID,
		Note:      req.Note,
		Status:    domain.New,
	}

	err = u.repo.VehicleStore.Create(vehicleOrder)
	if err != nil {
		return myerror.ErrVehicleCreate(err)
	}

	return nil
}
