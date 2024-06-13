package item

import (
	"garasystem/internal/core/domain"
	"garasystem/internal/core/myerror"
)

func (u *Service) GetList(req *domain.FilterItemRequest) ([]domain.ItemModel, error) {
	items, err := u.repo.ItemStore.GetList(req)
	if err != nil {
		return nil, myerror.ErrItemGetList(err)
	}

	var (
		categoriesMap = map[int64]domain.Category{}
		categoryIds   = make([]int64, 0)
		itemsModel    = make([]domain.ItemModel, 0)
	)

	for _, item := range items {
		categoriesMap[item.CategoryID] = domain.Category{ID: item.CategoryID}
	}

	for _, category := range categoriesMap {
		categoryIds = append(categoryIds, category.ID)
	}

	categories, err := u.repo.CategoryStore.GetList("id IN ?", categoryIds)
	if err != nil {
		return nil, myerror.ErrItemGetCategory(err)
	}

	for _, category := range categories {
		categoriesMap[category.ID] = category
	}

	for _, item := range items {
		itemModel := *item.MappingItemModel()

		category, ok := categoriesMap[item.CategoryID]
		if !ok || category.ID <= 0 {
			return nil, myerror.ErrItemCategoryNotFound(nil, item.CategoryID)
		}

		itemModel.Category = *category.MappingCategoryModel()

		itemsModel = append(itemsModel, itemModel)
	}

	return itemsModel, nil
}
