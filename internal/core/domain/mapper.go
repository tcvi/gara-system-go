package domain

func (u *User) MappingUserLogin() *UserLoginResponse {
	return &UserLoginResponse{
		ID:          u.ID,
		PhoneNumber: u.PhoneNumber,
		UserName:    u.UserName,
		CreatedAt:   u.CreatedAt,
		UpdatedAt:   u.UpdatedAt,
	}
}
