package ports

import "github.com/gcaldasl/srs-cli/internal/core/domain"

type CardRepository interface {
	Create(card *domain.Card) error
	Get(id int64) (*domain.Card, error)
	Update(card *domain.Card) error
	Delete(id int64) error
	ListDue() ([]*domain.Card, error)
}
