package services

import (
	"payment-Api/models"
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"errors"
)

// MockPaymentRepository is a mock for the PaymentRepository
type MockPaymentRepository struct {
	mock.Mock
}

func (m *MockPaymentRepository) CreatePayment(payment *models.Payment) error {
	args := m.Called(payment)
	return args.Error(0)
}

func (m *MockPaymentRepository) GetPaymentByReference(paymentReference string) (*models.Payment, error) {
	args := m.Called(paymentReference)
	if args.Get(0) != nil {
		return args.Get(0).(*models.Payment), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockPaymentRepository) GetPaymentsBySender(senderAccountRef string) ([]models.Payment, error) {
	args := m.Called(senderAccountRef)
	if args.Get(0) != nil {
		return args.Get(0).([]models.Payment), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockPaymentRepository) GetPaymentsByReceiver(receiverAccountRef string) ([]models.Payment, error) {
	args := m.Called(receiverAccountRef)
	if args.Get(0) != nil {
		return args.Get(0).([]models.Payment), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockPaymentRepository) GetTotalAmountByReceiver(receiverAccountRef string) (float64, error) {
	args := m.Called(receiverAccountRef)
	return args.Get(0).(float64), args.Error(1)
}

// Test suite for PaymentService

func TestProcessPayment(t *testing.T) {
	mockRepo := new(MockPaymentRepository)
	service := NewPaymentService(mockRepo)

	payment := &models.Payment{
		PaymentReference:   "pay123",
		Amount:             1000,
		SenderAccountRef:   "sender123",
		ReceiverAccountRef: "receiver123",
	}

	t.Run("Success", func(t *testing.T) {
		mockRepo.On("CreatePayment", payment).Return(nil).Once()

		err := service.ProcessPayment(payment)
		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Failure", func(t *testing.T) {
		mockRepo.On("CreatePayment", payment).Return(errors.New("create error")).Once()

		err := service.ProcessPayment(payment)
		assert.Error(t, err)
		mockRepo.AssertExpectations(t)
	})
}

func TestFetchPaymentByReference(t *testing.T) {
	mockRepo := new(MockPaymentRepository)
	service := NewPaymentService(mockRepo)

	payment := &models.Payment{
		PaymentReference: "pay123",
		Amount:           1000,
	}

	t.Run("Success", func(t *testing.T) {
		mockRepo.On("GetPaymentByReference", "pay123").Return(payment, nil).Once()

		result, err := service.FetchPaymentByReference("pay123")
		assert.NoError(t, err)
		assert.Equal(t, payment, result)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Failure", func(t *testing.T) {
		mockRepo.On("GetPaymentByReference", "invalid_ref").Return(nil, errors.New("not found")).Once()

		result, err := service.FetchPaymentByReference("invalid_ref")
		assert.Error(t, err)
		assert.Nil(t, result)
		mockRepo.AssertExpectations(t)
	})
}

func TestFetchPaymentsBySender(t *testing.T) {
	mockRepo := new(MockPaymentRepository)
	service := NewPaymentService(mockRepo)

	payments := []models.Payment{
		{PaymentReference: "pay123", Amount: 1000, SenderAccountRef: "sender123"},
	}

	t.Run("Success", func(t *testing.T) {
		mockRepo.On("GetPaymentsBySender", "sender123").Return(payments, nil).Once()

		result, err := service.FetchPaymentsBySender("sender123")
		assert.NoError(t, err)
		assert.Equal(t, payments, result)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Failure", func(t *testing.T) {
		mockRepo.On("GetPaymentsBySender", "invalid_sender").Return(nil, errors.New("not found")).Once()

		result, err := service.FetchPaymentsBySender("invalid_sender")
		assert.Error(t, err)
		assert.Nil(t, result)
		mockRepo.AssertExpectations(t)
	})
}

func TestFetchPaymentsByReceiver(t *testing.T) {
	mockRepo := new(MockPaymentRepository)
	service := NewPaymentService(mockRepo)

	payments := []models.Payment{
		{PaymentReference: "pay123", Amount: 1000, ReceiverAccountRef: "receiver123"},
	}

	t.Run("Success", func(t *testing.T) {
		mockRepo.On("GetPaymentsByReceiver", "receiver123").Return(payments, nil).Once()

		result, err := service.FetchPaymentsByReceiver("receiver123")
		assert.NoError(t, err)
		assert.Equal(t, payments, result)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Failure", func(t *testing.T) {
		mockRepo.On("GetPaymentsByReceiver", "invalid_receiver").Return(nil, errors.New("not found")).Once()

		result, err := service.FetchPaymentsByReceiver("invalid_receiver")
		assert.Error(t, err)
		assert.Nil(t, result)
		mockRepo.AssertExpectations(t)
	})
}

func TestFetchTotalAmountByReceiver(t *testing.T) {
	mockRepo := new(MockPaymentRepository)
	service := NewPaymentService(mockRepo)

	t.Run("Success", func(t *testing.T) {
		mockRepo.On("GetTotalAmountByReceiver", "receiver123").Return(2000.0, nil).Once()

		total, err := service.FetchTotalAmountByReceiver("receiver123")
		assert.NoError(t, err)
		assert.Equal(t, 2000.0, total)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Failure", func(t *testing.T) {
		mockRepo.On("GetTotalAmountByReceiver", "invalid_receiver").Return(0.0, errors.New("not found")).Once()

		total, err := service.FetchTotalAmountByReceiver("invalid_receiver")
		assert.Error(t, err)
		assert.Equal(t, 0.0, total)
		mockRepo.AssertExpectations(t)
	})
}
