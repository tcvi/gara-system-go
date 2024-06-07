package vehicleorder

import (
	"garasystem/internal/core/domain"
	"garasystem/internal/core/myerror"
)

func (u *Service) GetList(req domain.FilterRequest) ([]domain.VehicleOrderModel, error) {
	vehicleOrders, err := u.repo.VehicleStore.GetList(&req)
	if err != nil {
		return nil, myerror.ErrVehicleGetList(err)
	}

	var (
		userMap = make(map[int64]domain.UserModel)
		userIDs = make([]int64, 0)
	)

	for _, v := range vehicleOrders {
		_, ok := userMap[v.UserID]
		if !ok {
			userIDs = append(userIDs, v.UserID)
			userMap[v.UserID] = domain.UserModel{}
		}

		_, ok = userMap[v.HandlerID]
		if !ok {
			userIDs = append(userIDs, v.HandlerID)
			userMap[v.HandlerID] = domain.UserModel{}
		}
	}

	users, err := u.repo.UserStore.GetList("id IN (?)", userIDs)
	if err != nil {
		return nil, myerror.ErrVehicleGetUsers(err)
	}

	for _, user := range users {
		userMap[user.ID] = *user.MappingUserModel()
	}

	vehicleModel := make([]domain.VehicleOrderModel, 0)

	for _, v := range vehicleOrders {
		model := *v.MappingVehicleOrderModel()

		user, ok := userMap[v.UserID]
		if !ok || user.ID <= 0 {
			return nil, myerror.ErrVehicleUserNotFound(nil)
		}
		model.User = user

		handler, ok := userMap[v.HandlerID]
		if !ok || handler.ID <= 0 {
			return nil, myerror.ErrVehicleHandlerNotFound(nil)
		}
		model.Handler = handler

		vehicleModel = append(vehicleModel, model)
	}

	return vehicleModel, nil
}
