package ports

import (
	"garasystem/internal/core/domain"
	"github.com/labstack/echo/v4"
)

type CategoryStore interface {
	GetList(query interface{}, args ...interface{}) ([]domain.Category, error)
	Create(*domain.Category) error
	Update(*domain.Category) error
	Delete(id int64) error
	Get(query interface{}, args ...interface{}) (*domain.Category, error)
}

type CategoryHandler interface {
	GetList(c echo.Context) error
	Create(c echo.Context) error
	Update(c echo.Context) error
	Delete(c echo.Context) error
}

type CategoryService interface {
	GetList() ([]domain.CategoryModel, error)
	Create(req domain.CreateCategoryReq) error
	Update(req domain.UpdateCategoryReq) error
	Delete(id int64) error
}
