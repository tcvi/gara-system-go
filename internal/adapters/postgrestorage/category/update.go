package category

import (
	"garasystem/internal/core/domain"
	"github.com/pkg/errors"
)

func (s *Storage) Update(category *domain.Category) error {
	err := s.db.Updates(category).Error
	return errors.Wrap(err, "Update Category fail")
}
