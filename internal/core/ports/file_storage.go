package ports

import (
	"garasystem/internal/core/domain"
	"github.com/labstack/echo/v4"
	"mime/multipart"
	"time"
)

type FileStorage interface {
	UploadFile(file *multipart.FileHeader) (*domain.FileUpload, error)
	PresignObject(objectKey string, expire time.Duration) (string, error)
}

type FileHandler interface {
	Upload(c echo.Context) error
}

type FileService interface {
	Upload([]*multipart.FileHeader) ([]domain.FileUpload, error)
}
