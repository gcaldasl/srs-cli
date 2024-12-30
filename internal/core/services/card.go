package services

import (
	"time"

	"github.com/gcaldasl/srs-cli/internal/core/domain"
	"github.com/gcaldasl/srs-cli/internal/core/ports"
)

type CardService struct {
	repo ports.CardRepository
	sm2  *domain.SM2Calculator
}

func NewCardService(repo ports.CardRepository) *CardService {
	return &CardService{
		repo: repo,
		sm2:  domain.NewSM2Calculator(),
	}
}

func (s *CardService) CreateCard(frontSide, backSide string) error {
	card := &domain.Card{
		FrontSide:    frontSide,
		BackSide:     backSide,
		LastReviewed: time.Now(),
		NextReview:   time.Now().Add(24 * time.Hour),
		Interval:     1,
		EaseFactor:   2.5,
	}

	return s.repo.Create(card)
}

func (s *CardService) ReviewCard(cardID int64, quality int) error {
	card, err := s.repo.Get(cardID)
	if err != nil {
		return err
	}

	s.sm2.Calculate(card, quality)
	return s.repo.Update(card)
}

func (s *CardService) ListDueCards() ([]*domain.Card, error) {
	return s.repo.ListDue()
}
