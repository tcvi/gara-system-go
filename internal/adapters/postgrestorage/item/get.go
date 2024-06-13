package item

import "garasystem/internal/core/domain"

func (s *Storage) Get(query interface{}, args ...interface{}) (*domain.Item, error) {
	var item domain.Item

	db := s.db
	if query != nil && args != nil {
		db = db.Where(query, args)
	}

	err := db.First(&item).Error
	if err != nil {
		return nil, err
	}

	return &item, nil
}
