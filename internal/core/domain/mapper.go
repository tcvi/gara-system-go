package domain

func (u *User) MappingUserLogin() *UserLoginResponse {
	return &UserLoginResponse{
		UserModel: UserModel{
			ID:          u.ID,
			PhoneNumber: u.PhoneNumber,
			UserName:    u.UserName,
			CreatedAt:   u.CreatedAt,
			UpdatedAt:   u.UpdatedAt,
		},
	}
}

func (u *User) MappingUserModel() *UserModel {
	return &UserModel{
		ID:          u.ID,
		PhoneNumber: u.PhoneNumber,
		UserName:    u.UserName,
		CreatedAt:   u.CreatedAt,
		UpdatedAt:   u.UpdatedAt,
	}
}

func (v *VehicleOrder) MappingVehicleOrderModel() *VehicleOrderModel {
	return &VehicleOrderModel{
		ID:        v.ID,
		Note:      v.Note,
		User:      UserModel{ID: v.UserID},
		Handler:   UserModel{ID: v.HandlerID},
		Status:    v.Status,
		CreatedAt: v.CreatedAt,
		UpdatedAt: v.UpdatedAt,
	}
}

func (h *VehicleOrderStatusHistory) MappingHistoryVehicleOrderStatusModel() *VehicleOrderStatusHistoryModel {
	return &VehicleOrderStatusHistoryModel{
		ID:             h.ID,
		VehicleOrderID: h.VehicleOrderID,
		Status:         h.Status,
		CreatedAt:      h.CreatedAt,
		UpdatedAt:      h.UpdatedAt,
	}
}

func (c *Category) MappingCategoryModel() *CategoryModel {
	return &CategoryModel{
		ID:        c.ID,
		Name:      c.Name,
		CreatedAt: c.CreatedAt,
		UpdatedAt: c.UpdatedAt,
	}
}

func (i *Item) MappingItemModel() *ItemModel {
	return &ItemModel{
		ID:          i.ID,
		Name:        i.Name,
		Description: i.Description,
		Price:       i.Price,
		CreatedAt:   i.CreatedAt,
		UpdatedAt:   i.UpdatedAt,
	}
}

func (v *VehicleOrderItem) MappingVehicleOrderItemModel() *VehicleOrderItemModel {
	return &VehicleOrderItemModel{
		ID:        v.ID,
		Item:      ItemModel{ID: v.ItemID},
		Note:      v.Note,
		Price:     v.Price,
		Quantity:  v.Quantity,
		CreatedAt: v.CreatedAt,
		UpdatedAt: v.UpdatedAt,
	}
}
