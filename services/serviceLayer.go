package services

import (
	"payment-Api/models"
	"payment-Api/repositories"
)

// PaymentService interface for service methods
type PaymentService interface {
	ProcessPayment(payment *models.Payment) error
	FetchPaymentByReference(paymentReference string) (*models.Payment, error)
	FetchPaymentsBySender(senderAccountRef string) ([]models.Payment, error)
	FetchPaymentsByReceiver(receiverAccountRef string) ([]models.Payment,error)


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

//fetching the Payment data by 'Payment_Reference'
func (s *paymentService) FetchPaymentByReference(paymentReference string) (*models.Payment, error) {
    return s.paymentRepo.GetPaymentByReference(paymentReference)
}

//fetching the Payment data by 'sender_acc_ref'
func (s *paymentService) FetchPaymentsBySender(senderAccountRef string) ([]models.Payment, error){
	return s.paymentRepo.GetPaymentsBySender(senderAccountRef)
}

//fetching the payment details by 'receiver_acc-ref'
func (s *paymentService) FetchPaymentsByReceiver(receiverAccountRef string) ([]models.Payment,error){
	return s.paymentRepo.GetPaymentsByReceiver(receiverAccountRef)
}
