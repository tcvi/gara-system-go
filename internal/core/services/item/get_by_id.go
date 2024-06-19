package item

import (
	"garasystem/internal/core/domain"
	"garasystem/internal/core/myerror"
)

func (u *Service) GetById(ID int64) (*domain.ItemModel, error) {
	if ID <= 0 {
		return nil, myerror.ErrItemParamIDInvalid(nil)
	}

	item, err := u.repo.ItemStore.Get("id = ?", ID)
	if err != nil {
		return nil, myerror.ErrItemItemNotFound(err)
	}

	itemModel := item.MappingItemModel()

	category, err := u.repo.CategoryStore.Get("id = ?", item.CategoryID)
	if err != nil {
		return nil, myerror.ErrItemCategoryNotFound(err, item.CategoryID)
	}

	itemModel.Category = *category.MappingCategoryModel()

	return itemModel, nil
}
