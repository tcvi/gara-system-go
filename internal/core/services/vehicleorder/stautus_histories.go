package vehicleorder

import (
	"garasystem/internal/core/domain"
	"garasystem/internal/core/myerror"
)

func (s *Service) StatusHistories(vehicleID int64) ([]domain.VehicleOrderStatusHistoryModel, error) {
	histories, err := s.repo.VehicleStore.StatusHistories(vehicleID)
	if err != nil {
		return nil, myerror.ErrVehicleGetStatusHistories(err)
	}

	modelHistories := make([]domain.VehicleOrderStatusHistoryModel, 0)

	if len(histories) == 0 {
		return modelHistories, nil
	}

	for _, history := range histories {
		modelHistories = append(modelHistories, *history.MappingHistoryVehicleOrderStatusModel())
	}

	return modelHistories, nil
}
