package repositories

import (
	"payment-Api/models"
	"gorm.io/gorm"
)

// PaymentRepository interface for payment repository methods
type PaymentRepository interface {
	CreatePayment(payment *models.Payment) error
}

// paymentRepository struct implements the PaymentRepository interface
type paymentRepository struct {
	db *gorm.DB
}

// NewPaymentRepository returns an instance of PaymentRepository
func NewPaymentRepository(db *gorm.DB) PaymentRepository {
	return &paymentRepository{db: db}
}

// CreatePayment inserts a payment record into the database
func (r *paymentRepository) CreatePayment(payment *models.Payment) error {
	return r.db.Create(payment).Error
}
