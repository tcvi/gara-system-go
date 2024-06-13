package category

import (
	"garasystem/internal/core/domain"
	"garasystem/internal/core/myerror"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func validateUpdateCategory(req domain.UpdateCategoryReq) error {
	if req.Name == "" || req.ID <= 0 {
		return myerror.ErrCategoryDataInvalid(nil)
	}

	return nil
}

func (u *Service) Update(req domain.UpdateCategoryReq) error {
	err := validateUpdateCategory(req)
	if err != nil {
		return err
	}

	category, err := u.repo.CategoryStore.Get("id = ?", req.ID)
	if err != nil {
		return myerror.ErrCategoryNotFound(err)
	}

	categoryName, err := u.repo.CategoryStore.Get("name = ?", req.Name)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return myerror.ErrCategoryGet(err)
	}

	if categoryName != nil {
		return myerror.ErrCategoryExisted(nil)
	}

	category.Name = req.Name
	err = u.repo.CategoryStore.Update(category)
	if err != nil {
		return myerror.ErrCategoryUpdate(err)
	}

	return nil
}
