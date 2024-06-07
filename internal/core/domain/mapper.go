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
