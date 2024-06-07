package vehicleorder

import (
	"garasystem/internal/core/domain"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func (s *Storage) Update(vehicleOrder *domain.VehicleOrder, isLogStatusHistory bool) error {
	err := s.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Updates(vehicleOrder).Error; err != nil {
			return err
		}

		if !isLogStatusHistory {
			return nil
		}

		history := &domain.VehicleOrderStatusHistory{
			VehicleOrderID: vehicleOrder.ID,
			Status:         vehicleOrder.Status,
		}

		if err := tx.Create(history).Error; err != nil {
			return err
		}

		return nil
	})
	return errors.Wrap(err, "Update VehicleOrder fail")
}
