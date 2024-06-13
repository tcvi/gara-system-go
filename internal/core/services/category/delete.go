package category

import (
	"garasystem/internal/core/myerror"
)

func validateDeleteCategory(id int64) error {
	if id <= 0 {
		return myerror.ErrCategoryDataInvalid(nil)
	}

	return nil
}

func (u *Service) Delete(id int64) error {
	err := validateDeleteCategory(id)
	if err != nil {
		return err
	}

	_, err = u.repo.CategoryStore.Get("id = ?", id)
	if err != nil {
		return myerror.ErrCategoryNotFound(err)
	}

	err = u.repo.CategoryStore.Delete(id)
	if err != nil {
		return myerror.ErrCategoryDelete(err)
	}

	return nil
}
