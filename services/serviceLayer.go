package services

import (
	"payment-Api/models"
	"payment-Api/repositories"
)

// PaymentService interface for service methods
type PaymentService interface {
	ProcessPayment(payment *models.Payment) error
}

// paymentService struct implements the PaymentService interface
type paymentService struct {
	paymentRepo repositories.PaymentRepository
}

// NewPaymentService returns an instance of PaymentService
func NewPaymentService(paymentRepo repositories.PaymentRepository) PaymentService {
	return &paymentService{paymentRepo: paymentRepo}
}

// ProcessPayment contains business logic and persists the payment
func (s *paymentService) ProcessPayment(payment *models.Payment) error {
	// Business logic can be added here if needed
	return s.paymentRepo.CreatePayment(payment)
}