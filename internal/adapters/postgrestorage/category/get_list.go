package category

import (
	"garasystem/internal/core/domain"
	"github.com/pkg/errors"
)

func (s *Storage) GetList(query interface{}, args ...interface{}) ([]domain.Category, error) {
	var categories []domain.Category

	db := s.db
	if query != nil && args != nil {
		db = db.Where(query, args)
	}

	err := db.Find(&categories).Error
	if err != nil {
		return nil, errors.Wrap(err, "GetList category fail")
	}

	return categories, nil
}
