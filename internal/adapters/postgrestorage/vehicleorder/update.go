package vehicleorder

import (
	"garasystem/internal/core/domain"
	"github.com/pkg/errors"
)

func (s *Storage) Update(vehicleOrder *domain.VehicleOrder) error {
	err := s.db.Updates(vehicleOrder).Error
	return errors.Wrap(err, "Update VehicleOrder fail")
}
