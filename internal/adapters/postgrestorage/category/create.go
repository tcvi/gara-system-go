package category

import (
	"garasystem/internal/core/domain"
	"github.com/pkg/errors"
)

func (s *Storage) Create(category *domain.Category) error {
	err := s.db.Create(category).Error
	return errors.Wrap(err, "Create Category fail")
}
