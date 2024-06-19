package vehicleorderitem

import (
	"garasystem/internal/core/domain"
	"github.com/pkg/errors"
)

func (s *Storage) GetList(query interface{}, args ...interface{}) ([]domain.VehicleOrderItem, error) {
	var vehicleOrderItems []domain.VehicleOrderItem

	db := s.db
	if query != nil && args != nil {
		db = db.Where(query, args...)
	}

	err := db.Find(&vehicleOrderItems).Error
	if err != nil {
		return nil, errors.Wrap(err, "GetList vehicle order item fail")
	}

	return vehicleOrderItems, nil
}
