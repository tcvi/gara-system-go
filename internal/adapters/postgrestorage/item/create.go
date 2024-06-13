package item

import (
	"garasystem/internal/core/domain"
	"github.com/pkg/errors"
)

func (s *Storage) Create(item *domain.Item) error {
	err := s.db.Create(item).Error
	return errors.Wrap(err, "Create item fail")
}
