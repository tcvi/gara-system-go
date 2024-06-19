package myerror

import (
	"fmt"
	"net/http"
)

func ErrVehicleOrderItemGetList(err error) MyError {
	return MyError{
		Raw:       err,
		HTTPCode:  http.StatusInternalServerError,
		ErrorCode: 105001,
		Message:   "Get vehicle order items fail",
	}
}

func ErrVehicleOrderItemGetVehicleOrder(err error) MyError {
	return MyError{
		Raw:       err,
		HTTPCode:  http.StatusBadRequest,
		ErrorCode: 105002,
		Message:   "Get vehicle order not found",
	}
}

func ErrVehicleOrderItemInvalidData(err error) MyError {
	return MyError{
		Raw:       err,
		HTTPCode:  http.StatusBadRequest,
		ErrorCode: 105003,
		Message:   "Invalid data",
	}
}

func ErrVehicleOrderItemDuplicateItem(err error, itemID int64) MyError {
	return MyError{
		Raw:       err,
		HTTPCode:  http.StatusBadRequest,
		ErrorCode: 105004,
		Message:   fmt.Sprintf("Duplicate item id = %d", itemID),
	}
}

func ErrVehicleOrderItemCreate(err error) MyError {
	return MyError{
		Raw:       err,
		HTTPCode:  http.StatusBadRequest,
		ErrorCode: 105005,
		Message:   "Create vehicle order items fail",
	}
}

func ErrVehicleOrderItemNotFound(err error, id int64) MyError {
	return MyError{
		Raw:       err,
		HTTPCode:  http.StatusBadRequest,
		ErrorCode: 105005,
		Message:   fmt.Sprintf("Item id = %d not found", id),
	}
}

func ErrVehicleOrderItemGetItems(err error) MyError {
	return MyError{
		Raw:       err,
		HTTPCode:  http.StatusBadRequest,
		ErrorCode: 105006,
		Message:   "Get items fail",
	}
}

func ErrVehicleOrderItemUpdate(err error) MyError {
	return MyError{
		Raw:       err,
		HTTPCode:  http.StatusBadRequest,
		ErrorCode: 105007,
		Message:   "Update vehicle order items fail",
	}
}
