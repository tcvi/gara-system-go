package vehicleorder

import (
	"garasystem/internal/core/domain"
)

func (s *Storage) GetByID(id int64) (*domain.VehicleOrder, error) {
	var vehicleOrder domain.VehicleOrder

	err := s.db.Where("id = ?", id).First(&vehicleOrder).Error
	if err != nil {
		return nil, err
	}

	return &vehicleOrder, nil
}
