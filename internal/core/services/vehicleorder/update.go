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

func (u *Service) Update(req domain.UpdateVehicleOrderRequest) error {
	if err := validate(req); err != nil {
		return err
	}

	vehicleStore, err := u.repo.VehicleStore.GetByID(req.ID)
	if err != nil {
		return myerror.ErrVehicleOrderGet(err)
	}

	status, err := domain.ParseStatus(req.Status)
	if err != nil {
		return myerror.ErrVehicleStatusInvalid(err)
	}

	vehicleStore.HandlerID = req.HandlerID
	vehicleStore.Note = req.Note
	vehicleStore.Status = *status

	return u.repo.VehicleStore.Update(vehicleStore)
}
