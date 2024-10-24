package repositories

import (
	"payment-Api/models"
	"gorm.io/gorm"
	"fmt"
)

// PaymentRepository interface for payment repository methods
type PaymentRepository interface {
	CreatePayment(payment *models.Payment) error
	GetPaymentByReference(paymentReference string) (*models.Payment, error)
	
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



// GetPaymentByReference fetches a payment by its payment_reference using GORM
func (r *paymentRepository) GetPaymentByReference(paymentReference string) (*models.Payment, error) {
    var payment models.Payment
    result := r.db.Where("payment_reference = ?", paymentReference).First(&payment)
	fmt.Println("RESULT",result)
    if result.Error != nil {
        return nil, result.Error
    }
    return &payment, nil
}