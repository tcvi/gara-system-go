package file

import (
	"garasystem/internal/core/myerror"
	"garasystem/pkg/util"
	"github.com/labstack/echo/v4"
)

func (h *Handler) Upload(c echo.Context) error {
	form, err := c.MultipartForm()
	if err != nil {
		return util.Response.Error(c, myerror.ErrFileMissingMultipartForm(err))
	}

	files := form.File["files"]
	if len(files) == 0 {
		return util.Response.Error(c, myerror.ErrFileMissingFile(err))
	}

	filesResponse, err := h.service.Upload(files)
	if err != nil {
		return util.Response.Error(c, err.(myerror.MyError))
	}

	return util.Response.Success(c, filesResponse)
}
