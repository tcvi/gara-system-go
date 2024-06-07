package myerror

import "net/http"

func ErrVehicleOrderGet(err error) MyError {
	return MyError{
		Raw:       err,
		HTTPCode:  http.StatusBadRequest,
		ErrorCode: 102001,
		Message:   "Not found vehicle order",
	}
}

func ErrVehicleOrderGetUser(err error) MyError {
	return MyError{
		Raw:       err,
		HTTPCode:  http.StatusBadRequest,
		ErrorCode: 102002,
		Message:   "Not found user",
	}
}

func ErrVehicleOrderGetHandler(err error) MyError {
	return MyError{
		Raw:       err,
		HTTPCode:  http.StatusBadRequest,
		ErrorCode: 102003,
		Message:   "Not found handler",
	}
}

func ErrVehicleCreate(err error) MyError {
	return MyError{
		Raw:       err,
		HTTPCode:  http.StatusInternalServerError,
		ErrorCode: 102004,
		Message:   "Create Vehicle order fail",
	}
}

func ErrVehicleStatusInvalid(err error) MyError {
	return MyError{
		Raw:       err,
		HTTPCode:  http.StatusInternalServerError,
		ErrorCode: 102005,
		Message:   "Status invalid",
	}
}

func ErrVehicleDataInvalid(err error) MyError {
	return MyError{
		Raw:       err,
		HTTPCode:  http.StatusInternalServerError,
		ErrorCode: 102006,
		Message:   "Data invalid",
	}
}

func ErrVehicleGetList(err error) MyError {
	return MyError{
		Raw:       err,
		HTTPCode:  http.StatusInternalServerError,
		ErrorCode: 102007,
		Message:   "Get Vehicle Orders fail",
	}
}

func ErrVehicleGetUsers(err error) MyError {
	return MyError{
		Raw:       err,
		HTTPCode:  http.StatusInternalServerError,
		ErrorCode: 102008,
		Message:   "Get users fail",
	}
}

func ErrVehicleUserNotFound(err error) MyError {
	return MyError{
		Raw:       err,
		HTTPCode:  http.StatusInternalServerError,
		ErrorCode: 102009,
		Message:   "User not found",
	}
}

func ErrVehicleHandlerNotFound(err error) MyError {
	return MyError{
		Raw:       err,
		HTTPCode:  http.StatusInternalServerError,
		ErrorCode: 102009,
		Message:   "Handler not found",
	}
}

func ErrVehicleInvalidIDParam(err error) MyError {
	return MyError{
		Raw:       err,
		HTTPCode:  http.StatusBadRequest,
		ErrorCode: 102010,
		Message:   "Invalid id param",
	}
}
