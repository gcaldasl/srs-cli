package services

import (
	"time"

	"github.com/gcaldasl/srs-cli/internal/core/domain"
	"github.com/gcaldasl/srs-cli/internal/core/ports"
)

type CardService struct {
	repo ports.CardRepository
}

func NewCardService(repo ports.CardRepository) *CardService {
	return &CardService{repo: repo}
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

	if quality >= 3 {
		if card.Interval == 0 {
			card.Interval = 1
		} else if card.Interval == 1 {
			card.Interval = 6
		} else {
			card.Interval = int(float64(card.Interval) * card.EaseFactor)
		}
	} else {
		card.Interval = 1
	}

	card.EaseFactor = card.EaseFactor + (0.1 - (5-float64(quality))*(0.08+(5-float64(quality))*0.02))
	if card.EaseFactor < 1.3 {
		card.EaseFactor = 1.3
	}

	card.LastReviewed = time.Now()
	card.NextReview = time.Now().AddDate(0, 0, card.Interval)

	return s.repo.Update(card)
}

func (s *CardService) ListDueCards() ([]*domain.Card, error) {
	return s.repo.ListDue()
}
