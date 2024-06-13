package category

import (
	"garasystem/internal/core/domain"
	"garasystem/internal/core/myerror"
)

func (u *Service) GetList() ([]domain.CategoryModel, error) {
	categories, err := u.repo.CategoryStore.GetList(nil, nil)
	if err != nil {
		return nil, myerror.ErrCategoryCreate(err)
	}

	categoriesModel := make([]domain.CategoryModel, 0)
	for _, category := range categories {
		categoriesModel = append(categoriesModel, *category.MappingCategoryModel())
	}

	return categoriesModel, nil
}
