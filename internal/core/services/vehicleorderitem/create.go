package vehicleorderitem

import (
	"garasystem/internal/core/domain"
	"garasystem/internal/core/myerror"
)

func validateCreateOrderItems(s *Service, orderItems []domain.VehicleOrderItemReq, records []domain.VehicleOrderItem) error {
	if len(orderItems) == 0 {
		return myerror.ErrVehicleOrderItemInvalidData(nil)
	}

	orderItemIds := map[int64]bool{}

	for _, item := range orderItems {
		if item.ItemID <= 0 || item.Price <= 0 || item.Quantity <= 0 {
			return myerror.ErrVehicleOrderItemInvalidData(nil)
		}

		if orderItemIds[item.ItemID] {
			return myerror.ErrVehicleOrderItemDuplicateItem(nil, item.ItemID)
		}

		orderItemIds[item.ItemID] = true
	}

	for _, record := range records {
		if orderItemIds[record.ItemID] {
			return myerror.ErrVehicleOrderItemDuplicateItem(nil, record.ItemID)
		}
	}

	// Check item exist ?
	keys := make([]int64, 0, len(orderItemIds))

	for key, _ := range orderItemIds {
		keys = append(keys, key)
	}

	items, err := s.repo.ItemStore.GetList("id IN (?)", keys)
	if err != nil {
		return myerror.ErrVehicleOrderItemGetItems(err)
	}

	itemIds := map[int64]bool{}

	for _, item := range items {
		itemIds[item.ID] = true
	}

	for _, item := range orderItems {
		if !itemIds[item.ItemID] {
			return myerror.ErrVehicleOrderItemNotFound(nil, item.ItemID)
		}
	}

	return nil
}

func (s *Service) Create(req domain.CreateVehicleOrderItemRequest) error {
	_, err := s.repo.VehicleStore.GetByID(req.VehicleOrderID)
	if err != nil {
		return myerror.ErrVehicleOrderItemGetVehicleOrder(err)
	}

	items, err := s.repo.VehicleItemStore.GetList("vehicle_order_id = ? ", req.VehicleOrderID)
	if err != nil {
		return myerror.ErrVehicleOrderItemGetList(err)
	}

	err = validateCreateOrderItems(s, req.Items, items)
	if err != nil {
		return err
	}

	objs := make([]domain.VehicleOrderItem, 0)

	for _, item := range req.Items {
		obj := domain.VehicleOrderItem{
			ItemID:        item.ItemID,
			VehicleOderID: req.VehicleOrderID,
			Note:          item.Note,
			Price:         item.Price,
			Quantity:      item.Quantity,
		}

		objs = append(objs, obj)
	}

	err = s.repo.VehicleItemStore.Create(objs)
	if err != nil {
		return myerror.ErrVehicleOrderItemCreate(err)
	}

	return nil
}
