package vehicleorder

import (
	"garasystem/internal/core/domain"
	"garasystem/internal/core/myerror"
)

func validate(req domain.UpdateVehicleOrderRequest) error {
	if req.Note == "" || req.HandlerID <= 0 {
		return myerror.ErrVehicleDataInvalid(nil)
	}

	return nil
}

func (s *Service) Update(req domain.UpdateVehicleOrderRequest) error {
	if err := validate(req); err != nil {
		return err
	}

	vehicleStore, err := s.repo.VehicleStore.GetByID(req.ID)
	if err != nil {
		return myerror.ErrVehicleOrderGet(err)
	}

	status, err := domain.ParseStatus(req.Status)
	if err != nil {
		return myerror.ErrVehicleStatusInvalid(err)
	}

	isLogStatusHistory := *status != vehicleStore.Status

	vehicleStore.HandlerID = req.HandlerID
	vehicleStore.Note = req.Note
	vehicleStore.Status = *status

	return s.repo.VehicleStore.Update(vehicleStore, isLogStatusHistory)
}
