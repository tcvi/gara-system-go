package vehicleorder

import (
	"garasystem/internal/core/domain"
	"github.com/pkg/errors"
)

func (s *Storage) StatusHistories(vehicleID int64) ([]domain.VehicleOrderStatusHistory, error) {
	var histories []domain.VehicleOrderStatusHistory

	err := s.db.Where("vehicle_order_id = ?", vehicleID).Order("created_at DESC").Find(&histories).Error

	return histories, errors.Wrap(err, "Get status histories fail")
}
