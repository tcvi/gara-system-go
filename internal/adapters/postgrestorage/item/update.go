package item

import (
	"garasystem/internal/core/domain"
	"github.com/pkg/errors"
)

func (s *Storage) Update(item *domain.Item) error {
	err := s.db.Updates(item).Error
	return errors.Wrap(err, "Update item fail")
}
