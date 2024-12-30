package domain

import "time"

type SM2Calculator struct{}

func NewSM2Calculator() *SM2Calculator {
	return &SM2Calculator{}
}

func (s *SM2Calculator) Calculate(card *Card, quality int) {
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
}
