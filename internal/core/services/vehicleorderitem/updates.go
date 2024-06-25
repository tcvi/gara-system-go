package vehicleorderitem

import (
	"garasystem/internal/core/domain"
	"garasystem/internal/core/myerror"
	"time"
)

func validateUpdateOrderItems(s *Service, orderItems []domain.UpdateVehicleOrderItem, records []domain.VehicleOrderItem) error {
	var (
		updateOrderItemIds = map[int64]bool{}
		recordIds          = map[int64]bool{}
	)

	for _, record := range records {
		recordIds[record.ID] = true
	}

	for _, item := range orderItems {
		if item.ItemID <= 0 || item.Price <= 0 || item.Quantity <= 0 {
			return myerror.ErrVehicleOrderItemInvalidData(nil)
		}

		if updateOrderItemIds[item.ItemID] {
			return myerror.ErrVehicleOrderItemDuplicateItem(nil, item.ItemID)
		}

		updateOrderItemIds[item.ItemID] = true

		if item.ID > 0 && !recordIds[item.ID] {
			return myerror.ErrVehicleOrderItemNotFound(nil, item.ID)
		}
	}

	// Check item exist ?
	keys := make([]int64, 0, len(updateOrderItemIds))

	for key := range updateOrderItemIds {
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

func (s *Service) Updates(req domain.UpdateVehicleOrderItemRequest) error {
	_, err := s.repo.VehicleStore.GetByID(req.VehicleOrderID)
	if err != nil {
		return myerror.ErrVehicleOrderItemGetVehicleOrder(err)
	}

	items, err := s.repo.VehicleItemStore.GetList("vehicle_order_id = ? ", req.VehicleOrderID)
	if err != nil {
		return myerror.ErrVehicleOrderItemGetList(err)
	}

	err = validateUpdateOrderItems(s, req.Items, items)
	if err != nil {
		return err
	}

	var (
		itemMap = map[int64]domain.VehicleOrderItem{}
		objs    = make([]domain.VehicleOrderItem, 0)
	)

	for _, item := range items {
		itemMap[item.ID] = item
	}

	for _, item := range req.Items {
		if item.ID <= 0 {
			obj := domain.VehicleOrderItem{
				ItemID:        item.ItemID,
				VehicleOderID: req.VehicleOrderID,
				Note:          item.Note,
				Price:         item.Price,
				Quantity:      item.Quantity,
			}
			objs = append(objs, obj)
		} else {
			record, ok := itemMap[item.ID]
			if !ok {
				return myerror.ErrVehicleOrderItemNotFound(nil, item.ID)
			}

			record.ItemID = item.ItemID
			record.Note = item.Note
			record.Quantity = item.Quantity
			record.Price = item.Price
			record.UpdatedAt = time.Now()

			objs = append(objs, record)
		}
	}

	err = s.repo.VehicleItemStore.Updates(req.VehicleOrderID, objs)
	if err != nil {
		return myerror.ErrVehicleOrderItemUpdate(err)
	}

	return nil
}
