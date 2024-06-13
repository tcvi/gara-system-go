package category

import (
	"garasystem/internal/core/domain"
	"garasystem/internal/core/myerror"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func validateCreateCategory(req domain.CreateCategoryReq) error {
	if req.Name == "" {
		return myerror.ErrCategoryDataInvalid(nil)
	}

	return nil
}

func (u *Service) Create(req domain.CreateCategoryReq) error {
	err := validateCreateCategory(req)
	if err != nil {
		return err
	}

	category, err := u.repo.CategoryStore.Get("name = ?", req.Name)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return myerror.ErrCategoryGet(err)
	}

	if category != nil {
		return myerror.ErrCategoryExisted(nil)
	}

	newCategory := &domain.Category{
		Name: req.Name,
	}

	err = u.repo.CategoryStore.Create(newCategory)
	if err != nil {
		return myerror.ErrCategoryCreate(err)
	}

	return nil
}
