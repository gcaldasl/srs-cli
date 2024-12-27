package domain

import "time"

type Card struct {
	ID           int64
	FrontSide    string
	BackSide     string
	LastReviewed time.Time
	NextReview   time.Time
	Interval     int
	EaseFactor   float64
}
