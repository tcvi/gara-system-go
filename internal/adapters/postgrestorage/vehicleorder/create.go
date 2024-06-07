package vehicleorder

import (
	"garasystem/internal/core/domain"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func (s *Storage) Create(vehicleOrder *domain.VehicleOrder) error {
	err := s.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(vehicleOrder).Error; err != nil {
			return err
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

	return errors.Wrap(err, "Create VehicleOrder fail")
}
