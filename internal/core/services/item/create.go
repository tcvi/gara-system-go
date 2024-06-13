package item

import (
	"garasystem/internal/core/domain"
	"garasystem/internal/core/myerror"
)

func validateCreateItem(req domain.CreateItemReq) error {
	if req.Description == "" || req.Name == "" || req.Price <= 0 || req.CategoryID <= 0 {
		return myerror.ErrItemInvalidData(nil)
	}

	return nil
}

func (u *Service) Create(req domain.CreateItemReq) error {
	err := validateCreateItem(req)
	if err != nil {
		return err
	}

	_, err = u.repo.CategoryStore.Get("id = ?", req.CategoryID)
	if err != nil {
		return myerror.ErrItemCategoryNotFound(err, req.CategoryID)
	}

	item := &domain.Item{
		CategoryID:  req.CategoryID,
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
	}

	err = u.repo.ItemStore.Create(item)
	if err != nil {
		return myerror.ErrItemCreate(err)
	}

	return nil
}
