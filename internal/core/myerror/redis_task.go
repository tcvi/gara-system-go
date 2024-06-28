package myerror

import "net/http"

func ErrRedisTaskMarshalPushNotificationPayload(err error) MyError {
	return MyError{
		Raw:       err,
		HTTPCode:  http.StatusInternalServerError,
		ErrorCode: 107000,
		Message:   "failed marshal push notification payload",
	}
}

func ErrRedisTaskEnqueue(err error) MyError {
	return MyError{
		Raw:       err,
		HTTPCode:  http.StatusInternalServerError,
		ErrorCode: 107001,
		Message:   "failed enqueue task",
	}
}
