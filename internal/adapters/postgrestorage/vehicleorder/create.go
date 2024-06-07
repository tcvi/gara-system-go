package vehicleorder

import (
	"garasystem/internal/core/domain"
	"github.com/pkg/errors"
)

func (s *Storage) Create(vehicleOrder *domain.VehicleOrder) error {
	err := s.db.Create(vehicleOrder).Error
	return errors.Wrap(err, "Create VehicleOrder fail")
}
