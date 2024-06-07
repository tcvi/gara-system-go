package vehicleorder

import (
	"fmt"
	"garasystem/internal/core/domain"
	"github.com/pkg/errors"
)

func (s *Storage) GetList(filter *domain.FilterRequest) ([]domain.VehicleOrder, error) {
	var vehicleOrders []domain.VehicleOrder

	db := s.db
	if filter != nil {
		query := make(map[string]interface{})
		fmt.Println("Oke ", filter.UserID)

		if filter.Status != "" {
			query["status"] = filter.Status
		}

		if filter.HandlerID > 0 {
			query["handler_id"] = filter.HandlerID
		}

		if filter.UserID > 0 {
			query["user_id"] = filter.UserID
		}

		if filter.SearchKey != "" {
			search := fmt.Sprint("%", filter.SearchKey, "%")
			subQuery := s.db.Select("id").Where("user_name LIKE ?", search).Table("users")
			db = db.Where("user_id IN (?) OR note LIKE ?", subQuery, search)
		}

		if len(query) > 0 {
			db = db.Where(query)
		}
	}

	err := db.Order("updated_at Desc").Find(&vehicleOrders).Error
	if err != nil {
		return nil, errors.Wrap(err, "Get list VehicleOrder fail")
	}

	return vehicleOrders, nil
}
