package util

import (
	"garasystem/internal/core/myerror"
	"garasystem/internal/logger"
	"github.com/labstack/echo/v4"
	"net/http"
)

type response struct {
}

var Response response

func (response) Success(c echo.Context, data interface{}) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    http.StatusOK,
		"message": "OK",
		"data":    data,
	})
}

func (response) Error(c echo.Context, err myerror.MyError) error {
	logger.Log.Error(err)
	//sentry.WithContext(&c).Error(err.Raw)

	return c.JSON(err.HTTPCode, map[string]interface{}{
		"code":    err.ErrorCode,
		"message": err.Message,
	})
}
