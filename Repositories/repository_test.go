package repositories

import (
	//"errors"
	"payment-Api/models"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Setup helper function to initialize SQLite in-memory database
func setupTestDB() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	// Migrate the schema for testing
	err = db.AutoMigrate(&models.Payment{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

func TestCreatePayment(t *testing.T) {
	db, err := setupTestDB()
	assert.NoError(t, err)
	repo := NewPaymentRepository(db)

	payment := &models.Payment{
		PaymentReference:   "pay123",
		Amount:             1000,
		SenderAccountRef:   "sender123",
		ReceiverAccountRef: "receiver123",
		Currency:           "INR",
	}

	t.Run("Success", func(t *testing.T) {
		err := repo.CreatePayment(payment)
		assert.NoError(t, err)
		// Verify the payment was saved
		var fetchedPayment models.Payment
		db.First(&fetchedPayment, "payment_reference = ?", payment.PaymentReference)
		assert.Equal(t, payment.PaymentReference, fetchedPayment.PaymentReference)
	})

	t.Run("Failure", func(t *testing.T) {
		// Drop the payments table to simulate a failure
		db.Migrator().DropTable(&models.Payment{})
		err := repo.CreatePayment(payment)
		assert.Error(t, err)
	})
}

func TestGetPaymentByReference(t *testing.T) {
	db, err := setupTestDB()
	assert.NoError(t, err)
	repo := NewPaymentRepository(db)

	payment := &models.Payment{
		PaymentReference: "pay123",
		Amount:           1000,
	}
	db.Create(payment)

	t.Run("Success", func(t *testing.T) {
		result, err := repo.GetPaymentByReference("pay123")
		assert.NoError(t, err)
		assert.Equal(t, payment.PaymentReference, result.PaymentReference)
	})

	t.Run("Failure", func(t *testing.T) {
		result, err := repo.GetPaymentByReference("invalid_ref")
		assert.Error(t, err)
		assert.Nil(t, result)
	})
}

func TestGetPaymentsBySender(t *testing.T) {
	db, err := setupTestDB()
	assert.NoError(t, err)
	repo := NewPaymentRepository(db)

	payment := &models.Payment{
		PaymentReference: "pay123",
		Amount:           1000,
		SenderAccountRef: "sender123",
	}
	db.Create(payment)

	t.Run("Success", func(t *testing.T) {
		results, err := repo.GetPaymentsBySender("sender123")
		assert.NoError(t, err)
		assert.Equal(t, 1, len(results))
		assert.Equal(t, payment.SenderAccountRef, results[0].SenderAccountRef)
	})

	t.Run("Failure", func(t *testing.T) {
		results, err := repo.GetPaymentsBySender("invalid_sender")
		assert.NoError(t, err)
		assert.Equal(t, 0, len(results))
	})
}

func TestGetPaymentsByReceiver(t *testing.T) {
	db, err := setupTestDB()
	assert.NoError(t, err)
	repo := NewPaymentRepository(db)

	payment := &models.Payment{
		PaymentReference:   "pay123",
		Amount:             1000,
		ReceiverAccountRef: "receiver123",
	}
	db.Create(payment)

	t.Run("Success", func(t *testing.T) {
		results, err := repo.GetPaymentsByReceiver("receiver123")
		assert.NoError(t, err)
		assert.Equal(t, 1, len(results))
		assert.Equal(t, payment.ReceiverAccountRef, results[0].ReceiverAccountRef)
	})

	t.Run("Failure", func(t *testing.T) {
		results, err := repo.GetPaymentsByReceiver("invalid_receiver")
		assert.NoError(t, err)
		assert.Equal(t, 0, len(results))
	})
}

func TestGetTotalAmountByReceiver(t *testing.T) {
	db, err := setupTestDB()
	assert.NoError(t, err)
	repo := NewPaymentRepository(db)

	payments := []models.Payment{
		{PaymentReference: "pay123", Amount: 1000, ReceiverAccountRef: "receiver123"},
		{PaymentReference: "pay124", Amount: 1500, ReceiverAccountRef: "receiver123"},
	}
	for _, p := range payments {
		db.Create(&p)
	}

	t.Run("Success", func(t *testing.T) {
		totalAmount, err := repo.GetTotalAmountByReceiver("receiver123")
		assert.NoError(t, err)
		assert.Equal(t, 2500.0, totalAmount)
	})

	t.Run("Failure", func(t *testing.T) {
		totalAmount, err := repo.GetTotalAmountByReceiver("invalid_receiver")
		assert.NoError(t, err)
		assert.Equal(t, 0.0, totalAmount)
	})
}
