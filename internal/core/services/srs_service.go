package services

import "github.com/gcaldasl/srs-cli/internal/core/ports"

type SRSService struct {
	repo ports.CardRepository
}

func NewSRSService(repo ports.CardRepository) *SRSService {
	return &SRSService{repo: repo}
}
