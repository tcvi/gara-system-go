package vehicleorderitem

import (
	"garasystem/internal/core/domain"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func (s *Storage) Updates(vehicleOderID int64, items []domain.VehicleOrderItem) error {
	err := s.db.Transaction(func(tx *gorm.DB) error {
		itemIds := make([]int64, 0)

		for _, item := range items {
			if item.ID > 0 {
				itemIds = append(itemIds, item.ID)
			}
		}

		err := tx.
			Where("vehicle_order_id = ? AND id NOT IN (?)", vehicleOderID, itemIds).
			Delete(&domain.VehicleOrderItem{}).Error
		if err != nil {
			return err
		}

		for _, item := range items {
			err := tx.Where("id = ?", item.ID).Assign(item).FirstOrCreate(&item).Error
			if err != nil {
				return err
			}
		}

		return nil
	})

	return errors.Wrap(err, "Update items fail")
}
