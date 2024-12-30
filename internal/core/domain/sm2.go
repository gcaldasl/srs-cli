package domain

import "time"

type SM2Calculator struct{}

func NewSM2Calculator() *SM2Calculator {
	return &SM2Calculator{}
}

func (s *SM2Calculator) Calculate(card *Card, quality int) {
	if quality >= 3 {
		card.Interval = s.calculateInterval(card)
	} else {
		card.Interval = 1
	}

	card.EaseFactor = s.calculateEaseFactor(card.EaseFactor, quality)
	s.updateReviewDates(card)
}

func (s *SM2Calculator) calculateInterval(card *Card) int {
	switch card.Interval {
	case 0:
		return 1
	case 1:
		return 6
	default:
		return int(float64(card.Interval) * card.EaseFactor)
	}
}

func (s *SM2Calculator) calculateEaseFactor(currentEF float64, quality int) float64 {
	newEF := currentEF + (0.1 - (5-float64(quality))*(0.08+(5-float64(quality))*0.02))
	if newEF < 1.3 {
		return 1.3
	}
	return newEF
}

func (s *SM2Calculator) updateReviewDates(card *Card) {
	card.LastReviewed = time.Now()
	card.NextReview = time.Now().AddDate(0, 0, card.Interval)
}
