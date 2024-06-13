package category

import "garasystem/internal/core/domain"

func (s *Storage) Get(query interface{}, args ...interface{}) (*domain.Category, error) {
	var category domain.Category

	err := s.db.Where(query, args).First(&category).Error
	if err != nil {
		return nil, err
	}

	return &category, nil
}
