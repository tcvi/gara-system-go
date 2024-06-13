package item

import (
	"fmt"
	"garasystem/internal/core/domain"
	"github.com/pkg/errors"
)

func (s *Storage) GetList(filter *domain.FilterItemRequest) ([]domain.Item, error) {
	db := s.db

	if filter != nil {
		query := make(map[string]interface{})

		if filter.CategoryID > 0 {
			query["category_id"] = filter.CategoryID
		}

		if filter.SearchKey != "" {
			searchKey := fmt.Sprint("%", filter.SearchKey, "%")
			db = db.Where("name like (?) OR description LIKE ?", searchKey, searchKey)
		}

		if len(query) > 0 {
			db = db.Where(query)
		}
	}

	var items []domain.Item
	err := db.Find(&items).Error

	return items, errors.Wrap(err, "Get list item fail")
}
