package category

import (
	"garasystem/internal/core/domain"
	"github.com/pkg/errors"
)

func (s *Storage) Delete(id int64) error {
	err := s.db.Delete(domain.Category{ID: id}).Error
	return errors.Wrap(err, "Delete category fail")
}
