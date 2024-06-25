package myerror

import "net/http"

func ErrNotificationInvalidData(err error) MyError {
	return MyError{
		Raw:       err,
		HTTPCode:  http.StatusBadRequest,
		ErrorCode: 106000,
		Message:   "Invalid data",
	}
}

func ErrNotificationClientMessage(err error) MyError {
	return MyError{
		Raw:       err,
		HTTPCode:  http.StatusInternalServerError,
		ErrorCode: 106001,
		Message:   "Get client message fail",
	}
}

func ErrNotificationPush(err error) MyError {
	return MyError{
		Raw:       err,
		HTTPCode:  http.StatusInternalServerError,
		ErrorCode: 106002,
		Message:   "Push notification fail",
	}
}
