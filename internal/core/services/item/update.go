package item

import (
	"garasystem/internal/core/domain"
	"garasystem/internal/core/myerror"
)

func validateUpdateItem(req domain.UpdateItemReq) error {
	if req.ID <= 0 {
		return myerror.ErrItemParamIDInvalid(nil)
	}

	if req.Description == "" && req.Name == "" && req.Price <= 0 && req.CategoryID <= 0 {
		return myerror.ErrItemInvalidData(nil)
	}

	return nil
}

func (u *Service) Update(req domain.UpdateItemReq) error {
	err := validateUpdateItem(req)
	if err != nil {
		return err
	}

	item, err := u.repo.ItemStore.Get("id = ?", req.ID)
	if err != nil {
		return myerror.ErrItemItemNotFound(err)
	}

	if req.CategoryID > 0 {
		_, err := u.repo.CategoryStore.Get("id = ?", req.CategoryID)
		if err != nil {
			return myerror.ErrItemCategoryNotFound(err, req.CategoryID)
		}

		item.CategoryID = req.CategoryID
	}

	if req.Price > 0 {
		item.Price = req.Price
	}

	if req.Name != "" {
		item.Name = req.Name
	}

	if req.Description != "" {
		item.Description = req.Description
	}

	err = u.repo.ItemStore.Update(item)
	if err != nil {
		return myerror.ErrItemUpdate(err)
	}

	return nil
}
