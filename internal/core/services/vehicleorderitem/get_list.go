package vehicleorderitem

import (
	"fmt"
	"garasystem/internal/core/domain"
	"garasystem/internal/core/myerror"
)

func (s *Service) GetList(vehicleOderID int64) ([]domain.VehicleOrderItemModel, error) {
	_, err := s.repo.VehicleStore.GetByID(vehicleOderID)
	if err != nil {
		return nil, myerror.ErrVehicleOrderItemGetVehicleOrder(err)
	}

	orderItems, err := s.repo.VehicleItemStore.GetList("vehicle_order_id = ? ", vehicleOderID)
	if err != nil {
		return nil, myerror.ErrVehicleOrderItemGetList(err)
	}

	fmt.Print("++++ ", len(orderItems))

	items := make([]domain.VehicleOrderItemModel, 0)

	for _, item := range orderItems {
		itemDetail, err := s.itemService.GetById(item.ItemID)
		if err != nil {
			return nil, err
		}

		itemModel := *item.MappingVehicleOrderItemModel()
		itemModel.Item = *itemDetail

		items = append(items, itemModel)
	}

	return items, nil
}
