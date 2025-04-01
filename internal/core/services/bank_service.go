package services

import (
	"ams-service/internal/core/entities"
	"ams-service/internal/ports/primary"
	"ams-service/internal/ports/secondary"

	"github.com/rs/zerolog/log"
)

type BankService struct {
	repo secondary.BankRepository
}

func NewBankService(repo secondary.BankRepository) primary.BankService {
	return &BankService{repo: repo}
}

func (s *BankService) AddCreditCard(card entities.CreditCard) error {
	err := s.repo.AddCreditCard(card)
	if err != nil {
		log.Error().Err(err).Msg("Error adding credit card")
		return err
	}
	return nil
}

func (s *BankService) GetAllCreditCards() ([]entities.CreditCard, error) {
	cards, err := s.repo.GetAllCreditCards()
	if err != nil {
		log.Error().Err(err).Msg("Error getting credit cards")
		return nil, err
	}
	return cards, nil
}

func (s *BankService) Pay(request entities.PaymentRequest) error {
	err := s.repo.Pay(request)
	if err != nil {
		log.Error().Err(err).Msg("Error processing payment")
		return err
	}
	return nil
}

func (s *BankService) Refund(request entities.RefundRequest) error {
	err := s.repo.Refund(request)
	if err != nil {
		log.Error().Err(err).Msg("Error processing refund")
		return err
	}
	return nil
}
