package item

import (
	"garasystem/internal/core/domain"
	"github.com/pkg/errors"
)

func (s *Storage) GetList(query interface{}, args ...interface{}) ([]domain.Item, error) {
	var items []domain.Item
	db := s.db

	if query != nil && args != nil {
		db = db.Where(query, args...)
	}

	err := db.Find(&items).Error
	return items, errors.Wrap(err, "Get list vehicle order items fail")
}
