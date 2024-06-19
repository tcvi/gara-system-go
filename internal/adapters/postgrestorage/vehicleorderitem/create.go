package vehicleorderitem

import (
	"garasystem/internal/core/domain"
	"github.com/pkg/errors"
)

func (s *Storage) Create(items []domain.VehicleOrderItem) error {
	err := s.db.Create(items).Error
	return errors.Wrap(err, "Add vehicle order's items fail")
}
